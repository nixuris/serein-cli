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
	Use:   "delete [number]",
	Short: "Delete home-manager generations",
	Long:  `Delete home-manager generations with home-manager remove-generations.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generation := args[0]
		runNixCommand("home-manager", "remove-generations", generation)
	},
}
