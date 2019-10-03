package resource

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type BoardResource struct {
	wf *aw.Workflow
}

func NewBoardResource(wf *aw.Workflow) *BoardResource {
	return &BoardResource{
		wf: wf,
	}
}

func (r *BoardResource) GetAll(_ context.Context, projectKeyOrID string, name string) ([]jira.Board, error) {
	store := cache.NewBoardsCache(r.wf)

	boards, err := store.GetCache()
	if err != nil {
		return []jira.Board{}, err
	}
	if len(boards) != 0 {
		return boards, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []jira.Board{}, err
	}

	boards, err = client.GetAllBoards(projectKeyOrID, name)
	if err != nil {
		return []jira.Board{}, err
	}
	if len(boards) == 0 {
		return []jira.Board{}, nil
	}

	return store.Store(boards)
}

func (r *BoardResource) GetByID(_ context.Context, boardID int) (*jira.Board, error) {
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	board, err := client.GetBoardByID(boardID)
	if err != nil {
		return nil, err
	}

	return board, nil
}
