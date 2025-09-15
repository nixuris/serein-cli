package container

import (
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"serein/internal/shared"
)

// Standalone Container Commands

var ContainerBuildCmd = shared.NewCommand(
	"build [name] [path/to/dockerfile]",
	"Build a container image",
	cobra.ExactArgs(2),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"build", "-t", args[0], args[1]}, false)
	},
)

var ContainerDeleteCmd = shared.NewCommand(
	"delete [name]",
	"Delete a container",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"rm", args[0]}, false)
	},
)

var ContainerListCmd = shared.NewCommand(
	"list",
	"List all containers",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"ps", "-a"}, false)
	},
)

var ContainerExportCmd = shared.NewCommand(
	"export <name>",
	"Export a container to a tarball",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		outputFile := containerName + ".tar"
		RunPodman([]string{"export", containerName, "-o", outputFile}, false)
	},
)

var ContainerImportCmd = shared.NewCommand(
	"import <path-to-tar>",
	"Import a container from a tarball",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		tarPath := args[0]
		baseName := filepath.Base(tarPath)
		containerName := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		RunPodman([]string{"import", tarPath, containerName}, false)
	},
)

// Container Images Commands

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
		RunPodman([]string{"rmi", args[0]}, false)
	},
)

var containerImagesListCmd = shared.NewCommand(
	"list",
	"List all container images",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"images"}, false)
	},
)

var containerImagesExportCmd = shared.NewCommand(
	"export <image-name>",
	"Export an image to a tarball",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		// Sanitize image name for use as a filename
		safeName := strings.ReplaceAll(imageName, "/", "_")
		safeName = strings.ReplaceAll(safeName, ":", "-")
		outputFile := safeName + ".tar"
		RunPodman([]string{"save", "-o", outputFile, imageName}, false)
	},
)

var containerImagesImportCmd = shared.NewCommand(
	"import <path-to-tar>",
	"Import an image from a tarball",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman([]string{"load", "-i", args[0]}, false)
	},
)

func StandaloneFlags(parent *cobra.Command) {
	parent.AddCommand(ContainerBuildCmd)
	parent.AddCommand(ContainerDeleteCmd)
	parent.AddCommand(ContainerListCmd)
	parent.AddCommand(ContainerExportCmd)
	parent.AddCommand(ContainerImportCmd)
}

func init() {
	ContainerImagesCmd.AddCommand(containerImagesDeleteCmd)
	ContainerImagesCmd.AddCommand(containerImagesListCmd)
	ContainerImagesCmd.AddCommand(containerImagesExportCmd)
	ContainerImagesCmd.AddCommand(containerImagesImportCmd)
}
