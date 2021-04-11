package api

import (
	"context"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) MyFilters(ctx context.Context) ([]jira.Filter, error) {
	opts := jira.GetMyFiltersQueryOptions{
		IncludeFavourites: true,
	}

	filters, _, err := client.jira.Filter.GetMyFiltersWithContext(ctx, &opts)
	if err != nil {
		return []jira.Filter{}, err
	}

	items := []jira.Filter{}
	for _, filter := range filters {
		items = append(items, *filter)
	}

	return items, nil
}

func (client *Client) SearchFilters(ctx context.Context) ([]jira.FiltersListItem, error) {
	opts := jira.FilterSearchOptions{}

	items := []jira.FiltersListItem{}

	for {
		list, _, err := client.jira.Filter.SearchWithContext(ctx, &opts)
		if err != nil {
			return items, err
		}

		items = append(items, list.Values...)

		if list.IsLast {
			return items, nil
		}

		// next page
		opts.StartAt += int64(len(list.Values))
	}
}

func (client *Client) GetFilterByID(ctx context.Context, filterID int) (*jira.Filter, error) {
	filter, _, err := client.jira.Filter.GetWithContext(ctx, filterID)
	if err != nil {
		return nil, err
	}

	return filter, nil
}
