package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send(result chan<- error) error
	Receive(result chan<- error) error
}

type tclient struct {
	host    string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	err     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out, err io.Writer) TelnetClient {
	return &tclient{
		host:    address,
		timeout: timeout,
		in:      in,
		out:     out,
		err:     err,
	}
}

func (tc *tclient) Connect() error {
	var err error
	tc.conn, err = net.DialTimeout("tcp", tc.host, tc.timeout)
	if err != nil {
		return wrapErr(err, ErrUnableConnect)
	}
	return tc.stdErrLogs(fmt.Sprintf("Connected to server with address: %s", tc.host))
}

func (tc *tclient) Close() error {
	err := tc.conn.Close()
	if err != nil {
		return err
	}
	return tc.stdErrLogs(fmt.Sprintf("Connection to server  %s closed", tc.host))
}

func (tc *tclient) Send(result chan<- error) error {
	_, err := io.Copy(tc.conn, tc.in)
	if err != nil {
		result <- err
	}
	return tc.stdErrLogs(fmt.Sprintf("Sended message: %s", tc.in))
}

func (tc *tclient) Receive(result chan<- error) error {
	_, err := io.Copy(tc.out, tc.conn)
	if err != nil {
		result <- err
	}
	return tc.stdErrLogs(fmt.Sprintf("Received message: %s", tc.in))
}

func (tc *tclient) stdErrLogs(message string) error {
	_, err := tc.err.Write([]byte(message))
	return err
}
