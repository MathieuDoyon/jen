package list

import (
	"strings"

	"github.com/Samasource/jen/src/cmd/internal"
	"github.com/Samasource/jen/src/cmd/list/actions"
	"github.com/Samasource/jen/src/internal/shell"
	"github.com/spf13/cobra"
)

// New creates a cobra command
func New(options *internal.Options) *cobra.Command {
	c := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists available templates, actions, variables or scripts",
		RunE: func(_ *cobra.Command, args []string) error {
			return run(options, args)
		},
	}
	c.AddCommand(actions.New(options))
	return c
}

func run(options *internal.Options, args []string) error {
	execContext, err := options.NewContext()
	if err != nil {
		return err
	}

	return shell.Execute(execContext.GetShellVars(), "", strings.Join(args, " "))
}
