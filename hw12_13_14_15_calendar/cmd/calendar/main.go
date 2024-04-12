package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app/calendar"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/calendar_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./calendar_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := calendarconfig.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	err = sqlstorage.Migrate(config)
	if err != nil {
		log.Error("migration has failed", zap.String("error_message", err.Error()))
		os.Exit(1)
	}

	app, err := calendar.New(ctx, config)
	if err != nil {
		log.Error("failed to connect to db", zap.String("error_message", err.Error()))
		os.Exit(1)
	}

	server := internalhttp.NewServer(ctx, app, config, log)
	grpcServer := internalgrpc.NewGRPCServer(ctx, app, config.EventServerPort)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-ctx.Done()

		if err := server.Stop(); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
		grpcServer.Stop()
		wg.Done()
	}()

	go func() {
		if err = server.Start(); err != nil {
			log.Error("failed to start http server", zap.String("error_message", err.Error()))
			if err := server.Stop(); err != nil {
				log.Error("failed to stop http server: " + err.Error())
			}
		}
	}()

	go func() {
		if err = grpcServer.Start(); err != nil {
			log.Error("failed to start grpc server", zap.String("error", err.Error()))
			grpcServer.Stop()
		}
	}()

	log.Info("calendar is running...", zap.String("start_time", time.Now().String()))
	wg.Wait()
}
