package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var roundUpMul bool
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mul", "product"},
	Short:   "Multiply 2 numbers",
	Long:    "Multiply 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplication of %s and %s = %s.\n\n", args[0], args[1], Multiply(args[0], args[1], roundUpMul))
	},
}

func init() {
	multiplyCmd.Flags().BoolVarP(&roundUpMul, "round", "r", false, "Round up to 2 decimal places")
	rootCmd.AddCommand(multiplyCmd)
}
