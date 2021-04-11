package resource

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type SprintResource struct {
	wf *aw.Workflow
}

func NewSprintResource(wf *aw.Workflow) *SprintResource {
	return &SprintResource{
		wf: wf,
	}
}

func (r *SprintResource) GetAllByBoardID(_ context.Context, boardID int) ([]jira.Sprint, error) {
	store := cache.NewSprintsCache(r.wf)

	sprints, err := store.GetCache(boardID)
	if err != nil {
		return []jira.Sprint{}, err
	}

	if len(sprints) != 0 {
		return sprints, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []jira.Sprint{}, err
	}

	sprints, err = client.GetSprintsInBoard(boardID)
	if err != nil {
		return []jira.Sprint{}, err
	}

	if len(sprints) == 0 {
		return []jira.Sprint{}, nil
	}

	return store.Store(sprints, boardID)
}
