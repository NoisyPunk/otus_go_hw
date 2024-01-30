package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return err
	}

	size, err := fileSize(file)
	if err != nil {
		return ErrUnsupportedFile
	}
	if size == 0 {
		return ErrUnsupportedFile
	}
	if offset >= size {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 {
		err = copyFullFile(file, toPath, size)
		if err != nil {
			return err
		}
		return nil
	}
	err = copyWithBuffer(file, toPath, offset, limit, size)
	if err != nil {
		return err
	}
	return nil
}

func copyFullFile(file *os.File, toPath string, size int64) error {
	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(io.MultiWriter(newFile, makeProgressBar(size)), file)
	if err != nil {
		return err
	}

	return nil
}

func copyWithBuffer(file *os.File, toPath string, offset, limit, size int64) error {
	switch {
	case size < limit:
		limit = size
	case offset > limit:
		if limit+offset > size {
			limit = size - offset
		}
	}
	buf := make([]byte, limit)

	shift, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.ReadAt(buf, shift)
	if err != nil {
		if errors.Is(err, io.EOF) {
			_ = err
		} else {
			return err
		}
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(buf)
	_, err = io.Copy(io.MultiWriter(newFile, makeProgressBar(limit)), reader)
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}
	return nil
}

func fileSize(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	size := fileInfo.Size()
	return size, nil
}

func makeProgressBar(maxLenght int64) *progressbar.ProgressBar {
	barDescription := "Copying in progress"
	bar := progressbar.DefaultBytes(maxLenght, barDescription)
	return bar
}
