package main

import (
	"github.com/ExploitSuite/cdncheck/internal/runner"
	"github.com/projectdiscovery/gologger"
)

func main() {
	options := runner.ParseOptions()

	newRunner := runner.NewRunner(options)

	err := newRunner.Run()
	if err != nil {
		gologger.Fatal().Msgf("Could not run cdncheck enumeration: %s\n", err)
	}
}
