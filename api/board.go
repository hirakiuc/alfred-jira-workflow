package api

import (
	"github.com/andygrunwald/go-jira"
)

func (client *Client) GetAllBoards(projectKeyOrID string, word string) ([]jira.Board, error) {
	opts := jira.BoardListOptions{
		BoardType:      "kanban",
		Name:           word,
		ProjectKeyOrID: projectKeyOrID,
	}

	boards := []jira.Board{}

	for {
		list, _, err := client.jira.Board.GetAllBoards(&opts)
		if err != nil {
			return boards, err
		}

		boards = append(boards, list.Values...)

		if list.IsLast {
			return boards, nil
		}

		// next page
		opts.StartAt += len(list.Values)
	}
}

func (client *Client) GetAllSprints(boardID string) ([]jira.Sprint, error) {
	sprints, _, err := client.jira.Board.GetAllSprints(boardID)
	if err != nil {
		return []jira.Sprint{}, err
	}

	return sprints, nil
}
