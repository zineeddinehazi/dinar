package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dinar",
	Short: "Simple expense tracker.",
	Long: `
Command line tool to manage expenses through minimal ascii list.
It allows to add/remove/update items.
It displays the total cost by default.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
