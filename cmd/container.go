package cmd

import (
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage containers with podman aliases",
	Long:  `Manage containers with podman aliases.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Offer help in case of no sub-command
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)
}
