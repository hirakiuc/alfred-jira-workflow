package cache

import (
	"fmt"
	"strconv"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

type FiltersCache struct {
	*BaseCache
}

func NewFiltersCache(wf *aw.Workflow) *FiltersCache {
	return &FiltersCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *FiltersCache) getCacheWithKey(cacheKey string) ([]jira.Filter, error) {
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

func (cache *FiltersCache) storeWithKey(cacheKey string, filters []jira.Filter) ([]jira.Filter, error) {
	err := cache.storeRawData(cacheKey, filters)
	if err != nil {
		return filters, err
	}

	return filters, nil
}

//----------------------------------------------------------

func (cache *FiltersCache) getCacheKey() string {
	return "filters"
}

func (cache *FiltersCache) GetFiltersCache() ([]jira.Filter, error) {
	cacheKey := cache.getCacheKey()

	return cache.getCacheWithKey(cacheKey)
}

func (cache *FiltersCache) StoreFilters(filters []jira.Filter) ([]jira.Filter, error) {
	cacheKey := cache.getCacheKey()

	return cache.storeWithKey(cacheKey, filters)
}

//----------------------------------------------------------

func (cache *FiltersCache) getCacheKeyWithID(filterID int) string {
	return fmt.Sprintf("filter-%d", filterID)
}

func (cache *FiltersCache) GetFilterCache(filterID int) (*jira.Filter, error) {
	cacheKey := cache.getCacheKeyWithID(filterID)

	filters, err := cache.getCacheWithKey(cacheKey)
	if err != nil {
		return nil, err
	}

	if len(filters) == 0 {
		return nil, nil
	}

	return &filters[0], nil
}

func (cache *FiltersCache) StoreFilter(filter *jira.Filter) (*jira.Filter, error) {
	strID, err := strconv.Atoi(filter.ID)
	if err != nil {
		return nil, err
	}

	cacheKey := cache.getCacheKeyWithID(strID)

	filters, err := cache.storeWithKey(cacheKey, []jira.Filter{*filter})
	if err != nil {
		return nil, err
	}

	return &filters[0], nil
}
