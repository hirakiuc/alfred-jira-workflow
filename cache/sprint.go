package cache

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

type SprintsCache struct {
	*BaseCache
}

func NewSprintsCache(wf *aw.Workflow) *SprintsCache {
	return &SprintsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *SprintsCache) getCacheKey(boardID int) string {
	return fmt.Sprintf("sprints-board-%d", boardID)
}

func (cache *SprintsCache) GetCache(boardID int) ([]jira.Sprint, error) {
	cacheKey := cache.getCacheKey(boardID)

	sprints := []jira.Sprint{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &sprints)
	if err != nil {
		return []jira.Sprint{}, err
	}
	if sprints == nil {
		return []jira.Sprint{}, nil
	}

	return sprints, nil
}

func (cache *SprintsCache) Store(sprints []jira.Sprint, boardID int) ([]jira.Sprint, error) {
	cacheKey := cache.getCacheKey(boardID)

	_, err := cache.storeRawData(cacheKey, sprints)
	if err != nil {
		return sprints, err
	}

	return sprints, nil
}
