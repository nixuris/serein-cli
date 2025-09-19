package box

import (
	"strings"

	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var iosSidestore bool
var iosPair bool

var ContainerIosCmd = shared.NewCommand(
	"ios [args...]",
	"Manage iOS devices with containers with optional arguments: sidestore, pair",
	cobra.ArbitraryArgs,
	func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			switch strings.ToLower(arg) {
			case "sidestore":
				iosSidestore = true
			case "pair":
				iosPair = true
			}
		}

		if iosSidestore {
			RunContainerCommand(BuildIOSArgs("ghcr.io/sidestore/altcon", false), true)
		} else if iosPair {
			RunContainerCommand(BuildIOSArgs("ipairing", true), true)
		} else {
			_ = cmd.Help()
		}
	},
)

func init() {
	ContainerIosCmd.Flags().BoolVarP(&iosSidestore, "sidestore", "s", false, "Run sidestore container")
	ContainerIosCmd.Flags().BoolVarP(&iosPair, "pair", "p", false, "Run ipairing container")
}
