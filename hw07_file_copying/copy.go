package main

import (
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrParamsLessZero        = errors.New("offset or limit cannot be less than zero")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 || limit < 0 {
		return ErrParamsLessZero
	}

	file, err := os.Open(fromPath)
	defer func() { file.Close() }()
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
		return copyFullFile(file, toPath, size)
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

	readSeeker := io.ReadSeeker(file)
	_, err := readSeeker.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	_, err = io.CopyN(io.MultiWriter(newFile, makeProgressBar(limit)), readSeeker, limit)
	if err != nil {
		return err
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
