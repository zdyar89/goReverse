package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "substract",
	Aliases: []string{"sub"},
	Short:   "Subtract a number from a number",
	Long:    "Subtract a number from a number",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
