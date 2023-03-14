package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	binName := "filesystemcrawler"

	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Exit(result)
}
