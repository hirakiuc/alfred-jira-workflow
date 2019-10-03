package cli

import (
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
)

const (
	cmdTypeBoardSprint  = "sprint"
	cmdTypeBoardBacklog = "backlog"
	cmdTypeBoardIssue   = "issue"
)

// BoardCommandParser describe a parser for board subcommand
type BoardCommandParser struct {
	BoardName string
	tokenizer *Tokenizer
}

func NewBoardCommandParser(tokenizer *Tokenizer, name string) *BoardCommandParser {
	return &BoardCommandParser{
		BoardName: name,
		tokenizer: tokenizer,
	}
}

func (p *BoardCommandParser) Parse() subcommand.SubCommand {
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeBoardBacklog:
		return subcommand.NewBoardBacklogCommand(p.BoardName, opts)
	case cmdTypeBoardIssue:
		return subcommand.NewBoardIssueCommand(p.BoardName, opts)
	case cmdTypeBoardSprint:
		return subcommand.NewBoardSprintCommand(p.BoardName, opts)
	default:
		return subcommand.NewBoardHelpCommand(p.BoardName, opts)
	}
}
