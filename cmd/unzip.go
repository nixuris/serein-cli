package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var unzipCmd = &cobra.Command{
	Use:   "unzip [target-to-unarchive]",
	Short: "Extract archives with 7z",
	Long:  `Extract archives with 7z.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		sevenZipCmd := exec.Command("7z", "x", target)

		sevenZipCmd.Stdout = os.Stdout
		sevenZipCmd.Stderr = os.Stderr

		if err := sevenZipCmd.Run(); err != nil {
			fmt.Println("Error extracting:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)
}
