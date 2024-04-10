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

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app/sender"
	senderconfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/sender_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/sender_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := senderconfig.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	app, err := sender.New(log, config)
	if err != nil {
		fmt.Printf("can't connect to db: %s", err.Error())
		os.Exit(1)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		app.Consume(ctx)
		wg.Done()
	}()

	log.Info("Sender is running...", zap.String("start_time", time.Now().String()))
	<-ctx.Done()
	wg.Wait()
}
