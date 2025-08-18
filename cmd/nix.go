package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nixCmd = &cobra.Command{
	Use:   "nix",
	Short: "Nix related commands",
	Long:  `A collection of commands to manage Nix and NixOS systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

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

var nixHomeCmd = &cobra.Command{
	Use:   "home",
	Short: "Manage home-manager",
	Long:  `Manage home-manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

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

var nixHomeGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "List home-manager generations",
	Long:  `List home-manager generations with home-manager generations.`,
	Run: func(cmd *cobra.Command, args []string) {
		nixCmd := exec.Command("home-manager", "generations")
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error listing home-manager generations:", err)
			os.Exit(1)
		}
	},
}

var nixHomeGenDeleteCmd = &cobra.Command{
	Use:   "delete [number]",
	Short: "Delete home-manager generations",
	Long:  `Delete home-manager generations with home-manager remove-generations.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generation := args[0]

		nixCmd := exec.Command("home-manager", "remove-generations", generation)
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error deleting home-manager generations:", err)
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

var nixSysGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "List system generations",
	Long:  `List system generations with sudo nix-env --list-generations --profile /nix/var/nix/profiles/system.`,
	Run: func(cmd *cobra.Command, args []string) {
		nixCmd := exec.Command("sudo", "nix-env", "--list-generations", "--profile", "/nix/var/nix/profiles/system")
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error listing system generations:", err)
			os.Exit(1)
		}
	},
}

var nixSysGenDeleteCmd = &cobra.Command{
	Use:   "delete [number]",
	Short: "Delete system generations",
	Long:  `Delete system generations with sudo nix-env --profile /nix/var/nix/profiles/system --delete-generations.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generation := args[0]

		nixCmd := exec.Command("sudo", "nix-env", "--profile", "/nix/var/nix/profiles/system", "--delete-generations", generation)
		nixCmd.Stdout = os.Stdout
		nixCmd.Stderr = os.Stderr

		if err := nixCmd.Run(); err != nil {
			fmt.Println("Error deleting system generations:", err)
			os.Exit(1)
		}
	},
}

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
	rootCmd.AddCommand(nixCmd)
	nixCmd.AddCommand(nixCleanCmd)
	nixCmd.AddCommand(nixHomeCmd)
	nixHomeCmd.AddCommand(nixHomeBuildCmd)
	nixHomeCmd.AddCommand(nixHomeGenCmd)
	nixHomeGenCmd.AddCommand(nixHomeGenDeleteCmd)
	nixCmd.AddCommand(nixSysCmd)
	nixSysCmd.AddCommand(nixSysBuildCmd)
	nixSysCmd.AddCommand(nixSysGenCmd)
	nixSysGenCmd.AddCommand(nixSysGenDeleteCmd)
	nixCmd.AddCommand(nixUpdateCmd)
}
