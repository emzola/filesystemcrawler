package cmd

import (
	"errors"
	"fmt"
	"io"

	restore "github.com/emzola/filesystemcrawler/cmd/restore"
	walk "github.com/emzola/filesystemcrawler/cmd/walk"
)

func printUsage(w io.Writer) {
	usageMessage := `Usage: File System Crawler [walk|restore] -h.

A command-line tool which crawls into file system directories.`
	fmt.Fprintln(w, usageMessage)
	walk.HandleWalk(w, []string{"-h"})
	restore.HandleRestore(w, []string{"-h"})
}

func HandleCommand(w io.Writer, args []string) error {
	var err error

	if len(args) < 1 {
		return ErrInvalidSubCommand
	}

	switch args[0] {
	case "walk":
		err = walk.HandleWalk(w, args[1:])
	case "restore":
		err = restore.HandleRestore(w, args[1:])
	case "-h", "--help":
		printUsage(w)
	default:
		err = ErrInvalidSubCommand
	}

	if errors.Is(err, ErrInvalidSubCommand) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}

	return err
}
