package main

import (
	"bytes"
	"fmt"
	"os"
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

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "walktest")
	if err != nil {
		t.Fatal(err)
	}

	for k, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, k)
			fpath := filepath.Join(tempDir, fname)
			if err := os.WriteFile(fpath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}

	return tempDir, func() { os.RemoveAll(tempDir) }
}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		c           config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{
		{
			name:        "DeleteExtensionNoMatch",
			c:           config{ext: ".log", del: true},
			extNoDelete: ".gz",
			nDelete:     0,
			nNoDelete:   10,
			expected:    "",
		},
		{
			name:        "DeleteExtensionMatch",
			c:           config{ext: ".log", del: true},
			extNoDelete: "",
			nDelete:     10,
			nNoDelete:   0,
			expected:    "",
		},
		{
			name:        "DeleteExtensionMixed",
			c:           config{ext: ".log", del: true},
			extNoDelete: ".gz",
			nDelete:     5,
			nNoDelete:   5,
			expected:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.c.ext:       tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})
			defer cleanup()

			tc.c.root = tempDir

			if err := run(&buffer, &tc.c); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}

			filesLeft, err := os.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}

			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("Expected %d files left, got %d instead \n", tc.nNoDelete, len(filesLeft))
			}
		})
	}
}
