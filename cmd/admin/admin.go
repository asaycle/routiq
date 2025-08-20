package admin

import (
	"github.com/asaycle/routiq/cmd/admin/location"
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Routiq admin CLI",
}

func init() {
	adminCmd.AddCommand(location.Command())
}

func Command() *cobra.Command {
	return adminCmd
}
