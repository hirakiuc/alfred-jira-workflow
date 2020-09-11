package board

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type IssueCommand struct {
	BoardName string
	subcommand.BaseCommand
}

func NewIssueCommand(name string, args []string) IssueCommand {
	return IssueCommand{
		BoardName: name,
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd IssueCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	wf.NewItem("TBD show issues in this board").
		Valid(false)
}
