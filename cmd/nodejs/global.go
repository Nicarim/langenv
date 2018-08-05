package nodejs

import "github.com/spf13/cobra"

var GlobalCmd = &cobra.Command{
	Use: "global",
	Short: "Sets given version as globally used version for currently logged in user",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
