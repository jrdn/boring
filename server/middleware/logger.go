package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error("paniced", zap.Any("recovered", recovered), zap.ByteString("trace", debug.Stack()))
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Debug("request",
				zap.Int("status", wrapped.status),
				zap.String("method", r.Method),
				zap.String("path", r.URL.EscapedPath()),
				zap.Duration("duration", time.Since(start)),
			)
		}

		return http.HandlerFunc(fn)
	}
}
