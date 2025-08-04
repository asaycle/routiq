package server

import (
	"fmt"

	"github.com/asaycle/routiq.git/pkg/lib/config"
	"github.com/asaycle/routiq.git/pkg/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.FromContext(cmd.Context())
		fmt.Println("server start", cfg)
		srv := server.NewServer(cfg)
		return srv.Start()
	},
}
