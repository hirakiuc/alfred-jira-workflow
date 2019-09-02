package api

import (
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) SearchIssuesByFilter(filterID string, orderBy string) ([]jira.Issue, error) {
	jql := fmt.Sprintf("Filter = %s Order By %s", filterID, orderBy)
	return client.SearchIssuesByJQL(jql)
}

func (client *Client) SearchIssuesByJQL(jql string) ([]jira.Issue, error) {
	searchOption := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 200,
	}

	fmt.Fprintf(os.Stderr, "jql:%s\n", jql)
	issues, _, err := client.jira.Issue.Search(jql, &searchOption)
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
