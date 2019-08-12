package subcommand

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
)

type SubCommand interface {
	// Run the subcommand
	Run(ctx context.Context, wf *aw.Workflow)

	// Get the arguments
	Arguments() []string
}

type BaseCommand struct {
	Args []string
}

func (cmd BaseCommand) Arguments() []string {
	return cmd.Args
}

func (cmd BaseCommand) HasQuery() bool {
	return len(cmd.Args) > 0
}

func (cmd BaseCommand) Query() string {
	return strings.Join(cmd.Args, " ")
}
