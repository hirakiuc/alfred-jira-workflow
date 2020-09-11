package project

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type Command struct {
	subcommand.BaseCommand
}

func NewCommand(args []string) Command {
	return Command{
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd Command) Run(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewProjectResource(wf)

	list, err := r.List(ctx)
	if err != nil {
		wf.FatalError(err)

		return
	}

	d, err := decorator.NewProjectDecorator(wf)
	if err != nil {
		wf.FatalError(err)

		return
	}

	for _, prj := range list {
		v := prj

		d.SetTarget(&v)

		wf.NewItem(d.GetTitle()).
			Subtitle(d.GetSubtitle()).
			Arg(prj.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No projects found.", "")
}
