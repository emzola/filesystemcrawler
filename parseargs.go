package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func printUsage(w io.Writer, args []string) {
	parseArgs(w, args)
}

func parseArgs(w io.Writer, args []string) (*config, error) {
	c := &config{}
	var err error
	var ext string

	fs := flag.NewFlagSet("File System Crawler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.root, "root", ".", "Root directory to start")
	fs.StringVar(&ext, "ext", "", "File extension to filter out")
	fs.Int64Var(&c.size, "size", 0, "Minimum file size")
	fs.BoolVar(&c.list, "list", false, "List files only")
	fs.BoolVar(&c.del, "del", false, "Delete files")
	fs.StringVar(&c.logFile, "log", "", "Log deletes to this file")
	fs.StringVar(&c.archive, "archive", "", "Archive directory")
	fs.Usage = func() {
		usageMessage := `
A file system crawler application which crawls into file system directories looking for specific files.

Usage of %s: <options> [name]`
		fmt.Fprintf(w, usageMessage, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err = fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		printUsage(w, []string{"-h"})
		return c, fmt.Errorf("error: %s", "positional arguments must not be specified")
	}

	c.ext = strings.Split(ext, "|")

	return c, nil
}
