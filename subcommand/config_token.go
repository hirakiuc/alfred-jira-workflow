package subcommand

import (
	"context"
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-jira-workflow/secret"
)

// ConfigTokenCommand describe a subcommand to configure the api token.
type ConfigTokenCommand struct {
	Token string

	BaseCommand
}

// NewConfigTokenCommand return an instance of ConfigTokenCommand
func NewConfigTokenCommand(token string) ConfigTokenCommand {
	return ConfigTokenCommand{
		Token: token,

		BaseCommand: BaseCommand{
			Args: []string{},
		},
	}
}

// Run start this subcommand
func (cmd ConfigTokenCommand) Run(ctx context.Context, wf *aw.Workflow) {
	store := secret.NewStore(wf)

	if len(cmd.Token) == 0 {
		token, err := store.GetAPIToken()
		if err != nil {
			wf.FatalError(err)
			return
		}

		if len(token) == 0 {
			wf.WarnEmpty("No token found.", "")
			return
		}

		wf.NewItem(fmt.Sprintf("Found token:%s", token))

		return
	}

	err := store.Store(secret.KeyJIRAToken, cmd.Token)
	if err != nil {
		wf.FatalError(err)
		return
	}

	wf.NewItem("Success.")
}
