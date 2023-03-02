package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

func parseArgs(w io.Writer, args []string) (*config, error) {
	c := &config{}

	fs := flag.NewFlagSet("File System Crawler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.root, "root", ".", "Root directory to start")
	fs.StringVar(&c.ext, "ext", "", "File extension to filter out")
	fs.Int64Var(&c.size, "size", 0, "Minimum file size")
	fs.BoolVar(&c.list, "list", false, "List files only")

	fs.Usage = func() {
		usageMessage := `
A file system crawler application which crawls into file system directories looking for specific files.

Usage of %s: <options> [name]`
		fmt.Fprintf(w, usageMessage, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		return c, errors.New("positional arguments must not be specified")
	}

	return c, nil
}
