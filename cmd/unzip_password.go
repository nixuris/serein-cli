package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var unzipPasswordCmd = &cobra.Command{
	Use:   "password [password] [target-to-unarchive]",
	Short: "Extract archives with a password",
	Long:  `Extract archives with a password using 7z.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		target := args[1]

		sevenZipCmd := exec.Command("7z", "x", "-p"+password, target)

		sevenZipCmd.Stdout = os.Stdout
		sevenZipCmd.Stderr = os.Stderr

		if err := sevenZipCmd.Run(); err != nil {
			fmt.Println("Error extracting with password:", err)
			os.Exit(1)
		}
	},
}

func init() {
	unzipCmd.AddCommand(unzipPasswordCmd)
}
