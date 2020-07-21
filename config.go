package main

import (
	"runtime"

	"github.com/jessevdk/go-flags"
)

type config struct {
	DataDir string `short:"D" long:"datadir" description:"Directory to store data"`
}

// newConfigParser returns a new command line flags parser.
func newConfigParser(cfg *config, so *serviceOptions, options flags.Options) *flags.Parser {
	parser := flags.NewParser(cfg, options)
	if runtime.GOOS == "windows" {
		parser.AddGroup("Service Options", "Service Options", so)
	}
	return parser
}
