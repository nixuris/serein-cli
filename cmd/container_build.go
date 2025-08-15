package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var containerBuildCmd = &cobra.Command{
	Use:   "build [name] [path/to/dockerfile]",
	Short: "Build a container image",
	Long:  `Build a container image using podman.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		dockerfilePath := args[1]

		podmanCmd := exec.Command("podman", "build", "-t", name, dockerfilePath)
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error building container:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerCmd.AddCommand(containerBuildCmd)
}
