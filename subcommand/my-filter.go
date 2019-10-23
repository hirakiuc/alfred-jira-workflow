package subcommand

import (
	"context"
	"strconv"
	"strings"

	"github.com/andygrunwald/go-jira"
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

func (cmd MyFilterCommand) showMyFilters(ctx context.Context, wf *aw.Workflow, opts []string) {
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

	if len(opts) > 0 {
		wf.Filter(strings.Join(opts, " "))
	}

	// Show a warning in alfred if there are no filters.
	wf.WarnEmpty("No filters found.", "")
}

func (cmd MyFilterCommand) showFilteredIssues(_ context.Context, wf *aw.Workflow, filter *jira.Filter, opts []string) {
	r := resource.NewIssueResource(wf)

	issues, err := r.SearchIssues([]string{filter.Jql})
	if err != nil {
		wf.FatalError(err)
		return
	}

	d, err := decorator.NewIssueDecorator(wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, issue := range issues {
		v := issue
		d.SetTarget(&v)

		wf.NewItem(d.Title()).
			Arg(d.URL()).
			Valid(true)
	}

	if len(opts) > 0 {
		wf.Filter(strings.Join(opts, " "))
	}

	// Show a warning in alfred if there are no filters.
	wf.WarnEmpty("No issues found.", "")
}

func (cmd MyFilterCommand) Run(ctx context.Context, wf *aw.Workflow) {
	if len(cmd.Args) == 0 {
		// case:my-filter
		cmd.showMyFilters(ctx, wf, cmd.Args)
		return
	}

	firstArg := cmd.Args[0]

	opts := []string{}
	if len(cmd.Args) > 1 {
		opts = cmd.Args[1:]
	}

	filterID, err := strconv.Atoi(firstArg)
	if err != nil {
		// case:my-filter [word] ...
		cmd.showMyFilters(ctx, wf, cmd.Args)
		return
	}

	r := resource.NewFilterResource(wf)

	filter, err := r.GetFilterByID(filterID)
	if err != nil || filter == nil {
		// case:my-filter [word]... No such filter found
		cmd.showMyFilters(ctx, wf, cmd.Args)
		return
	}

	// Show filtered issues
	cmd.showFilteredIssues(ctx, wf, filter, opts)
}
