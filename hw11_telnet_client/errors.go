package main

import "fmt"

var (
	ErrUnableConnect = fmt.Errorf("unable to connect to server")
	ErrEmptyAddr     = fmt.Errorf("host and port shouldn't be empty")
)

func wrapErr(err, expectedErr error) error {
	return fmt.Errorf("%w: %w", expectedErr, err)
}
