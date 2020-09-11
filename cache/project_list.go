package cache

import (
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/model"
)

type ProjectListCache struct {
	*BaseCache
}

func NewProjectListCache(wf *aw.Workflow) *ProjectListCache {
	return &ProjectListCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *ProjectListCache) getCacheKey() string {
	return "project-list"
}

func (cache *ProjectListCache) GetCache() ([]model.ProjectData, error) {
	cacheKey := cache.getCacheKey()

	list := []model.ProjectData{}

	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &list)
	if err != nil {
		return []model.ProjectData{}, err
	}

	if list == nil {
		return []model.ProjectData{}, nil
	}

	return list, nil
}

func (cache *ProjectListCache) Store(list []model.ProjectData) ([]model.ProjectData, error) {
	cacheKey := cache.getCacheKey()

	err := cache.storeRawData(cacheKey, list)
	if err != nil {
		return list, err
	}

	return list, nil
}
