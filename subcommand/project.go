package subcommand

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
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

func (cmd ProjectCommand) fetchProjectList(_ context.Context, wf *aw.Workflow) (jira.ProjectList, error) {
	store := cache.NewProjectListCache(wf)

	list, err := store.GetCache()
	if err != nil {
		return jira.ProjectList{}, err
	}

	if len(list) != 0 {
		return list, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return jira.ProjectList{}, err
	}

	result, err := client.GetProjects()
	if err != nil {
		return jira.ProjectList{}, err
	}

	if result == nil {
		return jira.ProjectList{}, nil
	}

	return store.Store(*result)
}

func (cmd ProjectCommand) Run(ctx context.Context, wf *aw.Workflow) {
	list, err := cmd.fetchProjectList(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, prj := range list {
		wf.NewItem(projectTitle(prj.Key, prj.Name)).
			Subtitle(projectSubtitle(prj.ID, prj.Self)).
			Arg(prj.Self).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No projects found.", "")
}
