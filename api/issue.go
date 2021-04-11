package api

import (
	"context"

	"github.com/andygrunwald/go-jira"
)

const (
	IssuesPerRequests int = 200
)

func (client *Client) SearchIssues(ctx context.Context, jql string) ([]jira.Issue, error) {
	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: IssuesPerRequests,
	}

	issues, _, err := client.jira.Issue.SearchWithContext(ctx, jql, &opts)
	if err != nil {
		return []jira.Issue{}, err
	}

	return issues, nil
}

func (client *Client) GetIssue(ctx context.Context, issueID string) (*jira.Issue, error) {
	opts := jira.GetQueryOptions{}

	issue, _, err := client.jira.Issue.GetWithContext(ctx, issueID, &opts)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
