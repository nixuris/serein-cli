package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nixSysBuildCmd = &cobra.Command{
	Use:   "build [path/to/flake]",
	Short: "Build a NixOS system",
	Long:  `Build a NixOS system with sudo nixos-rebuild switch --impure --flake.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flakePath := args[0]

		nixCmd := exec.Command("sudo", "nixos-rebuild", "switch", "--impure", "--flake", flakePath)
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error building system:", err)
			os.Exit(1)
		}
	},
}

var nixSysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Manage NixOS system",
	Long:  `Manage NixOS system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	nixSysCmd.AddCommand(nixSysBuildCmd)
	rootCmd.AddCommand(nixSysCmd)
}
