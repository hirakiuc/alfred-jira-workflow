package cli

import (
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

const (
	cmdTypeBoardSprint  = "sprint"
	cmdTypeBoardBacklog = "backlog"
	cmdTypeBoardIssue   = "issue"
)

// BoardCommandParser describe a parser for board subcommand.
type BoardCommandParser struct {
	tokenizer *Tokenizer
}

func NewBoardCommandParser(tokenizer *Tokenizer) *BoardCommandParser {
	return &BoardCommandParser{
		tokenizer: tokenizer,
	}
}

func (p *BoardCommandParser) Parse() subcommand.SubCommand {
	// cmd: board
	// => show  board list
	args := p.tokenizer.RestOfTokens()
	if len(args) == 0 {
		return subcommand.NewBoardCommand(args)
	}

	// cmd: board {boardID} ...
	// => show board help
	// cmd: board {boardID} {token} ...
	// => show subcommand
	boardID := p.tokenizer.NextToken()

	return p.parseBoardCommands(boardID)
}

func (p *BoardCommandParser) parseBoardCommands(boardID string) subcommand.SubCommand {
	// cmd: board {token} [options...]
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeBoardBacklog:
		return subcommand.NewBoardBacklogCommand(boardID, opts)

	case cmdTypeBoardIssue:
		return subcommand.NewBoardIssueCommand(boardID, opts)

	case cmdTypeBoardSprint:
		return subcommand.NewBoardSprintCommand(boardID, opts)

	default:
		options := append([]string{boardID, token}, opts...)

		return subcommand.NewBoardHelpCommand(options)
	}
}
