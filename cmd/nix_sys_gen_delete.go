package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

func init() {
	nixSysGenCmd.AddCommand(nixSysGenDeleteCmd)
}
