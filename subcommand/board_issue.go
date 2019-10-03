package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

type BoardIssueCommand struct {
	BoardName string
	BaseCommand
}

func NewBoardIssueCommand(name string, args []string) BoardIssueCommand {
	return BoardIssueCommand{
		BoardName: name,
		BaseCommand: BaseCommand{
			Args: args,
		},
	}
}

func (cmd BoardIssueCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	wf.NewItem("TBD show issues in this board").
		Valid(false)
}
