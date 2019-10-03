package api

import (
	"strings"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) GetAllBoards(projectKeyOrID string, word string) ([]jira.Board, error) {
	opts := jira.BoardListOptions{
		BoardType:      "scrum",
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

func (client *Client) GetBoardByName(name string) (*jira.Board, error) {
	boards, err := client.GetAllBoards("", name)
	if err != nil {
		return nil, err
	}

	for _, board := range boards {
		if strings.TrimSpace(board.Name) == strings.TrimSpace(name) {
			return &board, nil
		}
	}

	// No such board found.
	return nil, nil
}

func (client *Client) GetBoardByID(boardID int) (*jira.Board, error) {
	board, _, err := client.jira.Board.GetBoard(boardID)
	if err != nil {
		return nil, err
	}

	return board, nil
}
