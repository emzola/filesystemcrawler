package walk

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
)

type walkConfig struct {
	ext     []string
	size    int64
	list    bool
	root    string
	del     bool
	logFile string
	out     io.Writer
	archive string
	modDate time.Time
}

func HandleWalk(w io.Writer, args []string) error {
	c := &walkConfig{}
	var err error
	var ext string
	var modDate string

	fs := flag.NewFlagSet("walk", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.root, "root", ".", "Root directory to start")
	fs.StringVar(&ext, "ext", "", "Filter by file extension")
	fs.Int64Var(&c.size, "size", 0, "Filter by minimum file size")
	fs.BoolVar(&c.list, "list", false, "List files")
	fs.StringVar(&c.archive, "archive", "", "Archive directory")
	fs.StringVar(&modDate, "date", "", "Filter by modified date (format: 2006-Jan-02)")
	fs.StringVar(&c.logFile, "log", "", "Log file deletes to this file")
	fs.BoolVar(&c.del, "del", false, "Delete files")
	fs.Usage = func() {
		usageMessage := `
walk: Finds files which match a specific criteria and executes actions.

walk: <options>`
		fmt.Fprint(w, usageMessage)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	err = fs.Parse(args)
	if err != nil {
		return err
	}

	if fs.NArg() != 0 {
		return fmt.Errorf("error: %s", "positional arguments must not be specified")
	}

	if len(modDate) != 0 {
		modDate, err := time.Parse("2006-Jan-02", modDate)
		if err != nil {
			return err
		}
		c.modDate = modDate
	}

	c.ext = strings.Split(ext, "|")

	err = run(w, c)
	if err != nil {
		return err
	}

	return nil
}
