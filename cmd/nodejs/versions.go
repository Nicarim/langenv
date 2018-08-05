package nodejs

import (
	"github.com/spf13/cobra"
	"langenv/providers"
	"fmt"
)

func init() {
	VersionsCmd.Flags().BoolVarP(&versionsAllPatches, "show-patches", "p", false,
		"Show all patch versions available.")
}

var versionsAllPatches bool

var VersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Lists available versions for given environment",
	Run: func(cmd *cobra.Command, args []string) {
		versions := providers.GetVersionsAvailable(versionsAllPatches)
		for _, element := range versions {
			fmt.Println(element.Version)
		}
	},
}
