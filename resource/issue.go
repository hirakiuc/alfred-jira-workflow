package resource

import (
	"context"
	"fmt"
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

func (r *IssueResource) SearchByJQL(ctx context.Context, jql string) ([]jira.Issue, error) {
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

	issues, err = client.SearchIssues(ctx, jql)
	if err != nil {
		return []jira.Issue{}, err
	}

	if len(issues) == 0 {
		return []jira.Issue{}, nil
	}

	return store.StoreByJQL(jql, issues)
}

func (r *IssueResource) SearchIssues(ctx context.Context, args []string) ([]jira.Issue, error) {
	jql := args2jql(args)

	return r.SearchByJQL(ctx, jql)
}

func (r *IssueResource) ListByProjectID(ctx context.Context, prjID string) ([]jira.Issue, error) {
	jql := fmt.Sprintf("project = %s ORDER BY created desc, priority desc", prjID)

	return r.SearchByJQL(ctx, jql)
}
