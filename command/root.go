package command

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "zipaddr",
		Short: "TBD",
	}

	root.AddCommand(newNormalizeCommand())
	root.AddCommand(newServeCommand())
	root.AddCommand(newUpdateCommand())

	return root
}
