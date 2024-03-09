package internalhttp

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"net/http"
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
	mux.HandleFunc("/hello", getHello)

	addr := s.host + ":" + s.port

	go func() error {
		err := http.ListenAndServe(addr, mux)
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
	// TODO
	return nil
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, OTUS!\n")
}
