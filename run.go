package main

import (
	"io"
	"io/fs"
	"path/filepath"
)

func run(w io.Writer, c *config) error {
	return filepath.Walk(c.root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, c.ext, c.size, info) {
			return nil
		}

		if c.list {
			return listFile(w, path)
		}

		return listFile(w, path)
	})
}
