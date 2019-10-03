package api

import "github.com/andygrunwald/go-jira"

func (client *Client) GetSprintsInBoard(boardID int) ([]jira.Sprint, error) {
	opts := jira.GetAllSprintsOptions{
		// https://developer.atlassian.com/cloud/jira/software/rest/#api-rest-agile-1-0-board-boardId-sprint-get
		State: "active,future",
	}

	sprints := []jira.Sprint{}

	for {
		list, _, err := client.jira.Board.GetAllSprintsWithOptions(boardID, &opts)
		if err != nil {
			return []jira.Sprint{}, err
		}

		sprints = append(sprints, list.Values...)

		if list.IsLast {
			return sprints, nil
		}

		// next page
		opts.StartAt += len(list.Values)
	}
}
