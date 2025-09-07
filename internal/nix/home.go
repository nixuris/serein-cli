package nix

import (
	"github.com/spf13/cobra"
)

func init() {
	HomeCmd.AddCommand(HomeBuildCmd)
	HomeCmd.AddCommand(HomeGenCmd)
	HomeCmd.AddCommand(HomeGenDeleteCmd)
}

var HomeCmd = &cobra.Command{
	Use:   "home",
	Short: "Manage home-manager",
	Long:  `Manage home-manager.`,
}

var HomeBuildCmd = &cobra.Command{
	Use:   "build [path/to/flake]",
	Short: "Build a home-manager configuration",
	Long:  `Build a home-manager configuration with home-manager switch --flake.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flakePath := args[0]
		runNixCommand("home-manager", "switch", "--flake", flakePath)
	},
}

var HomeGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "List home-manager generations",
	Long:  `List home-manager generations with home-manager generations.`,
	Run: func(cmd *cobra.Command, args []string) {
		runNixCommand("home-manager", "generations")
	},
}

var HomeGenDeleteCmd = &cobra.Command{
	Use:   "delete [numbers...]",
	Short: "Delete home-manager generations",
	Long:  `Delete home-manager generations with home-manager remove-generations. Can accept multiple numbers and ranges (e.g., 1-10).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generations := parseGenerations(args)
		cmdArgs := append([]string{"remove-generations"}, generations...)
		runNixCommand("home-manager", cmdArgs...)
	},
}
