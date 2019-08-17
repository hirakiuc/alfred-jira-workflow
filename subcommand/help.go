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
func NewHelpCommand(args []string) HelpCommand {
	return HelpCommand{
		BaseCommand{
			Args: args,
		},
	}
}

func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "issue",
			desc: "search issue by jql",
		},
		{
			name: "board",
			desc: "",
		},
		{
			name: "project",
			desc: "",
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

	// Show a warning in Alfred if there are no repos
	wf.WarnEmpty("No commands found.", "")
}
