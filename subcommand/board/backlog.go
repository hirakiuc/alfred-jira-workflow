package board

import (
	"context"

	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type BacklogCommand struct {
	BoardName string
	subcommand.BaseCommand
}

func NewBacklogCommand(name string, args []string) BacklogCommand {
	return BacklogCommand{
		BoardName: name,
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

// nolint:wsl // not implemented yet
func (cmd BacklogCommand) Run(_ctx context.Context, wf *aw.Workflow) {
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

	wf.NewItem("TBD show backlog in this board").
		Valid(false)
}
