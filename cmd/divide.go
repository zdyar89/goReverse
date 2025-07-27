package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var roundUpDiv bool
var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"div"},
	Short:   "Divide 2 numbers",
	Long:    "Divide 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Division of %s by %s = %s.\n\n", args[0], args[1], Divide(args[0], args[1], roundUpDiv))
	},
}

func init() {
	divideCmd.Flags().BoolVarP(&roundUpDiv, "round", "r", false, "Round up to 2 decimal places")
	rootCmd.AddCommand(divideCmd)
}
