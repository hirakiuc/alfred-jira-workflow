package resource

import (
	"strings"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
)

type IssueResource struct {
	wf *aw.Workflow
}

func NewIssueResource(wf *aw.Workflow) *IssueResource {
	return &IssueResource{
		wf: wf,
	}
}

func args2jql(args []string) string {
	return strings.Join(args, " AND ")
}

func (r *IssueResource) SearchIssues(args []string) ([]jira.Issue, error) {
	jql := args2jql(args)
	store := cache.NewIssuesCache(r.wf)

	issues, err := store.GetCacheByJQL(jql)
	if err != nil {
		return []jira.Issue{}, err
	}

	if len(issues) > 0 {
		return issues, nil
	}

	client, err := api.NewClient()
	if err != nil {
		return []jira.Issue{}, err
	}

	issues, err = client.SearchIssues(jql)
	if err != nil {
		return []jira.Issue{}, err
	}

	if len(issues) == 0 {
		return []jira.Issue{}, nil
	}

	return store.StoreByJQL(jql, issues)
}
