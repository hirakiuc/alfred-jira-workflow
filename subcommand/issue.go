package subcommand

import (
	"context"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type IssueCommand struct {
	BaseCommand
}

func NewIssueCommand(args []string) IssueCommand {
	return IssueCommand{
		BaseCommand{
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

func (cmd IssueCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient()
	if err != nil {
		wf.FatalError(err)
		return
	}

	issues, err := client.SearchIssues(cmd.Arguments())
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, issue := range issues {
		wf.NewItem(IssueTitle(issue)).
			Arg(IssueURL(client.BaseURL(), issue)).
			Valid(true)
	}

	/*
		if cmd.HasQuery() {
			wf.Filter(cmd.Query())
		}
	*/

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
