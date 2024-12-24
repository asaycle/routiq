package server

import (
	"fmt"

	"github.com/asaycle/motify.git/pkg/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("server start")
		srv := server.NewServer()
		return srv.Start()
	},
}
