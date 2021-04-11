package issue

import (
	"context"
	"fmt"
	"os"

	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type Command struct {
	subcommand.BaseCommand
}

func NewCommand(args []string) Command {
	return Command{
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func splitArgs(args []string) ([]string, string) {
	for _, w := range args {
		fmt.Fprintf(os.Stderr, "word:%s\n", w)
	}

	l := len(args)
	switch l {
	case 0:
		return []string{}, ""
	case 1:
		return args, ""
	default:
		return args[0:(l - 1)], args[l-1]
	}
}

func (cmd Command) Run(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewIssueResource(wf)

	opts, word := splitArgs(cmd.Args)
	for _, w := range opts {
		fmt.Fprintf(os.Stderr, "w:%s\n", w)
	}

	issues, err := r.SearchIssues(ctx, opts)
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

	if len(word) > 0 {
		wf.Filter(word)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
