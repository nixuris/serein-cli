package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nixCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up Nix store",
	Long:  `Clean up Nix store with sudo nix-collect-garbage -d.`,
	Run: func(cmd *cobra.Command, args []string) {
		nixCmd := exec.Command("sudo", "nix-collect-garbage", "-d")
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error cleaning up Nix store:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(nixCleanCmd)
}
