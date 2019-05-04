package command

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/okeyaki/zipaddr/storage"
)

func newServeCommand() *cobra.Command {
	var config = struct {
		DataDir string `json:"data-dir"`
		Listen  string `json:"listen"`
	}{
		DataDir: "",
		Listen:  "",
	}

	serve := &cobra.Command{
		Use:   "serve",
		Short: "郵便番号検索サーバーを起動します",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := storage.SqliteStorage{
				DataDir: config.DataDir,
			}
			if err := s.Open(); err != nil {
				return err
			}
			defer s.Close()

			e := echo.New()
			e.HideBanner = true

			e.Use(middleware.Logger())

			e.GET("/addrs/:zipcode", func(c echo.Context) error {
				addrs, err := s.FindByZipcode(c.Param("zipcode"))
				if err != nil {
					return err
				}

				return c.JSON(http.StatusOK, addrs)
			})

			e.Start(config.Listen)

			return nil
		},
	}

	serve.Flags().StringVarP(&config.DataDir, "data-dir", "", "", "TBD")
	serve.Flags().StringVarP(&config.Listen, "listen", "", "0.0.0.0:3000", "TBD")

	serve.MarkFlagRequired("data-dir")

	return serve
}
