package cmd

import (
	"example/go-dinar/pkg"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the entire list",
	Long:  `Clear evry item on the expense list.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Too many arguments for [ clear ] command.")
			return
		}
		var answer string
		fmt.Print("Are you sure ? (Y) Yes / (N) No : ")
		fmt.Scanf("%s", &answer)
		if strings.ToLower(answer) == "n" {
			fmt.Println("Operation cancelled by the user.")
		} else if strings.ToLower(answer) == "y" {
			items := make([]pkg.Item, 0)
			out := pkg.EncodeJSON(items)
			os.WriteFile("data.json", out, 0644)
			fmt.Println("List cleared!")
		} else {
			fmt.Println("Answer not expected, operation cancelled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
