package cli

import (
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand/filter"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand/issue"
)

const (
	issueToken    = "issue"
	boardToken    = "board"
	projectToken  = "project"
	myfilterToken = "my-filter"
)

type Parser struct {
	tokenizer *Tokenizer
}

type SubCommandParser interface {
	Parse() subcommand.SubCommand
}

func NewParser() *Parser {
	return &Parser{
		tokenizer: NewTokenizer(),
	}
}

func ParseArgs(args []string) (*subcommand.SubCommand, error) {
	p := NewParser()

	return p.Parse(args)
}

func (p *Parser) Parse(args []string) (*subcommand.SubCommand, error) {
	err := p.tokenizer.Tokenize(args)
	if err != nil {
		return nil, err
	}

	token := p.tokenizer.NextToken()
	cmd := p.createSubCommand(token)

	return &cmd, nil
}

func (p *Parser) createSubCommand(token string) subcommand.SubCommand {
	args := p.tokenizer.RestOfTokens()

	switch token {
	case issueToken:
		return issue.NewCommand(args)

	case boardToken:
		parser := NewBoardCommandParser(p.tokenizer)

		return parser.Parse()

	case projectToken:
		parser := NewProjectCommandParser(p.tokenizer)

		return parser.Parse()

	case myfilterToken:
		return filter.NewMyCommand(args)

	default:
		options := append([]string{token}, args...)

		return subcommand.NewHelpCommand(options)
	}
}
