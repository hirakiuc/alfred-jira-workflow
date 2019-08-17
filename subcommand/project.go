package subcommand

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type ProjectCommand struct {
	BaseCommand
}

func NewProjectCommand(args []string) ProjectCommand {
	return ProjectCommand{
		BaseCommand{
			Args: args,
		},
	}
}

func projectTitle(key string, name string) string {
	return fmt.Sprintf("%s - %s", key, name)
}

func projectSubtitle(prjID string, prjSelf string) string {
	return fmt.Sprintf("%s - %s", prjID, prjSelf)
}

func (cmd ProjectCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient()
	if err != nil {
		wf.FatalError(err)
		return
	}

	projects, err := client.GetProjects()
	if err != nil {
		wf.FatalError(err)
		return
	}

	if projects == nil {
		projects = &jira.ProjectList{}
	}

	for _, project := range *projects {
		wf.NewItem(projectTitle(project.Key, project.Name)).
			Subtitle(projectSubtitle(project.ID, project.Self)).
			Arg(project.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No projects found.", "")
}
