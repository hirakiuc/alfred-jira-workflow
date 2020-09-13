package cli

import (
	"github.com/hirakiuc/alfred-jira-workflow/subcommand"
	"github.com/hirakiuc/alfred-jira-workflow/subcommand/project"
)

const (
	cmdTypeProjectIssue = "issue"
)

type ProjectCommandParser struct {
	tokenizer *Tokenizer
}

func NewProjectCommandParser(tokenizer *Tokenizer) *ProjectCommandParser {
	return &ProjectCommandParser{
		tokenizer: tokenizer,
	}
}

func (p *ProjectCommandParser) Parse() subcommand.SubCommand {
	// cmd: project
	args := p.tokenizer.RestOfTokens()
	if len(args) == 0 {
		return project.NewCommand(args)
	}

	// cmd: project {projectID} ...
	prjID := p.tokenizer.NextToken()

	return p.parseProjectCommand(prjID)
}

func (p *ProjectCommandParser) parseProjectCommand(projectID string) subcommand.SubCommand {
	// cmd: project {projectID} ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeProjectIssue:
		return project.NewIssueCommand(projectID, opts)

	default:
		options := append([]string{projectID, token}, opts...)

		return project.NewHelpCommand(options)
	}
}
