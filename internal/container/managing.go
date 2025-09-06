package container

import (
	"github.com/spf13/cobra"
)

var ContainerBuildCmd = &cobra.Command{
	Use:   "build [name] [path/to/dockerfile]",
	Short: "Build a container image",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"build", "-t", args[0], args[1]}, false, "Error building container:")
	},
}

var ContainerDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"rm", args[0]}, false, "Error deleting container:")
	},
}

var ContainerListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all containers",
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"ps", "-a"}, false, "Error listing containers:")
	},
}

var ContainerImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Manage container images",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var containerImagesDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a container image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"rmi", args[0]}, false, "Error deleting image:")
	},
}

var containerImagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all container images",
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"images"}, false, "Error listing images:")
	},
}

func init() {
	ContainerImagesCmd.AddCommand(containerImagesDeleteCmd)
	ContainerImagesCmd.AddCommand(containerImagesListCmd)
}
