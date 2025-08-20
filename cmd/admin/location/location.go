package location

import "github.com/spf13/cobra"

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "location cli",
}

func init() {
	locationCmd.AddCommand(importCmd)
}

func Command() *cobra.Command {
	return locationCmd
}
