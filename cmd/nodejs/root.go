package nodejs

import (
	"github.com/spf13/cobra"
)

func init() {
	RootNodeCmd.AddCommand(VersionsCmd)
	RootNodeCmd.AddCommand(InstallCmd)
	RootNodeCmd.AddCommand(GlobalCmd)
}

var RootNodeCmd = &cobra.Command{
	Use:     "nodejs",
	Aliases: []string{"node"},
	Short:   "Management command of node environment.",
}
