package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for math ops",
	Long:  "zero is a cli tool for math ops",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while func '%s'\n", err)
		os.Exit(1)
	}
}
