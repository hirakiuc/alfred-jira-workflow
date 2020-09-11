package cache

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

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

func (cache *IssuesCache) getCacheWithKey(key string, cacheAge time.Duration) ([]jira.Issue, error) {
	issues := []jira.Issue{}

	err := cache.getRawCache(key, cacheAge, &issues)
	if err != nil {
		return []jira.Issue{}, err
	}

	if issues == nil {
		return []jira.Issue{}, nil
	}

	return issues, nil
}

func (cache *IssuesCache) storeWithKey(key string, issues []jira.Issue) ([]jira.Issue, error) {
	err := cache.storeRawData(key, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}

// ---------------------------------------------

func (cache *IssuesCache) getCacheKeyWithJQL(jql string) string {
	hash := sha256.Sum256([]byte(strings.TrimSpace(jql)))

	return fmt.Sprintf("issues-jql-%x", hash)
}

func (cache *IssuesCache) GetCacheByJQL(jql string) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKeyWithJQL(jql)

	return cache.getCacheWithKey(cacheKey, getMaxCacheAge())
}

func (cache *IssuesCache) StoreByJQL(jql string, issues []jira.Issue) ([]jira.Issue, error) {
	cacheKey := cache.getCacheKeyWithJQL(jql)

	return cache.storeWithKey(cacheKey, issues)
}

// ---------------------------------------------
