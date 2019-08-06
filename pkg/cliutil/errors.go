package cliutil

import (
	"github.com/solo-io/go-utils/errors"
)

var (
	NoFileOrSubcommandError = errors.New("please provide a file flag or subcommand")

	NoSubcommandError = errors.New("please select a subcommand")
)
