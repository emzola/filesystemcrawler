package main

import (
	"fmt"
	"io"
	"os"
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
