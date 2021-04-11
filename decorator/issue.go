package decorator

import (
	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type IssueDecorator struct {
	wf      *aw.Workflow
	baseURL string
	issue   *jira.Issue
}

func NewIssueDecorator(wf *aw.Workflow) (*IssueDecorator, error) {
	// TBD
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	return &IssueDecorator{
		wf:      wf,
		issue:   nil,
		baseURL: client.BaseURL(),
	}, nil
}

func (d *IssueDecorator) SetTarget(issue *jira.Issue) {
	d.issue = issue
}

func (d *IssueDecorator) Title() string {
	issue := d.issue
	if issue == nil || issue.Fields == nil {
		return ""
	}

	return issue.Fields.Summary
}

func (d *IssueDecorator) URL() string {
	issue := d.issue
	if issue == nil {
		return ""
	}

	// TBD
	return d.baseURL + "browse/" + issue.Key
}
