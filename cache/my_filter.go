package cache

import (
	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

type MyFiltersCache struct {
	*BaseCache
}

func NewMyFiltersCache(wf *aw.Workflow) *MyFiltersCache {
	return &MyFiltersCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *MyFiltersCache) getCacheKey() string {
	return "my-filters"
}

func (cache *MyFiltersCache) GetCache() ([]jira.Filter, error) {
	cacheKey := cache.getCacheKey()

	filters := []jira.Filter{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &filters)
	if err != nil {
		return []jira.Filter{}, err
	}
	if filters == nil {
		return []jira.Filter{}, nil
	}

	return filters, nil
}

func (cache *MyFiltersCache) Store(filters []jira.Filter) ([]jira.Filter, error) {
	cacheKey := cache.getCacheKey()

	_, err := cache.storeRawData(cacheKey, filters)
	if err != nil {
		return filters, err
	}

	return filters, nil
}
