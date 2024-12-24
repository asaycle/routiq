package main

import (
	"fmt"

	"github.com/asaycle/motify.git/cmd/server"
	"github.com/spf13/cobra"
)

var motifyCmd = &cobra.Command{
	Use:   "motify",
	Short: "A server CLI for managing gRPC and gRPC-Gateway",
	Long:  "This CLI tool allows you to start gRPC and gRPC-Gateway servers.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand (e.g., grpc, gateway)")
	},
}

func init() {
	motifyCmd.AddCommand(server.Command())
}

func main() {
	if err := motifyCmd.Execute(); err != nil {
		panic(err)
	}
}
