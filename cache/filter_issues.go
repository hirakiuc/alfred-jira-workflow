package cache

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

type FilterIssuesCache struct {
	*BaseCache
}

func NewFilterIssuesCache(wf *aw.Workflow) *FilterIssuesCache {
	return &FilterIssuesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *FilterIssuesCache) getCacheKey(filterID string) string {
	return fmt.Sprintf("filtered-issues-%s", filterID)
}

func (cache *FilterIssuesCache) GetCache(filterID string) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKey(filterID)

	issues := []jira.Issue{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &issues)
	if err != nil {
		return []jira.Issue{}, err
	}
	if issues == nil {
		return []jira.Issue{}, nil
	}

	return issues, nil
}

func (cache *FilterIssuesCache) Store(issues []jira.Issue, filterID string) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKey(filterID)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
