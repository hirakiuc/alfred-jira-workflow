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

func ConvertProject(prj *jira.Project) *ProjectData {
	return &ProjectData{
		Expand:     prj.Expand,
		Self:       prj.Self,
		ID:         prj.ID,
		Key:        prj.Key,
		Name:       prj.Name,
		AvatarUrls: prj.AvatarUrls,
		// Ignore this prop because jira.Project does not have ProjectTypeKey property.
		// ProjectTypeKey:  prj.ProjectTypeKey,
		ProjectCategory: prj.ProjectCategory,
		IssueTypes:      prj.IssueTypes,
	}
}

// For sort function

type ProjectDataList []ProjectData

func (list ProjectDataList) Len() int {
	return len(list)
}

func (list ProjectDataList) Less(i int, j int) bool {
	return list[i].ID < list[j].ID
}

func (list ProjectDataList) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}
