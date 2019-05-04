package command

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/okeyaki/zipaddr/stream/write/flow"
	"github.com/okeyaki/zipaddr/stream/write/sink"
)

func newNormalizeCommand() *cobra.Command {
	var config = struct {
		File string `json:"file"`
	}{
		File: "",
	}

	normalize := &cobra.Command{
		Use:   "normalize",
		Short: "郵便番号データを正規化します",
		RunE: func(cmd *cobra.Command, args []string) error {
			var src io.Reader
			if config.File == "" {
				src = os.Stdin
			} else {
				file, err := os.Open(config.File)
				if err != nil {
					return err
				}
				defer file.Close()

				src = file
			}

			sink := sink.StreamSink{
				Stream: os.Stdout,
			}

			if err := sink.Write(flow.Normalize(src)); err != nil {
				return err
			}

			return nil
		},
	}

	normalize.Flags().StringVarP(&config.File, "file", "f", "", "TBD")

	return normalize
}
