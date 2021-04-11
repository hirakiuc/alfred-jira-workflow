package decorator

import (
	"fmt"

	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/model"
)

type ProjectDecorator struct {
	wf  *aw.Workflow
	prj *model.ProjectData
}

func NewProjectDecorator(wf *aw.Workflow) (*ProjectDecorator, error) {
	return &ProjectDecorator{
		wf: wf,
	}, nil
}

func (d *ProjectDecorator) SetTarget(prj *model.ProjectData) {
	d.prj = prj
}

func (d *ProjectDecorator) GetTitle() string {
	return fmt.Sprintf("%s - %s", d.prj.Key, d.prj.Name)
}

func (d *ProjectDecorator) GetSubtitle() string {
	return fmt.Sprintf("%s - %s", d.prj.ID, d.prj.Self)
}
