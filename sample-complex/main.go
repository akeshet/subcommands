// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// sample - Sample app to demonstrate example usage of package subcommand.
package main

import (
	"github.com/maruel/subcommands"
	"github.com/maruel/subcommands/subcommandstest"
	"log"
	"os"
)

var application = &subcommands.DefaultApplication{
	Name:  "sample",
	Title: "Sample tool to act as a skeleton for subcommands usage.",
	// Commands will be shown in this exact order, so you'll likely want to put
	// them in alphabetical order or in logical grouping.
	Commands: []*subcommands.Command{
		cmdGreet,
		subcommands.CmdHelp,
		cmdAsk,
		cmdSleep,
	},
}

type SampleApplication interface {
	subcommands.Application

	// Add anything desired, in particular if you'd like to crete a fake
	// application during testing.
}

type sample struct {
	*subcommands.DefaultApplication
}

func main() {
	subcommands.KillStdLog()
	application := &subcommands.DefaultApplication{
		Logger: subcommands.NewLogger(),
		Name:   "sample",
		Title:  "Sample tool to act as a skeleton for subcommands usage.",
		// Commands will be shown in this exact order, so you'll likely want to put
		// them in alphabetical order or in logical grouping.
		Commands: []*subcommands.Command{
			cmdGreet,
			subcommands.CmdHelp,
			cmdAsk,
			cmdSleep,
		},
	}

	application.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// Disable log by default, unless -verbose is specified.
	//d := &sample{application, log.New(nullWriter(0), "", log.LstdFlags|log.Lmicroseconds)}
	d := &sample{application, log.New(application.GetErr(), "")}
	os.Exit(subcommands.Run(d, nil))
}
