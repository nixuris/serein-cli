package cmd

import (
	"github.com/spf13/cobra"
)

var musicCmd = &cobra.Command{
	Use:   "music",
	Short: "Music related utilities",
	Long:  `Music related utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(musicCmd)
}
