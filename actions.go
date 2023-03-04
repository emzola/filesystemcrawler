package main

import (
	"fmt"
	"io"
	"io/fs"
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

func delFile(path string) error {
	return os.Remove(path)
}
