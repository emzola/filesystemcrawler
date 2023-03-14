package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	usageMessage := "A command-line tool which crawls into file system directories, find files and executes actions.\n\nUsage of %s: <options> [name]File System Crawler\nOptions: \n  -archive string\n    \tArchive directory\n  -date string\n    \tFilter by modified date (format: 2006-Jan-02)\n  -del\n    \tDelete files\n  -ext string\n    \tFilter by file extension\n  -list\n    \tList files\n  -log string\n    \tLog file deletes to this file\n  -root string\n    \tRoot directory to start (default \".\")\n  -size int\n    \tFilter by minimum file size\n"
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"ParseHelpArgument", []string{"-h"}, usageMessage, errors.New("flag: help requested")},
		{"ParseInvalidArgument", []string{"-foo"}, "flag provided but not defined: -foo\n" + usageMessage, errors.New("invalid subcommand specified")},
		{"ParseValidArgument", []string{"-root", "testdata", "-date", "2022-Mar-11", "-list"}, "testdata\\dir.log\ntestdata\\dir2\\script.sh\n", nil},
		{"ParsePositionalArgument", []string{"-root", "testdata", "positional argument"}, "", ErrPosArgSpecified},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := parseArgs(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
