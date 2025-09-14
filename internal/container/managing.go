package container

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var ContainerBuildCmd = shared.NewCommand(
	"build [name] [path/to/dockerfile]",
	"Build a container image",
	cobra.ExactArgs(2),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"build", "-t", args[0], args[1]}, false, "Error building container:")
	},
)

var ContainerDeleteCmd = shared.NewCommand(
	"delete [name]",
	"Delete a container",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"rm", args[0]}, false, "Error deleting container:")
	},
)

var ContainerListCmd = shared.NewCommand(
	"list",
	"List all containers",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"ps", "-a"}, false, "Error listing containers:")
	},
)

var ContainerImagesCmd = shared.NewCommand(
	"images",
	"Manage container images",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
)

var containerImagesDeleteCmd = shared.NewCommand(
	"delete [id]",
	"Delete a container image",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"rmi", args[0]}, false, "Error deleting image:")
	},
)

var containerImagesListCmd = shared.NewCommand(
	"list",
	"List all container images",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"images"}, false, "Error listing images:")
	},
)

func StandaloneFlags(parent *cobra.Command) {
	parent.AddCommand(ContainerBuildCmd)
	parent.AddCommand(ContainerDeleteCmd)
	parent.AddCommand(ContainerListCmd)
}

func init() {
	ContainerImagesCmd.AddCommand(containerImagesDeleteCmd)
	ContainerImagesCmd.AddCommand(containerImagesListCmd)
}
