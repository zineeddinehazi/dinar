package cmd

import (
	"example/go-dinar/pkg"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an item in the list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			fmt.Println("Not enough arguments for [ update ] command.")
			return
		} else if len(args) > 4 {
			fmt.Println("Too many arguments for [ update ] command.")
			return
		}
		if args[0] == " " || args[1] == " " || args[2] == " " || args[3] == " " {
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
				newTitle := args[1]
				newCost, _ := strconv.ParseFloat(args[2], 64)
				newQuantity, _ := strconv.Atoi(args[3])
				items[i] = pkg.Item{
					ID:       uuid.NewString()[:8],
					Title:    newTitle,
					Cost:     newCost,
					Quantity: newQuantity,
				}
				out := pkg.EncodeJSON(items)
				os.WriteFile("data.json", out, 0644)
				fmt.Printf("Item with ID [ %s ] has been updated.\n", item.ID)
				fmt.Printf("Title: %s\nCost: DA %.1f\nQuantity: %d\n", newTitle, newCost, newQuantity)
				return
			}
		}
		fmt.Println("No item matches the provided ID.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
