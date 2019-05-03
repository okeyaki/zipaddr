package command

import (
	"os"
)

func Execute() {
	cmd := newRootCommand()
	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)

		os.Exit(1)
	}
}
