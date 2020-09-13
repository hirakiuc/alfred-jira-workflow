package api

import (
	"context"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) GetProjectByID(_ context.Context, prjID string) (*jira.Project, error) {
	project, _, err := client.jira.Project.Get(prjID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (client *Client) GetProjects() (*jira.ProjectList, error) {
	opts := jira.GetQueryOptions{}

	list, _, err := client.jira.Project.ListWithOptions(&opts)
	if err != nil {
		return list, err
	}

	return list, nil
}
