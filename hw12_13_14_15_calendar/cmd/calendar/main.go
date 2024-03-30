package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := configs.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	calendar, err := app.New(ctx, config)
	if err != nil {
		fmt.Printf("can't connect to db: %s", err.Error())
		cancel()
		os.Exit(1)
	}

	server := internalhttp.NewServer(ctx, calendar, config, log)
	grpcServer := internalgrpc.NewGRPCServer(ctx, calendar, config.EventServerPort)

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
			log.Error("failed to start http server", zap.String("error", err.Error()))
			cancel()
			os.Exit(1)
		}
	}()

	go func() {
		if err = grpcServer.Start(); err != nil {
			log.Error("failed to start grpc server", zap.String("error", err.Error()))
			cancel()
			os.Exit(1)
		}
	}()

	log.Info("calendar is running...")
	wg.Wait()
}
