package cli

import "github.com/hirakiuc/alfred-jira-workflow/subcommand"

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

	opts := parser.tokenizer.RestOfTokens()

	return subcommand.NewIssueCommand(opts)
}

/*
func (parser *Parser) createSubCommandParser(token string) SubCommandParser {
	args := parser.tokenizer.RestOfTokens()

	switch token {
	default:
	}
}
*/
