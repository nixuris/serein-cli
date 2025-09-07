package nix

import (
	"github.com/spf13/cobra"
)

func init() {
	SysCmd.AddCommand(SysBuildCmd)
	SysCmd.AddCommand(SysGenCmd)
	SysCmd.AddCommand(SysGenDeleteCmd)
}

var SysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Manage NixOS system",
	Long:  `Manage NixOS system.`,
}

var SysBuildCmd = &cobra.Command{
	Use:   "build [path/to/flake]",
	Short: "Build a NixOS system",
	Long:  `Build a NixOS system with sudo nixos-rebuild switch --impure --flake.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flakePath := args[0]
		runNixCommand("sudo", "nixos-rebuild", "switch", "--impure", "--flake", flakePath)
	},
}

var SysGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "List system generations",
	Long:  `List system generations with sudo nix-env --list-generations --profile /nix/var/nix/profiles/system.`,
	Run: func(cmd *cobra.Command, args []string) {
		runNixCommand("sudo", "nix-env", "--list-generations", "--profile", "/nix/var/nix/profiles/system")
	},
}

var SysGenDeleteCmd = &cobra.Command{
	Use:   "delete [numbers...]",
	Short: "Delete system generations",
	Long:  `Delete system generations with sudo nix-env --profile /nix/var/nix/profiles/system --delete-generations. Can accept multiple numbers and ranges (e.g., 1-10).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generations := parseGenerations(args)
		cmdArgs := append([]string{"nix-env", "--profile", "/nix/var/nix/profiles/system", "--delete-generations"}, generations...)
		runNixCommand("sudo", cmdArgs...)
	},
}
