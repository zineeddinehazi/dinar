package cmd

import (
	"example/go-dinar/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an item from the list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments for [ remove ] command.")
			return
		} else if len(args) > 1 {
			fmt.Println("Too many arguments for [ remove ] command.")
			return
		}
		if args[0] == " " {
			fmt.Println("Empty values are not allowed.")
			return
		}
		items, err := pkg.DecodeJSON()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "failed to load items: %v\n", err)
			return
		}
		for i, item := range items {
			if item.ID == args[0] {
				items = append(items[:i], items[i+1:]...)
				out := pkg.EncodeJSON(items)
				os.WriteFile("data.json", out, 0644)
				fmt.Printf("Item with title [ %s ] has been deleted.\n", item.Title)
				return
			}
		}
		fmt.Println("No item matches the provided ID.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
