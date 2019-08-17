package subcommand

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type BoardCommand struct {
	BaseCommand
}

func NewBoardCommand(args []string) BoardCommand {
	return BoardCommand{
		BaseCommand{
			Args: args,
		},
	}
}

func boardSubTitle(board jira.Board) string {
	return fmt.Sprintf("%d - %s", board.ID, board.Self)
}

func (cmd BoardCommand) fetchBoards(_ context.Context, wf *aw.Workflow) ([]jira.Board, error) {
	store := cache.NewBoardsCache(wf)

	boards, err := store.GetCache()
	if err != nil {
		return []jira.Board{}, err
	}
	if len(boards) != 0 {
		return boards, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []jira.Board{}, err
	}

	boards, err = client.GetAllBoards("", "")
	if err != nil {
		return []jira.Board{}, err
	}

	if len(boards) == 0 {
		return []jira.Board{}, nil
	}

	return store.Store(boards)
}

func (cmd BoardCommand) Run(ctx context.Context, wf *aw.Workflow) {
	boards, err := cmd.fetchBoards(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, board := range boards {
		wf.NewItem(board.Name).
			Subtitle(boardSubTitle(board)).
			Arg(board.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No boards found.", "")
}
