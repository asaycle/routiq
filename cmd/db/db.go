package db

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	return dbCmd
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "db command",
}

func init() {
	dbCmd.AddCommand(migrateCmd)
}
