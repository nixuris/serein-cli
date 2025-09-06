package archive

import (
	"github.com/spf13/cobra"
)

func init() {
	UnzipCmd.AddCommand(unzipPasswordCmd)
}

var UnzipCmd = &cobra.Command{
	Use:   "unzip [target-to-unarchive]",
	Short: "Extract archives with 7z",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdToRun := BuildExtractCommand(args[0], "")
		RunWithOutput(cmdToRun, "Error extracting:")
	},
}

var unzipPasswordCmd = &cobra.Command{
	Use:   "password [password] [target-to-unarchive]",
	Short: "Extract archives with a password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmdToRun := BuildExtractCommand(args[1], args[0])
		RunWithOutput(cmdToRun, "Error extracting with password:")
	},
}
