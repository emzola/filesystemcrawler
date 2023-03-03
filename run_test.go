package main

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {

	logPath := filepath.Join("testdata", "dir.log") + "\n"
	shPath := filepath.Join("testdata", "dir2", "script.sh") + "\n"

	testCases := []struct {
		name     string
		c        config
		expected string
	}{
		{name: "NoFilter", c: config{ext: "", size: 0, list: true, root: "testdata"}, expected: logPath + shPath},
		{name: "FilterExtensionMatch", c: config{ext: ".log", size: 0, list: true, root: "testdata"}, expected: logPath},
		{name: "FilterExtensionSizeMatch", c: config{ext: ".log", size: 10, list: true, root: "testdata"}, expected: logPath},
		{name: "FilterExtensionSizeNoMatch", c: config{ext: ".log", size: 30, list: true, root: "testdata"}, expected: ""},
		{name: "FilterExtensionNoMatch", c: config{ext: ".gz", size: 0, list: true, root: "testdata"}, expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := run(&buffer, &tc.c); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}

		})
	}
}
