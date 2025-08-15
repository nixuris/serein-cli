package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

func init() {
	nixHomeCmd.AddCommand(nixHomeGenCmd)
}
