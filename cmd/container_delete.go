package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var containerDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a container",
	Long:  `Delete a container using podman rm.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		podmanCmd := exec.Command("podman", "rm", name)
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error deleting container:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerCmd.AddCommand(containerDeleteCmd)
}
