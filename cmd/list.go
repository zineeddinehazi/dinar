package cmd

import (
	"fmt"
	"os"

	"example/go-dinar/pkg"

	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Too many arguments for [ list ] command.")
			return
		}
		items, err := pkg.DecodeJSON()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "failed to load items: %v\n", err)
			return
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "Item", "Cost", "Quantity"})
		var sum float64
		for i := range items {
			t.AppendRow([]interface{}{items[i].ID, items[i].Title, fmt.Sprintf("DA %.1f", items[i].Cost), items[i].Quantity})
			sum += items[i].Cost * float64(items[i].Quantity)
		}
		t.AppendSeparator()
		t.AppendFooter(table.Row{fmt.Sprintf("num of items: %d", len(items)), "Total", fmt.Sprintf("DA %.1f", sum), "-----X-----"})
		t.Render()
		if len(items) == 0 {
			fmt.Println("The list is empty!")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
