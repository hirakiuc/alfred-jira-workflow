package decorator

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	aw "github.com/deanishe/awgo"

	"github.com/hirakiuc/alfred-jira-workflow/api"
)

type BoardDecorator struct {
	wf      *aw.Workflow
	baseURL string
	board   *jira.Board
}

func NewBoardDecorator(wf *aw.Workflow) (*BoardDecorator, error) {
	// TBD
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	return &BoardDecorator{
		wf:      wf,
		board:   nil,
		baseURL: client.BaseURL(),
	}, nil
}

func (d *BoardDecorator) SetTarget(board *jira.Board) {
	d.board = board
}

func (d *BoardDecorator) Title() string {
	b := d.board
	if b == nil {
		return ""
	}

	return b.Name
}

func (d *BoardDecorator) Subtitle() string {
	b := d.board
	if b == nil {
		return ""
	}

	return fmt.Sprintf("%d - %s", b.ID, b.Self)
}
