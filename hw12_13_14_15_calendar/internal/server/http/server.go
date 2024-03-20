package internalhttp

import (
	"context"
	"io"
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
	host        string
	port        string
}

type Application interface {
	CreateEvent(ctx context.Context, data storage.Event, userID uuid.UUID) (uuid.UUID, error)
}

func NewServer(app Application, config *configs.Config) *Server {
	return &Server{
		Application: app,
		host:        config.Host,
		port:        config.Port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	l := logger.FromContext(ctx)

	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(s.getHello), l))

	addr := s.host + ":" + s.port

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	go func() error {
		err := server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("server started", zap.String("server address", addr))
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	_ = ctx
	// TODO
	return nil
}

func (s *Server) getHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, OTUS!\n")
}
