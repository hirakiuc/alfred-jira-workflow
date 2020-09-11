package resource

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
	"github.com/hirakiuc/alfred-jira-workflow/model"
)

type ProjectResource struct {
	wf *aw.Workflow
}

func NewProjectResource(wf *aw.Workflow) *ProjectResource {
	return &ProjectResource{
		wf: wf,
	}
}

func (r *ProjectResource) List(ctx context.Context) ([]model.ProjectData, error) {
	store := cache.NewProjectListCache(r.wf)

	list, err := store.GetCache()
	if err != nil {
		return []model.ProjectData{}, err
	}

	if len(list) != 0 {
		return list, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []model.ProjectData{}, err
	}

	result, err := client.GetProjects()
	if err != nil {
		return []model.ProjectData{}, err
	}

	if result == nil {
		return []model.ProjectData{}, nil
	}

	projects := model.ConvertProjectList(*result)

	return store.Store(projects)
}
