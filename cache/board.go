package cache

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
)

type BoardsCache struct {
	*BaseCache
}

func NewBoardsCache(wf *aw.Workflow) *BoardsCache {
	return &BoardsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *BoardsCache) getCacheKey() string {
	return fmt.Sprintf("boards")
}

func (cache *BoardsCache) GetCache() ([]jira.Board, error) {
	cacheKey := cache.getCacheKey()

	boards := []jira.Board{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &boards)
	if err != nil {
		return []jira.Board{}, err
	}
	if boards == nil {
		return []jira.Board{}, nil
	}

	return boards, nil
}

func (cache *BoardsCache) Store(boards []jira.Board) ([]jira.Board, error) {
	cacheKey := cache.getCacheKey()

	_, err := cache.storeRawData(cacheKey, boards)
	if err != nil {
		return boards, err
	}

	return boards, nil
}
