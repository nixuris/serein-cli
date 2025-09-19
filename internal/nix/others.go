package nix

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

func BasicNixCmds(parent *cobra.Command) {
	parent.AddCommand(nixUpdateCmd)
	parent.AddCommand(nixSearchCmd)
	parent.AddCommand(nixHashCmd)
	parent.AddCommand(nixCleanCmd)
	parent.AddCommand(nixLintCmd)
}

var nixHashCmd = shared.NewCommand(
	"hash",
	"Prefetch a Nix package URL and convert its hash to base64",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runNixCommand("sh", "-c",
			"nix-prefetch-url --type sha256 \"$1\" | xargs nix hash convert --hash-algo sha256 --to base64",
			"_", args[0])
	},
)

var nixSearchCmd = shared.NewCommand(
	"search",
	"Search for nix packages",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		nixpkgsName := args[0]
		runNixCommand("nix", "search", "nixpkgs", nixpkgsName)
	},
)

var nixUpdateCmd = shared.NewCommand(
	"update",
	"Update flakes",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		runNixCommand("nix", "flake", "update")
	},
)

var nixCleanCmd = shared.NewCommand(
	"clean",
	"Clean up Nix store",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		runNixCommand("sudo", "nix-collect-garbage", "-d")
	},
)

var nixLintCmd = shared.NewCommand(
	"lint",
	"Linter For Nix",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		runNixCommand("nix-shell", "-p", "alejandra", "--run", "alejandra .")
	},
)
