package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"langenv/cmd/nodejs"
)

var rootCmd = &cobra.Command{
	Use:   "langenv",
	Short: "langenv is a tool to manage environment of languages available in your system and switch easily between them.",

}

func init() {
	rootCmd.AddCommand(nodejs.RootNodeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
