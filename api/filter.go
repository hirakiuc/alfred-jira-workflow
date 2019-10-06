package api

import (
	"github.com/andygrunwald/go-jira"
)

func (client *Client) MyFilters() ([]jira.Filter, error) {
	opts := jira.GetMyFiltersQueryOptions{
		IncludeFavourites: true,
	}

	filters, _, err := client.jira.Filter.GetMyFilters(&opts)
	if err != nil {
		return []jira.Filter{}, err
	}

	items := []jira.Filter{}
	for _, filter := range filters {
		items = append(items, *filter)
	}

	return items, nil
}

func (client *Client) SearchFilters() ([]jira.FiltersListItem, error) {
	opts := jira.FilterSearchOptions{}

	items := []jira.FiltersListItem{}

	for {
		list, _, err := client.jira.Filter.Search(&opts)
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

func (client *Client) GetFilterByID(filterID int) (*jira.Filter, error) {
	filter, _, err := client.jira.Filter.Get(filterID)
	if err != nil {
		return nil, err
	}

	return filter, nil
}
