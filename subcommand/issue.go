package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
)

type IssueCommand struct {
	BaseCommand
}

func NewIssueCommand(args []string) IssueCommand {
	return IssueCommand{
		BaseCommand{
			Args: args,
		},
	}
}

func (cmd IssueCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient()
	if err != nil {
		wf.FatalError(err)
		return
	}

	issues, err := client.SearchIssues(cmd.Arguments())
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

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
