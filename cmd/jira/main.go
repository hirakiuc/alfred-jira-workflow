package main

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/cli"
)

// Your workflow starts here.
func run(wf *aw.Workflow) func() {
	return func() {
		cmd, err := cli.ParseArgs(wf.Args())
		if err != nil {
			wf.FatalError(err)

			return
		}

		ctx := context.Background()

		(*cmd).Run(ctx, wf)

		// Add a "Script Filter" result
		// wf.NewItem("First result!")

		// Send results to Alfred
		wf.SendFeedback()
	}
}

func main() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables.
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf := aw.New()

	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run(wf))
}
