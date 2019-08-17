package cache

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

// IssuesCache describe an instance of issue cache.
type IssuesCache struct {
	*BaseCache
}

func NewIssuesCache(wf *aw.Workflow) *IssuesCache {
	return &IssuesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *IssuesCache) getCacheKey(jql string) string {
	return fmt.Sprintf("issues-%s", jql)
}

func (cache *IssuesCache) GetCache(jql string) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKey(jql)

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

func (cache *IssuesCache) Store(jql string, issues []jira.Issue) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKey(jql)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
