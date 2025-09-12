package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var rootCmd = &cobra.Command{
	Use:   "serein",
	Short: "Serein is an opinionated CLI tool.",
	Long:  `Serein is an opinionated CLI tool that provides aliases for multiple utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Root command does nothing by itself
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&shared.DryRun, "dry-run", false, "print the command that would be executed instead of executing it")
}
