package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) SearchIssues(args []string) ([]jira.Issue, error) {
	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 200,
	}

	jql := args2jql(args)
	fmt.Fprintf(os.Stderr, "jql:%s\n", jql)
	issues, _, err := client.jira.Issue.Search(jql, &opts)
	if err != nil {
		return []jira.Issue{}, err
	}

	return issues, nil
}

func args2jql(args []string) string {
	return strings.Join(args, " AND ")
}
