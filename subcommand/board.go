package subcommand

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
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

	d, err := decorator.NewBoardDecorator(wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, board := range boards {
		v := board
		d.SetTarget(&v)

		wf.NewItem(d.Title()).
			Subtitle(d.Subtitle()).
			Autocomplete(fmt.Sprintf("board %s ", v.Name)).
			Arg(v.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No boards found.", "")
}
