package subcommand

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
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

func (cmd MyFilterCommand) fetchMyFilters(_ context.Context, wf *aw.Workflow) ([]jira.Filter, error) {
	store := cache.NewMyFiltersCache(wf)

	filters, err := store.GetCache()
	if err != nil {
		return []jira.Filter{}, err
	}
	if len(filters) != 0 {
		return filters, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []jira.Filter{}, err
	}

	filters, err = client.MyFilters()
	if err != nil {
		return []jira.Filter{}, err
	}
	if len(filters) == 0 {
		return []jira.Filter{}, nil
	}

	return store.Store(filters)
}

func (cmd MyFilterCommand) Run(ctx context.Context, wf *aw.Workflow) {
	filters, err := cmd.fetchMyFilters(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, filter := range filters {
		wf.NewItem(filter.Name).
			Subtitle(filter.Description).
			Arg(filter.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in alfred if there are no filters.
	wf.WarnEmpty("No filters found.", "")
}
