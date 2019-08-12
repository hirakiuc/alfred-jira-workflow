package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// HelpCommand describe a subcommand to show the subcommands.
type HelpCommand struct {
	BaseCommand
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand() HelpCommand {
	return HelpCommand{
		BaseCommand{
			Args: []string{},
		},
	}
}

func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "{query}",
			desc: "search issue by jql",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete(cmd.name).
			Valid(false)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}
}
