package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var tempShell bool

var containerShellCmd = &cobra.Command{
	Use:   "shell [name]",
	Short: "Start a shell in a container",
	Long:  `Start a shell in a container.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		var podmanArgs []string
		if tempShell {
			podmanArgs = []string{"run", "--rm", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), name, "/bin/bash"}
		} else {
			podmanArgs = []string{"run", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), name, "/bin/bash"}
		}

		podmanCmd := exec.Command("podman", podmanArgs...)
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr
		podmanCmd.Stdin = os.Stdin

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error starting shell:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerShellCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
	containerCmd.AddCommand(containerShellCmd)
}
