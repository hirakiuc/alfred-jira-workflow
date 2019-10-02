package cli

import "github.com/hirakiuc/alfred-jira-workflow/subcommand"

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
		return subcommand.NewIssueCommand(args)
	case boardToken:
		return p.parseBoardSubCommannd()
	case projectToken:
		return subcommand.NewProjectCommand(args)
	case myfilterToken:
		return subcommand.NewMyFilterCommand(args)
	default:
		options := append([]string{token}, args...)
		return subcommand.NewHelpCommand(options)
	}
}

func (p *Parser) parseBoardSubCommannd() subcommand.SubCommand {
	args := p.tokenizer.RestOfTokens()
	if len(args) == 0 {
		// 'board', ... => show board list
		return subcommand.NewBoardCommand(args)
	}

	// 'board', token, ...
	token := p.tokenizer.NextToken()
	parser := NewBoardCommandParser(p.tokenizer, token)
	return parser.Parse()
}
