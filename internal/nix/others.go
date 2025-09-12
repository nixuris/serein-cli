package nix

import (
	"github.com/spf13/cobra"
)

func BasicNixCmds(parent *cobra.Command) {
	parent.AddCommand(nixUpdateCmd)
	parent.AddCommand(nixSearchCmd)
	parent.AddCommand(nixCleanCmd)
	parent.AddCommand(nixLintCmd)
}

var nixSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for nix packages",
	Long:  `Search for nix packages with nix search.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nixpkgsName := args[0]
		runNixCommand("nix", "search", "nixpkgs", nixpkgsName)
	},
}

var nixUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update flakes",
	Long:  `Update flakes with nix flake update.`,
	Run: func(cmd *cobra.Command, args []string) {
		runNixCommand("nix", "flake", "update")
	},
}

var nixCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up Nix store",
	Long:  `Clean up Nix store with sudo nix-collect-garbage -d.`,
	Run: func(cmd *cobra.Command, args []string) {
		runNixCommand("sudo", "nix-collect-garbage", "-d")
	},
}

var nixLintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Linter For Nix",
	Long:  `Format Nix Files With Alejandra.`,
	Run: func(cmd *cobra.Command, args []string) {
		runNixCommand("nix-shell", "-p", "alejandra", "--run", "alejandra .")
	},
}
