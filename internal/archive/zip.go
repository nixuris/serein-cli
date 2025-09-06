package archive

import (
	"github.com/spf13/cobra"
)

func init() {
	ZipCmd.AddCommand(zipPasswordCmd)
}

var ZipCmd = &cobra.Command{
	Use:   "zip [archive-name] [target-to-archive]",
	Short: "Archive files with 7z",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		archiveName := args[0]
		targets := ExpandTargets(args[1:])
		cmdToRun := BuildArchiveCommand(archiveName, targets, "")
		RunWithOutput(cmdToRun, "Error archiving:")
	},
}

var zipPasswordCmd = &cobra.Command{
	Use:   "password [archive-name] [target-to-archive]",
	Short: "Archive files with 7z and a password",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		archiveName := args[0]
		targets := ExpandTargets(args[1:])
		cmdToRun := BuildArchiveCommand(archiveName, targets, "your-password")
		RunWithOutput(cmdToRun, "Error archiving with password:")
	},
}
