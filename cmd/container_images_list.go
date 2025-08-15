package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var containerImagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all container images",
	Long:  `List all container images using podman images.`,
	Run: func(cmd *cobra.Command, args []string) {
		podmanCmd := exec.Command("podman", "images")
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error listing images:", err)
			os.Exit(1)
		}
	},
}

var containerImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Manage container images",
	Long:  `Manage container images.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	containerImagesCmd.AddCommand(containerImagesListCmd)
	containerCmd.AddCommand(containerImagesCmd)
}
