package server

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	return serverCmd
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server command",
}

func init() {
	serverCmd.AddCommand(startCmd)
}
