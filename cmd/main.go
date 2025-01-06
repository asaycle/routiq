package main

import (
	"fmt"

	"github.com/asaycle/routiq.git/cmd/server"
	"github.com/spf13/cobra"
)

var routiqCmd = &cobra.Command{
	Use:   "routiq",
	Short: "A server CLI for managing gRPC and gRPC-Gateway",
	Long:  "This CLI tool allows you to start gRPC and gRPC-Gateway servers.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand (e.g., grpc, gateway)")
	},
}

func init() {
	routiqCmd.AddCommand(server.Command())
}

func main() {
	if err := routiqCmd.Execute(); err != nil {
		panic(err)
	}
}
