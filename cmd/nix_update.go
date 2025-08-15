package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nixUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update flakes",
	Long:  `Update flakes with nix flake update.`,
	Run: func(cmd *cobra.Command, args []string) {
		nixCmd := exec.Command("nix", "flake", "update")
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error updating flakes:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(nixUpdateCmd)
}
