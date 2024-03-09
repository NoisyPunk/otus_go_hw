package main

import (
	"context"
	"flag"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
)

var (
	configFile string
	storage    app.Storage
)

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	config := configs.GetConfig(configFile)

	log := logger.New(config.LogLevel)
	ctx := logger.ContextLogger(context.Background(), log)
	if config.InmemStore {
		storage = memorystorage.New()
		log.Debug("inmemory storage is used for server")
	}
	calendar := app.New(ctx, storage)

	server := internalhttp.NewServer(calendar, config)

	ctx, cancel := signal.NotifyContext(ctx,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
