package subcommand

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
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

func (cmd BoardSprintCommand) fetchSprints(
	_ context.Context, wf *aw.Workflow, client *api.Client, boardID int) (
	[]jira.Sprint, error) {
	store := cache.NewSprintsCache(wf)

	sprints, err := store.GetCache(boardID)
	if err != nil {
		return []jira.Sprint{}, err
	}
	if len(sprints) != 0 {
		return sprints, nil
	}

	sprints, err = client.GetSprintsInBoard(boardID)
	if err != nil {
		return []jira.Sprint{}, err
	}
	if len(sprints) == 0 {
		return []jira.Sprint{}, nil
	}

	return store.Store(sprints, boardID)
}

func (cmd BoardSprintCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient()
	if err != nil {
		wf.FatalError(err)
		return
	}

	fmt.Fprintf(os.Stderr, "name:%s", cmd.BoardName)

	//	board, err := client.GetBoardByName(cmd.BoardName)
	// TBD: temporary, fetch By BoardID
	boardID, err := strconv.Atoi(cmd.BoardName)
	if err != nil {
		wf.FatalError(err)
		return
	}

	board, err := client.GetBoardByID(boardID)
	if err != nil {
		wf.FatalError(err)
		return
	}
	if board == nil {
		wf.FatalError(errors.New("no such board found"))
		return
	}

	// fetch sprints in the board
	sprints, err := cmd.fetchSprints(_ctx, wf, client, board.ID)
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
