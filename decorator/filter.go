package decorator

import (
	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type FilterDecorator struct {
	wf      *aw.Workflow
	baseURL string
	filter  *jira.Filter
}

func NewFilterDecorator(wf *aw.Workflow) (*FilterDecorator, error) {
	// TBD
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	return &FilterDecorator{
		wf:      wf,
		baseURL: client.BaseURL(),
		filter:  nil,
	}, nil
}

func (d *FilterDecorator) SetTarget(filter *jira.Filter) {
	d.filter = filter
}

func (d *FilterDecorator) Title() string {
	f := d.filter
	if f == nil {
		return ""
	}

	return f.Name
}

func (d *FilterDecorator) Subtitle() string {
	f := d.filter
	if f == nil {
		return ""
	}

	return f.Description
}
