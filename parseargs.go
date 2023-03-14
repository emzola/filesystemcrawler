package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
)

type config struct {
	ext     []string
	size    int64
	list    bool
	root    string
	del     bool
	logFile string
	out     io.Writer
	archive string
	modDate time.Time
}

var ErrPosArgSpecified = errors.New("positional argument specified")

func parseArgs(w io.Writer, args []string) error {
	c := &config{}
	var err error
	var ext string
	var modDate string

	fs := flag.NewFlagSet("File System Crawler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.root, "root", ".", "Root directory to start")
	fs.StringVar(&ext, "ext", "", "Filter by file extension")
	fs.Int64Var(&c.size, "size", 0, "Filter by minimum file size")
	fs.BoolVar(&c.list, "list", false, "List files")
	fs.StringVar(&c.archive, "archive", "", "Archive directory")
	fs.StringVar(&modDate, "date", "", "Filter by modified date (format: 2006-Jan-02)")
	fs.StringVar(&c.logFile, "log", "", "Log file deletes to this file")
	fs.BoolVar(&c.del, "del", false, "Delete files")
	fs.Usage = func() {
		usageMessage := `A command-line tool which crawls into file system directories, find files and executes actions.

Usage of %s: <options> [name]`
		fmt.Fprint(w, usageMessage, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err = fs.Parse(args)
	if err != nil {
		return err
	}

	if fs.NArg() != 0 {
		return ErrPosArgSpecified
	}

	if len(modDate) != 0 {
		modDate, err := time.Parse("2006-Jan-02", modDate)
		if err != nil {
			return err
		}
		c.modDate = modDate
	}

	c.ext = strings.Split(ext, "|")

	err = run(w, c)
	if err != nil {
		return err
	}

	return nil
}
