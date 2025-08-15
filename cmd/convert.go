package cmd

import (
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Music related conversion utilities",
	Long:  `Music related conversion utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	musicCmd.AddCommand(convertCmd)
}
