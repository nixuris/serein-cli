package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var containerListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all containers",
	Long:  `List all containers using podman ps -a.`,
	Run: func(cmd *cobra.Command, args []string) {
		podmanCmd := exec.Command("podman", "ps", "-a")
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error listing containers:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerCmd.AddCommand(containerListCmd)
}
