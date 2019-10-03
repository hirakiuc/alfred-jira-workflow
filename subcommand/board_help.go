package subcommand

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
)

// BoardHelpCommand describe a command to show the menus on board subcommands.
type BoardHelpCommand struct {
	BaseCommand
}

func NewBoardHelpCommand(args []string) BoardHelpCommand {
	return BoardHelpCommand{
		BaseCommand: BaseCommand{
			Args: args,
		},
	}
}

func (cmd BoardHelpCommand) showBoardLists(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewBoardResource(wf)
	boards, err := r.GetAll(ctx, "", "")
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

	// Show a warning in Alfred if there are no items.
	wf.WarnEmpty("No boards found.", "")
}

func (cmd BoardHelpCommand) showHelpMenus(_ context.Context, wf *aw.Workflow, board *jira.Board) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "sprint",
			desc: "show sprints in this board",
		},
		{
			name: "backlog",
			desc: "show tickets in backlog of this board",
		},
		{
			name: "issue",
			desc: "show tickets in this board",
		},
	}

	for _, c := range subcommands {
		wf.NewItem(c.name).
			Subtitle(c.desc).
			Autocomplete(fmt.Sprintf("board %s %s ", board.Name, c.name)).
			Valid(false)
	}

	// Args: [BoardID, ...]
	if len(cmd.Args) > 1 {
		args := cmd.Args[1:]
		wf.Filter(strings.Join(args, " "))
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No items found.", "")
}

func (cmd BoardHelpCommand) Run(ctx context.Context, wf *aw.Workflow) {
	// boards {boardID}, [arg, ...]
	// If boardID is the numeric string & such board found, show help
	if len(cmd.Args) == 0 {
		// => show boards if no arguments are gives
		cmd.showBoardLists(ctx, wf)
		return
	}

	boardID, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		// looks like filter string
		// => show boards list and filter by that string
		cmd.showBoardLists(ctx, wf)
		return
	}

	r := resource.NewBoardResource(wf)
	board, err := r.GetByID(ctx, boardID)
	if err != nil {
		wf.FatalError(err)
		return
	}
	if board == nil {
		// no such board found
		// => show boards list and filter by that string
		cmd.showBoardLists(ctx, wf)
		return
	}

	// Found that board
	// => show menus for the board
	cmd.showHelpMenus(ctx, wf, board)
}
