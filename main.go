package main

import (
	"os"

	"github.com/emzola/filesystemcrawler/cmd"
)

func main() {
	err := cmd.HandleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
