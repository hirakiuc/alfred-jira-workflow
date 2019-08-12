package api

import "github.com/andygrunwald/go-jira"

func (client *Client) GetProjects() (*jira.ProjectList, error) {
	opts := jira.GetQueryOptions{}

	list, _, err := client.jira.Project.ListWithOptions(&opts)
	if err != nil {
		return list, err
	}

	return list, nil
}
