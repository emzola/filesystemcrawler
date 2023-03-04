package main

import (
	"io"
	"io/fs"
	"log"
	"path/filepath"
)

func run(w io.Writer, c *config) error {
	delLogger := log.New(c.wLog, "DELETED FILE: ", log.LstdFlags)
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

		if c.del {
			return delFile(path, delLogger)
		}

		return listFile(w, path)
	})
}
