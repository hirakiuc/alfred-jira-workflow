package model

import (
	"github.com/andygrunwald/go-jira"
)

// ProjectData describe an element of jira.ProjectList.
type ProjectData struct {
	Expand          string               `json:"expand" structs:"expand"`
	Self            string               `json:"self" structs:"self"`
	ID              string               `json:"id" structs:"id"`
	Key             string               `json:"key" structs:"key"`
	Name            string               `json:"name" structs:"name"`
	AvatarUrls      jira.AvatarUrls      `json:"avatarUrls" structs:"avatarUrls"`
	ProjectTypeKey  string               `json:"projectTypeKey" structs:"projectTypeKey"`
	ProjectCategory jira.ProjectCategory `json:"projectCategory,omitempty" structs:"projectsCategory,omitempty"`
	IssueTypes      []jira.IssueType     `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
}

func ConvertProjectList(list jira.ProjectList) []ProjectData {
	projects := make([]ProjectData, len(list))

	for i, prj := range list {
		projects[i] = ProjectData{
			Expand:     prj.Expand,
			Self:       prj.Self,
			ID:         prj.ID,
			Key:        prj.Key,
			Name:       prj.Name,
			AvatarUrls: prj.AvatarUrls,
			// An element of ProjectList has the ProjectTypeKey, but jira.Project does not...
			//			ProjectTypeKey:  prj.ProjectTypeKey,
			ProjectCategory: prj.ProjectCategory,
			IssueTypes:      prj.IssueTypes,
		}
	}

	return projects
}
