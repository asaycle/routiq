package server

import (
	"fmt"

	"github.com/asaycle/routiq.git/pkg/lib/config"
	"github.com/asaycle/routiq.git/pkg/server"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return xerrors.Errorf("Error loading config: %w", err)
		}
		fmt.Println("server start", cfg)

		srv := server.NewServer()
		return srv.Start()
	},
}
