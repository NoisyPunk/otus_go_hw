package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app/scheduler"
	schedulerConfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./scheduler_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := schedulerConfig.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	app, err := scheduler.New(ctx, log, config)
	if err != nil {
		fmt.Printf("can't connect to rmq: %s", err.Error())
		os.Exit(1)
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		app.OldEventRemover(ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		app.Notifier(ctx)
		wg.Done()
	}()

	log.Info("Scheduler is running...", zap.String("start_time", time.Now().String()))
	<-ctx.Done()
	wg.Wait()
}
