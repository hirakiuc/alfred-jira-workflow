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
	parser := NewParser()
	return parser.Parse(args)
}

func (parser *Parser) Parse(args []string) (*subcommand.SubCommand, error) {
	err := parser.tokenizer.Tokenize(args)
	if err != nil {
		return nil, err
	}

	token := parser.tokenizer.NextToken()
	cmd := parser.createSubCommand(token)
	return &cmd, nil
}

func (parser *Parser) createSubCommand(token string) subcommand.SubCommand {
	args := parser.tokenizer.RestOfTokens()

	switch token {
	case issueToken:
		return subcommand.NewIssueCommand(args)
	case boardToken:
		return subcommand.NewBoardCommand(args)
	case projectToken:
		return subcommand.NewProjectCommand(args)
	case myfilterToken:
		return subcommand.NewMyFilterCommand(args)
	default:
		options := append([]string{token}, args...)
		return subcommand.NewHelpCommand(options)
	}
}

/*
func (parser *Parser) createSubCommandParser(token string) SubCommandParser {
	args := parser.tokenizer.RestOfTokens()

	switch token {
	case issueToken:
		return subcommand.NewIssueCommand(args)
	case boardToken:
		return subcommand.NewBoardCommand(args)
	case projectToken:
		return subcommand.NewProjectCommand(args)
	default:
		return subcommand.HelpCommand(args)
	}
}
*/
