package project

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type IssueCommand struct {
	prjID string
	subcommand.BaseCommand
}

func NewIssueCommand(prjID string, args []string) IssueCommand {
	return IssueCommand{
		prjID: prjID,
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd IssueCommand) Run(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewIssueResource(wf)

	word := strings.Join(cmd.Args, " ")

	// Fetch issues
	issues, err := r.ListByProjectID(ctx, cmd.prjID)
	if err != nil {
		wf.FatalError(err)

		return
	}

	d, err := decorator.NewIssueDecorator(wf)
	if err != nil {
		wf.FatalError(err)

		return
	}

	for _, issue := range issues {
		v := issue
		d.SetTarget(&v)

		wf.NewItem(d.Title()).
			Arg(d.URL()).
			Valid(true)
	}

	if len(word) > 0 {
		wf.Filter(word)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
