package container

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var iosSidestore bool
var iosPair bool

var ContainerIosCmd = shared.NewCommand(
	"ios",
	"Manage iOS devices with containers",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		switch {
		case iosSidestore:
			RunPodman(BuildIOSArgs("ghcr.io/sidestore/altcon", false), true, "Error running sidestore container:")
		case iosPair:
			RunPodman(BuildIOSArgs("ipairing", true), true, "Error running ipairing container:")
		default:
			_ = cmd.Help()
		}
	},
)

func init() {
	ContainerIosCmd.Flags().BoolVarP(&iosSidestore, "sidestore", "s", false, "Run sidestore container")
	ContainerIosCmd.Flags().BoolVarP(&iosPair, "pair", "p", false, "Run ipairing container")
}
