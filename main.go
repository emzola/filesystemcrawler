package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
	root string
}

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

func main() {
	c, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = run(os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
