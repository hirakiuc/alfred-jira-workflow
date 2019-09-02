package filter

import (
	"context"
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
	"github.com/hirakiuc/alfred-jira-workflow/cache"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

type IssueCommand struct {
	FilterID string
	subcommand.BaseCommand
}

func NewIssueCommand(filterID string, args []string) IssueCommand {
	return IssueCommand{
		FilterID: filterID,
		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func IssueURL(baseURL string, issue jira.Issue) string {
	// TBD
	return baseURL + "browse/" + issue.Key
}

func IssueTitle(issue jira.Issue) string {
	if issue.Fields == nil {
		return ""
	}

	return issue.Fields.Summary
}

func (cmd IssueCommand) fetchIssues(_ context.Context, wf *aw.Workflow, client *api.Client) ([]jira.Issue, error) {
	store := cache.NewFilterIssuesCache(wf)

	issues, err := store.GetCache(cmd.FilterID)
	if err != nil {
		return []jira.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.SearchIssuesByFilter(cmd.FilterID, "updated DESC")
	if err != nil {
		return []jira.Issue{}, err
	}
	if len(issues) == 0 {
		return issues, nil
	}

	return store.Store(issues, cmd.FilterID)
}

func (cmd IssueCommand) Run(ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient()
	if err != nil {
		wf.FatalError(err)
		return
	}

	issues, err := cmd.fetchIssues(ctx, wf, client)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, issue := range issues {
		wf.NewItem(IssueTitle(issue)).
			Arg(IssueURL(client.BaseURL(), issue)).
			Valid(true)
	}

	if cmd.HasQuery() {
		fmt.Fprintf(os.Stderr, "query:%s\n", cmd.Query())
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
