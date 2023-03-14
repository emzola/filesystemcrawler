package main

import (
	"os"
)

func main() {
	err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
