package walk

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func filterOut(path string, ext []string, minSize int64, modDate time.Time, info fs.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	var extMatch bool
	for _, e := range ext {
		if e == "" || filepath.Ext(path) == e {
			extMatch = true
		}
	}

	if !extMatch {
		return true
	}

	if !modDate.IsZero() {
		return modDate.After(info.ModTime())
	}

	return false
}

func listFile(w io.Writer, path string) error {
	_, err := fmt.Fprintln(w, path)
	return err
}

func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

func writeToLog(c *walkConfig) (*os.File, error) {
	file, err := os.OpenFile(c.logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return file, nil
}

func archiveFile(destDir, root, path string) error {
	info, err := os.Stat(destDir)
	if err != nil {
		return err
	}

	// Check whether destination is a directory
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", destDir)
	}

	// Determine the relative directory of the file to be archived in relation to it's source root path
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		return err
	}

	dest := fmt.Sprintf("%s.gz", filepath.Base(path))
	targetPath := filepath.Join(destDir, relDir, dest)

	// Create target directory tree
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	// Create the compressed archive
	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	zw := gzip.NewWriter(out)

	zw.Name = filepath.Base(path)

	if _, err := io.Copy(zw, in); err != nil {
		return err
	}

	if err := zw.Close(); err != nil {
		return err
	}

	return out.Close()
}
