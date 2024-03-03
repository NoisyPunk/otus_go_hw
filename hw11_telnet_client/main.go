package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10, "timeout for connection, 10s is default")
}

func main() {
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)

	if host == "" || port == "" {
		log.Fatal(ErrEmptyAddr)
	}

	address := net.JoinHostPort(host, port)

	reader := io.ReadCloser(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	serviceWriter := bufio.NewWriter(os.Stderr)

	client := NewTelnetClient(address, timeout, reader, writer, serviceWriter)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	result := make(chan error, 1)
	go func() {
		_ = client.Send(result)
	}()
	go func() {
		_ = client.Receive(result)
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-result:
		if err != nil {
			client.Close()
			log.Fatal(err)
		}
	case <-signals:
	}
	client.Close()
}
