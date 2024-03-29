package internalhttp

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Server struct {
	Application Application
	server      http.Server
}

type Application interface {
	CreateEvent(ctx context.Context, data storage.Event, userID uuid.UUID) (uuid.UUID, error)
}

func NewServer(app Application, config *configs.Config) *Server {
	return &Server{
		Application: app,
		server: http.Server{
			Addr:              net.JoinHostPort(config.Host, config.Port),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	l := logger.FromContext(ctx)

	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(s.getHello), l))

	s.server.Handler = mux

	go func() error {
		err := s.server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("server started", zap.String("server address", s.server.Addr))
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) getHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, OTUS!\n")
}
