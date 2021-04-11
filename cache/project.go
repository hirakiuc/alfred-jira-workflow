package cache

import (
	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/model"
)

type ProjectCache struct {
	*BaseCache
}

func NewProjectCache(wf *aw.Workflow) *ProjectCache {
	return &ProjectCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (c *ProjectCache) GetCache(prjID string) (*model.ProjectData, error) {
	listCache := NewProjectListCache(c.wf)

	list, err := listCache.GetCache()
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	for _, prj := range list {
		if prj.ID == prjID {
			return &prj, nil
		}
	}

	return nil, nil
}

func (c *ProjectCache) mergeProjects(projects []model.ProjectData, prj model.ProjectData) []model.ProjectData {
	for idx, v := range projects {
		if v.ID == prj.ID {
			projects[idx] = prj

			return projects
		}
	}

	projects = append(projects, prj)

	return projects
}

func (c *ProjectCache) Store(prj *model.ProjectData) (*model.ProjectData, error) {
	listCache := NewProjectListCache(c.wf)

	projects, err := listCache.GetCache()
	if err != nil {
		return nil, err
	}

	projects = c.mergeProjects(projects, *prj)

	_, err = listCache.Store(projects)
	if err != nil {
		return nil, err
	}

	return prj, nil
}
