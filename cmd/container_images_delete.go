package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var containerImagesDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a container image",
	Long:  `Delete a container image using podman rmi.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		podmanCmd := exec.Command("podman", "rmi", id)
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error deleting image:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerImagesCmd.AddCommand(containerImagesDeleteCmd)
}
