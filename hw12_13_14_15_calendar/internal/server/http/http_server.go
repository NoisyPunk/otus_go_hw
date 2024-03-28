package internalhttp

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
)

type HTTPEventServer struct {
	application app.Application
	logger      *zap.Logger
	server      http.Server
}

func NewServer(app app.Application, config *configs.Config, logger *zap.Logger) *HTTPEventServer {
	return &HTTPEventServer{
		application: app,
		logger:      logger,
		server: http.Server{
			Addr:              net.JoinHostPort(config.Host, config.Port),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

func (s *HTTPEventServer) Start(ctx context.Context) error {
	l := logger.FromContext(ctx)

	mux := http.NewServeMux()
	mux.Handle("/create", loggingMiddleware(http.HandlerFunc(s.CreateEvent), l))
	mux.Handle("/update", loggingMiddleware(http.HandlerFunc(s.UpdateEvent), l))
	mux.Handle("/delete", loggingMiddleware(http.HandlerFunc(s.DeleteEvent), l))
	mux.Handle("/daily", loggingMiddleware(http.HandlerFunc(s.EventsDailyList), l))
	mux.Handle("/weekly", loggingMiddleware(http.HandlerFunc(s.EventsWeeklyList), l))
	mux.Handle("/monthly", loggingMiddleware(http.HandlerFunc(s.EventsMonthlyList), l))

	s.server.Handler = mux

	go func() error {
		err := s.server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("http server started", zap.String("server address", s.server.Addr))
	<-ctx.Done()
	return nil
}

func (s *HTTPEventServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
