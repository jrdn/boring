package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jrdn/boring/log"
	"github.com/jrdn/boring/server/middleware"
	"github.com/jrdn/boring/sig"
	"go.uber.org/zap"
)

func New(listen string, serverLogger *zap.Logger, requestLogger *zap.Logger) *Server {
	if serverLogger == nil {
		serverLogger = log.GetLogger("server")
		serverLogger.Info("starting server", zap.String("server_addr", listen))
	}

	if requestLogger == nil {
		requestLogger = serverLogger.Named("request")
	}

	srv := &http.Server{
		Addr: listen,
	}

	server := &Server{
		HttpServer:    srv,
		Logger:        serverLogger,
		RequestLogger: requestLogger,
		Router:        chi.NewRouter(),
	}

	server.Middleware(middleware.LoggingMiddleware(requestLogger))

	return server
}

type Server struct {
	HttpServer    *http.Server
	StopCallback  func(*zap.Logger)
	Logger        *zap.Logger
	RequestLogger *zap.Logger
	Router        chi.Router
}

func (s *Server) Mount(pattern string, handler http.Handler) {
	s.Router.Mount(pattern, handler)
}

func (s *Server) Middleware(middleware func(http.Handler) http.Handler) {
	s.Router = s.Router.With(middleware)
}

func (s *Server) Run() {
	s.HttpServer.Handler = s.Router

	go func() {
		_ = s.HttpServer.ListenAndServe()
	}()

	sig.WaitUntilSignalled()
	_ = s.HttpServer.Shutdown(context.Background())

	if s.StopCallback != nil {
		s.StopCallback(s.Logger)
	}
}
