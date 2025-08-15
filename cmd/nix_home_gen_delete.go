package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

func init() {
	nixHomeGenCmd.AddCommand(nixHomeGenDeleteCmd)
}
