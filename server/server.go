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

func New(listen string) *Server {
	l := log.GetLogger("server")
	l.Info("starting server", zap.String("server_addr", listen))

	srv := &http.Server{
		Addr:    listen,
		Handler: chi.NewRouter(),
	}

	server := &Server{
		HttpServer: srv,
		Logger:     l,
	}

	server.Middleware(middleware.LoggingMiddleware(l))
	return server
}

type Server struct {
	HttpServer   *http.Server
	Router       *chi.Mux
	StopCallback func(*zap.Logger)
	Logger       *zap.Logger
}

func (s *Server) Mount(pattern string, handler http.Handler) {
	s.Router.Mount(pattern, handler)
}

func (s *Server) Middleware(middleware func(http.Handler) http.Handler) {
	s.Router.With(middleware)
}

func (s *Server) Run() {
	go func() {
		_ = s.HttpServer.ListenAndServe()
	}()

	sig.WaitUntilSignalled()
	_ = s.HttpServer.Shutdown(context.Background())

	if s.StopCallback != nil {
		s.StopCallback(s.Logger)
	}
}
