package cmd

import (
	"github.com/spf13/cobra"
	"serein/internal/container"
)

func init() {
	rootCmd.AddCommand(containerCmd)

	container.StandaloneFlags(containerCmd)
	containerCmd.AddCommand(container.ContainerImagesCmd)
	containerCmd.AddCommand(container.ContainerIosCmd)
	containerCmd.AddCommand(container.ContainerShellCmd)
	containerCmd.AddCommand(container.ContainerSilentCmd)
}

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage containers with podman aliases",
	Long:  `Manage containers with podman aliases.`,
}
