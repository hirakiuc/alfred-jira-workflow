package subcommand

import (
	"context"
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
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

func (cmd BoardCommand) Run(ctx context.Context, wf *aw.Workflow) {
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

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No boards found.", "")
}
