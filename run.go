package main

import (
	"io"
	"io/fs"
	"log"
	"path/filepath"
)

func run(w io.Writer, c *config) error {
	var delLogger *log.Logger

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

		if len(c.logFile) != 0 {
			file, err := writeToLog(c)
			if err != nil {
				return err
			}
			c.out = file
			defer file.Close()
		}

		if len(c.archive) != 0 {
			if err := archiveFile(c.archive, c.root, path); err != nil {
				return err
			}
		}

		if c.del {
			delLogger = log.New(c.out, "DELETED FILE: ", log.LstdFlags)
			return delFile(path, delLogger)
		}

		return listFile(w, path)
	})
}
