package command

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "zipaddr",
		Short: "TBD",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	root.AddCommand(newNormalizeCommand())
	root.AddCommand(newServeCommand())
	root.AddCommand(newUpdateCommand())

	return root
}
