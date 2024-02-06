package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset bigger than filesize", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test.txt", 10000, 1000)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})
	t.Run("empty file", func(t *testing.T) {
		err := Copy("testdata/zero.txt", "testdata/test.txt", 1, 1000)
		require.Equal(t, ErrUnsupportedFile, err)
	})
	t.Run("offset equal than filesize", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test.txt", 6617, 6617)
		require.ErrorIs(t, ErrOffsetExceedsFileSize, err)
	})
	t.Run("params less than zero", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test.txt", -1, -1)
		require.ErrorIs(t, ErrParamsLessZero, err)
	})
}
