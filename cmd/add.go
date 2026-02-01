package cmd

import (
	"fmt"
	"os"

	"example/dinar/pkg"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new item (purchase) to the list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments for [ add ] command.")
			fmt.Println("command: dinar add <item name> <price> <quantity>")
			return
		} else if len(args) > 3 {
			fmt.Println("Too many arguments for [ add ] command.")
			fmt.Println("command: dinar add <item name> <price> <quantity>")
			return
		}
		items, err := pkg.DecodeJSON()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "failed to load items: %v\n", err)
			return
		}
		newTitle := args[0]
		newCost, _ := strconv.ParseFloat(args[1], 64)
		newQuantity, _ := strconv.Atoi(args[2])
		newItem := pkg.Item{
			ID:       uuid.NewString()[:8],
			Title:    newTitle,
			Cost:     newCost,
			Quantity: newQuantity,
		}
		items = append(items, newItem)
		out := pkg.EncodeJSON(items)
		os.WriteFile("data.json", out, 0644)
		fmt.Println("A new item is added to the list successfully!")
		fmt.Printf("Title: %s\nCost: DA %.1f\nQuantity: %d\n", newItem.Title, newItem.Cost, newItem.Quantity)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
