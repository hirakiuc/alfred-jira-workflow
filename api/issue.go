package api

import (
	"github.com/andygrunwald/go-jira"
)

func (client *Client) SearchIssues(jql string) ([]jira.Issue, error) {
	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 200,
	}

	issues, _, err := client.jira.Issue.Search(jql, &opts)
	if err != nil {
		return []jira.Issue{}, err
	}

	return issues, nil
}

func (client *Client) GetIssue(issueID string) (*jira.Issue, error) {
	opts := jira.GetQueryOptions{}

	issue, _, err := client.jira.Issue.Get(issueID, &opts)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
