package project

import (
	"context"
	"fmt"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/decorator"
	"github.com/hirakiuc/alfred-jira-workflow/model"
	"github.com/hirakiuc/alfred-jira-workflow/resource"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type HelpCommand struct {
	subcommand.BaseCommand
}

func NewHelpCommand(args []string) HelpCommand {
	return HelpCommand{
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd HelpCommand) showProjectList(ctx context.Context, wf *aw.Workflow) {
	r := resource.NewProjectResource(wf)

	projects, err := r.List(ctx)
	if err != nil {
		wf.FatalError(err)

		return
	}

	d, err := decorator.NewProjectDecorator(wf)
	if err != nil {
		wf.FatalError(err)

		return
	}

	for _, project := range projects {
		v := project
		d.SetTarget(&v)

		wf.NewItem(d.GetTitle()).
			Subtitle(d.GetSubtitle()).
			Arg(v.Self).
			Valid(true)
	}

	if cmd.HasQueryWithOffset(1) {
		wf.Filter(cmd.QueryWithOffset(1))
	}

	// Show a warning in Alfred if there are no items.
	wf.WarnEmpty("No projects found.", "")
}

func (cmd HelpCommand) showHelp(_ context.Context, wf *aw.Workflow, prj *model.ProjectData) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "issue",
			desc: "show issues in this project",
		},
	}

	for _, c := range subcommands {
		wf.NewItem(c.name).
			Subtitle(c.desc).
			Autocomplete(fmt.Sprintf("project %s %s ", prj.ID, c.name)).
			Valid(false)
	}

	if len(cmd.Args) > 1 {
		args := cmd.Args[1:]
		wf.Filter(strings.Join(args, " "))
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No items found.", " ")
}

func (cmd HelpCommand) Run(ctx context.Context, wf *aw.Workflow) {
	if len(cmd.Args) == 0 {
		cmd.showProjectList(ctx, wf)

		return
	}

	projectID := cmd.Args[0]

	r := resource.NewProjectResource(wf)

	project, err := r.GetByID(ctx, projectID)
	if err != nil {
		wf.FatalError(err)

		return
	}

	if project == nil {
		// no such project found
		// => show boards list and filter by that string
		cmd.showProjectList(ctx, wf)

		return
	}

	// Found that project
	// => show menu for the project
	cmd.showHelp(ctx, wf, project)
}
