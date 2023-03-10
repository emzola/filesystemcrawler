package cmd

import "errors"

var (
	ErrInvalidSubCommand = errors.New("invalid subcommand specified")
)
