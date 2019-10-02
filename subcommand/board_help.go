package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// BoardHelpCommand describe a command to show the menus on board subcommands.
type BoardHelpCommand struct {
	BoardName string
	BaseCommand
}

func NewBoardHelpCommand(name string, args []string) BoardHelpCommand {
	return BoardHelpCommand{
		BoardName: name,
		BaseCommand: BaseCommand{
			Args: args,
		},
	}
}

func (cmd BoardHelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "sprint",
			desc: "show sprints in this board",
		},
		{
			name: "backlog",
			desc: "show tickets in backlog of this board",
		},
		{
			name: "issue",
			desc: "show tickets in this board",
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

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No items found.", "")
}
