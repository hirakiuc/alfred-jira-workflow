package subcommand

import (
	"context"
	"errors"
	"strconv"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
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

func (cmd BoardSprintCommand) getBoard(ctx context.Context, wf *aw.Workflow) (*jira.Board, error) {
	r := resource.NewBoardResource(wf)

	// board, err := client.GetBoardByName(cmd.BoardName)
	// TBD: temporary, fetch By BoardID
	boardID, err := strconv.Atoi(cmd.BoardName)
	if err != nil {
		return nil, err
	}

	board, err := r.GetByID(ctx, boardID)
	if err != nil {
		return nil, err
	}
	if board == nil {
		return nil, errors.New("no such board found")
	}

	return board, nil
}

func (cmd BoardSprintCommand) Run(ctx context.Context, wf *aw.Workflow) {
	board, err := cmd.getBoard(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// fetch sprints in the board
	r := resource.NewSprintResource(wf)
	sprints, err := r.GetAllByBoardID(ctx, board.ID)
	if err != nil {
		wf.FatalError(err)
		return
	}

	d, err := decorator.NewSprintDecorator(wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, sprint := range sprints {
		v := sprint
		d.SetTarget(&v)

		wf.NewItem(d.Title()).
			Subtitle(d.SubTitle()).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items.
	wf.WarnEmpty("No sprints found.", "")
}
