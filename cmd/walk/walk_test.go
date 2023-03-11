package walk

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleWalk(t *testing.T) {
	usageMessage := "\nwalk: Finds files which match a specific criteria and executes actions.\n\nwalk: <options>\nOptions: \n  -archive string\n    \tArchive directory\n  -date string\n    \tFilter by modified date (format: 2006-Jan-02)\n  -del\n    \tDelete files\n  -ext string\n    \tFilter by file extension\n  -list\n    \tList files\n  -log string\n    \tLog file deletes to this file\n  -root string\n    \tRoot directory to start (default \".\")\n  -size int\n    \tFilter by minimum file size\n"
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"HandleWalkHelpFlag", []string{"-h"}, usageMessage, errors.New("flag: help requested")},
		{"HandleWalkInvalidFlag", []string{"-foo"}, "flag provided but not defined: -foo\n" + usageMessage, errors.New("invalid subcommand specified")},
		{"HandleWalkValidFlags", []string{"-root", "testdata", "-date", "2022-Mar-11", "-list"}, "testdata\\dir.log\ntestdata\\dir2\\script.sh\n", nil},
		{"HandleWalkPositionalArgument", []string{"-root", "testdata", "positional argument"}, "", ErrPosArgSpecified},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := HandleWalk(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
