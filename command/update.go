package command

import (
	"github.com/okeyaki/zipaddr/stream/write/flow"
	"github.com/okeyaki/zipaddr/stream/write/sink"
	"github.com/okeyaki/zipaddr/stream/write/source"
	"github.com/spf13/cobra"
)

func newUpdateCommand() *cobra.Command {
	var config = struct {
		File    string `json:"file"`
		DataDir string `json:"data-dir"`
	}{
		DataDir: "",
	}

	update := &cobra.Command{
		Use:   "update",
		Short: "TBD",
		RunE: func(cmd *cobra.Command, args []string) error {
			src := source.URLSource{
				URL: "https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/ken_all.zip",
			}
			if err := src.Open(); err != nil {
				return err
			}
			defer src.Close()

			sink := sink.SqliteSink{
				DataDir: config.DataDir,
			}
			if err := sink.Open(); err != nil {
				return err
			}
			defer sink.Close()

			if err := sink.Write(flow.Normalize(src)); err != nil {
				return err
			}

			return nil

		},
	}

	update.Flags().StringVarP(&config.DataDir, "data-dir", "", "", "TBD")

	update.MarkFlagRequired("data-dir")

	return update
}
