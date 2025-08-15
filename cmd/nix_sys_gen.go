package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

func init() {
	nixSysCmd.AddCommand(nixSysGenCmd)
}
