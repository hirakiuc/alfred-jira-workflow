package cache

import (
	"sort"

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

func (c *ProjectListCache) getCacheKey() string {
	return "project-list"
}

func (c *ProjectListCache) GetCache() ([]model.ProjectData, error) {
	cacheKey := c.getCacheKey()

	list := []model.ProjectData{}

	err := c.getRawCache(cacheKey, getMaxCacheAge(), &list)
	if err != nil {
		return []model.ProjectData{}, err
	}

	if list == nil {
		return []model.ProjectData{}, nil
	}

	return list, nil
}

func (c *ProjectListCache) Store(list []model.ProjectData) ([]model.ProjectData, error) {
	cacheKey := c.getCacheKey()

	sort.Sort(model.ProjectDataList(list))

	err := c.storeRawData(cacheKey, list)
	if err != nil {
		return list, err
	}

	return list, nil
}
