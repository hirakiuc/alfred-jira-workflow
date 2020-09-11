package cache

import (
	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
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

func (cache *ProjectListCache) GetCache() (jira.ProjectList, error) {
	cacheKey := cache.getCacheKey()

	list := jira.ProjectList{}

	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &list)
	if err != nil {
		return jira.ProjectList{}, err
	}

	if list == nil {
		return jira.ProjectList{}, nil
	}

	return list, nil
}

func (cache *ProjectListCache) Store(list jira.ProjectList) (jira.ProjectList, error) {
	cacheKey := cache.getCacheKey()

	err := cache.storeRawData(cacheKey, list)
	if err != nil {
		return list, err
	}

	return list, nil
}
