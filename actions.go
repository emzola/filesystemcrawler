package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func filterOut(path, ext string, minSize int64, info fs.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if len(ext) != 0 && filepath.Ext(path) != ext {
		return true
	}

	return false
}

func listFile(w io.Writer, path string) error {
	_, err := fmt.Fprintln(w, path)
	return err
}

func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

func writeToLog(c *config) (*os.File, error) {
	file, err := os.OpenFile(c.logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return file, nil
}
