package walk

import (
	"os"
	"testing"
	"time"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		ext      []string
		minSize  int64
		modDate  time.Time
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", []string{""}, 0, time.Time{}, false},
		{"FilterExtensionMatch", "testdata/dir.log", []string{".log"}, 0, time.Time{}, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", []string{".sh", ".pdf"}, 0, time.Time{}, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", []string{".log"}, 10, time.Time{}, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", []string{".log"}, 30, time.Time{}, true},
		{"FilterExtensionTimeAfter", "testdata/dir.log", []string{".log"}, 0, time.Now(), true},
		{"FilterExtensionTimeBefore", "testdata/dir.log", []string{".log"}, 0, time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.path)
			if err != nil {
				t.Fatal(err)
			}

			f := filterOut(tc.path, tc.ext, tc.minSize, tc.modDate, info)

			if f != tc.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}
