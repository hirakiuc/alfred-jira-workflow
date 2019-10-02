package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

type BoardSprintCommand struct {
	BoardName string
	BaseCommand
}

func NewBoardSprintCommand(name string, args []string) BoardSprintCommand {
	return BoardSprintCommand{
		BoardName: name,
		BaseCommand: BaseCommand{
			Args: args,
		},
	}
}

func (cmd BoardSprintCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	wf.NewItem("TBD show sprints in this board.").
		Valid(false)
}
