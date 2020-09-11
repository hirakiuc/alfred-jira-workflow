package resource

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type FilterResource struct{ wf *aw.Workflow }

func NewFilterResource(wf *aw.Workflow) *FilterResource {
	return &FilterResource{
		wf: wf,
	}
}

func (r *FilterResource) MyFilters(_ context.Context) ([]jira.Filter, error) {
	store := cache.NewFiltersCache(r.wf)

	filters, err := store.GetFiltersCache()
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

	return store.StoreFilters(filters)
}

func (r *FilterResource) GetFilterByID(filterID int) (*jira.Filter, error) {
	store := cache.NewFiltersCache(r.wf)

	filter, err := store.GetFilterCache(filterID)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		return filter, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	filter, err = client.GetFilterByID(filterID)
	if err != nil {
		return nil, err
	}

	if filter == nil {
		return nil, nil
	}

	return store.StoreFilter(filter)
}
