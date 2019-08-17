package cli

import "github.com/hirakiuc/alfred-jira-workflow/subcommand"

const (
	issueToken   = "issue"
	boardToken   = "board"
	projectToken = "project"
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

func ParseArgs(args []string) subcommand.SubCommand {
	parser := NewParser()
	return parser.Parse(args)
}

func (parser *Parser) Parse(args []string) subcommand.SubCommand {
	parser.tokenizer.Tokenize(args)

	token := parser.tokenizer.NextToken()
	return parser.createSubCommand(token)
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
