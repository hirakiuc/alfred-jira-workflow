package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

type BoardBacklogCommand struct {
	BoardName string
	BaseCommand
}

func NewBoardBacklogCommand(name string, args []string) BoardBacklogCommand {
	return BoardBacklogCommand{
		BoardName: name,
		BaseCommand: BaseCommand{
			Args: args,
		},
	}
}

func (cmd BoardBacklogCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	wf.NewItem("TBD show backlog in this board").
		Valid(false)
	/*
		TBD
		// Fetch the board
		board, err := fetchBoardByName(wf, cmd.BoardName)
		if err != nil {
			wf.FatalError(err)
			return
		}

		// Fetch tickets in the backlog of this board
		issues := []*jira.Issue{}

		for _, issue := range issues {
			wf.NewItem()
		}
	*/
}
