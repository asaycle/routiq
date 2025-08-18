package main

import (
	"context"
	"fmt"

	"github.com/asaycle/routiq/cmd/db"
	"github.com/asaycle/routiq/cmd/server"
	"github.com/asaycle/routiq/pkg/lib/config"
	"github.com/asaycle/routiq/pkg/sentry"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var routiqCmd = &cobra.Command{
	Use:   "routiq",
	Short: "A server CLI for managing gRPC and gRPC-Gateway",
	Long:  "This CLI tool allows you to start gRPC and gRPC-Gateway servers.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		sentry.InitializeSentry()
		cfg, err := config.Load()
		if err != nil {
			return xerrors.Errorf("failed to load config: %w", err)
		}
		ctx := context.WithValue(cmd.Context(), config.ConfigKey, cfg)
		cmd.SetContext(ctx)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand (e.g., grpc, gateway)")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		sentry.Flush()
	},
}

func init() {
	routiqCmd.AddCommand(server.Command())
	routiqCmd.AddCommand(db.Command())
}

func main() {
	if err := routiqCmd.Execute(); err != nil {
		panic(err)
	}
}
