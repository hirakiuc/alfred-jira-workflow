package api

import (
	"context"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) GetProjectByID(ctx context.Context, prjID string) (*jira.Project, error) {
	project, _, err := client.jira.Project.GetWithContext(ctx, prjID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (client *Client) GetProjects(ctx context.Context) (*jira.ProjectList, error) {
	opts := jira.GetQueryOptions{}

	list, _, err := client.jira.Project.ListWithOptionsWithContext(ctx, &opts)
	if err != nil {
		return list, err
	}

	return list, nil
}
