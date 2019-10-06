package resource

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type FilterResource struct {
	wf *aw.Workflow
}

func NewFilterResource(wf *aw.Workflow) *FilterResource {
	return &FilterResource{
		wf: wf,
	}
}

func (r *FilterResource) MyFilters(_ context.Context) ([]jira.Filter, error) {
	store := cache.NewMyFiltersCache(r.wf)
	filters, err := store.GetCache()
	if err != nil {
		return []jira.Filter{}, err
	}
	if len(filters) > 0 {
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
