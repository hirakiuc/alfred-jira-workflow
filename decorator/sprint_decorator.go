package decorator

import (
	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type SprintDecorator struct {
	wf      *aw.Workflow
	baseURL string
	sprint  *jira.Sprint
}

func NewSprintDecorator(wf *aw.Workflow) (*SprintDecorator, error) {
	// TBD
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	return &SprintDecorator{
		wf:      wf,
		sprint:  nil,
		baseURL: client.BaseURL(),
	}, nil
}

func (d *SprintDecorator) SetTarget(sprint *jira.Sprint) {
	d.sprint = sprint
}

func (d *SprintDecorator) Title() string {
	return d.sprint.Name
}

func (d *SprintDecorator) SubTitle() string {
	return "TBD:description"
}

func (d *SprintDecorator) URL() string {
	return "TBD"
}
