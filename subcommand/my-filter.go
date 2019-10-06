package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
)

type MyFilterCommand struct {
	BaseCommand
}

func NewMyFilterCommand(args []string) MyFilterCommand {
	return MyFilterCommand{
		BaseCommand{
			Args: args,
		},
	}
}

func (cmd MyFilterCommand) Run(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewFilterResource(wf)
	filters, err := r.MyFilters(ctx)
	if err != nil {
		wf.FatalError(err)
		return
	}

	d, err := decorator.NewFilterDecorator(wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, filter := range filters {
		v := filter
		d.SetTarget(&v)

		wf.NewItem(d.Title()).
			Subtitle(d.Subtitle()).
			Arg(filter.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in alfred if there are no filters.
	wf.WarnEmpty("No filters found.", "")
}
