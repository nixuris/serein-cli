package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nixHomeBuildCmd = &cobra.Command{
	Use:   "build [path/to/flake]",
	Short: "Build a home-manager configuration",
	Long:  `Build a home-manager configuration with home-manager switch --flake.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flakePath := args[0]

		nixCmd := exec.Command("home-manager", "switch", "--flake", flakePath)
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error building home-manager configuration:", err)
			os.Exit(1)
		}
	},
}

var nixHomeCmd = &cobra.Command{
	Use:   "home",
	Short: "Manage home-manager",
	Long:  `Manage home-manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	nixHomeCmd.AddCommand(nixHomeBuildCmd)
	rootCmd.AddCommand(nixHomeCmd)
}
