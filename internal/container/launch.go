package container

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var tempShell bool
var shellMount, shellUsb, shellIp bool

var silentMount, silentUsb, silentIp bool

var ContainerShellCmd = shared.NewCommand(
	"shell [name]",
	"Start a shell in a container",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman(BuildShellArgs(args[0], tempShell, shellMount, shellUsb, shellIp), true)
	},
)

var ContainerSilentCmd = shared.NewCommand(
	"silent [name]",
	"Run a container in the background",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		RunPodman(BuildDetachedArgs(args[0], silentMount, silentUsb, silentIp), false)
	},
)

func init() {
	ContainerShellCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
	ContainerShellCmd.Flags().BoolVarP(&shellMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerShellCmd.Flags().BoolVarP(&shellUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerShellCmd.Flags().BoolVarP(&shellIp, "ip", "", false, "Passthrough usbmuxd for iPhone")

	ContainerSilentCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerSilentCmd.Flags().BoolVarP(&silentUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerSilentCmd.Flags().BoolVarP(&silentIp, "ip", "", false, "Passthrough usbmuxd for iPhone")
}
