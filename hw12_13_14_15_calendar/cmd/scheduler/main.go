package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app/scheduler"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/scheduler_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := scheduler_config.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	app, err := scheduler.New(ctx, config)
	if err != nil {
		fmt.Printf("can't connect to db: %s", err.Error())
		cancel()
		os.Exit(1)
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-ctx.Done()
		wg.Done()
	}()

	go func() {
		app.OldEventRemover(ctx)
	}()

	go func() {
		app.Notifier(ctx)
	}()

	log.Info("Scheduler is running...", zap.String("start_time", time.Now().String()))
	wg.Wait()
}
