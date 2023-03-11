package cmd

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := "Usage: File System Crawler [walk|restore] -h.\n\nA command-line tool which crawls into file system directories.\n\nwalk: Finds files which match a specific criteria and executes actions.\n\nwalk: <options>\nOptions: \n  -archive string\n    \tArchive directory\n  -date string\n    \tFilter by modified date (format: 2006-Jan-02)\n  -del\n    \tDelete files\n  -ext string\n    \tFilter by file extension\n  -list\n    \tList files\n  -log string\n    \tLog file deletes to this file\n  -root string\n    \tRoot directory to start (default \".\")\n  -size int\n    \tFilter by minimum file size\n"
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"HandleCommandNoArgument", []string{}, "invalid subcommand specified\n" + usageMessage, ErrInvalidSubCommand},
		{"HandleCommandHelpArgument", []string{"-h"}, usageMessage, nil},
		{"HandleCommandInvalidArgument", []string{"-foo"}, "invalid subcommand specified\n" + usageMessage, ErrInvalidSubCommand},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := HandleCommand(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
