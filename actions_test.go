package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		ext      []string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", []string{""}, 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", []string{".log"}, 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", []string{".sh", ".pdf"}, 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", []string{".log"}, 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", []string{".log"}, 30, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.path)
			if err != nil {
				t.Fatal(err)
			}

			f := filterOut(tc.path, tc.ext, tc.minSize, info)

			if f != tc.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}
