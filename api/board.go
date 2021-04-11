package api

import (
	"context"
	"strings"

	"github.com/andygrunwald/go-jira"
)

func (client *Client) GetAllBoards(ctx context.Context, projectKeyOrID string, word string) ([]jira.Board, error) {
	opts := jira.BoardListOptions{
		BoardType:      "scrum",
		Name:           word,
		ProjectKeyOrID: projectKeyOrID,
	}

	boards := []jira.Board{}

	for {
		list, _, err := client.jira.Board.GetAllBoardsWithContext(ctx, &opts)
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

func (client *Client) GetBoardByName(ctx context.Context, name string) (*jira.Board, error) {
	boards, err := client.GetAllBoards(ctx, "", name)
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

func (client *Client) GetBoardByID(ctx context.Context, boardID int) (*jira.Board, error) {
	board, _, err := client.jira.Board.GetBoardWithContext(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return board, nil
}
