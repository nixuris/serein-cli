package box

import (
	"strings"

	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var tempShell bool
var shellMount, shellUsb, shellIp bool

var silentMount, silentUsb, silentIp bool

var ContainerShellCmd = shared.NewCommand(
	"shell [args...] [name]",
	"Start a shell in a container with optional arguments: temp, mount, usb, ip",
	cobra.MinimumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		imageName := args[len(args)-1]
		for _, arg := range args[:len(args)-1] {
			switch strings.ToLower(arg) {
			case "temp":
				tempShell = true
			case "mount":
				shellMount = true
			case "usb":
				shellUsb = true
			case "ip":
				shellIp = true
			}
		}
		RunContainerCommand(BuildShellArgs(imageName, tempShell, shellMount, shellUsb, shellIp), true)
	},
)

var ContainerSilentCmd = shared.NewCommand(
	"silent [args...] [name]",
	"Run a container in the background with optional arguments: mount, usb, ip",
	cobra.MinimumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		imageName := args[len(args)-1]
		for _, arg := range args[:len(args)-1] {
			switch strings.ToLower(arg) {
			case "mount":
				silentMount = true
			case "usb":
				silentUsb = true
			case "ip":
				silentIp = true
			}
		}
		RunContainerCommand(BuildDetachedArgs(imageName, silentMount, silentUsb, silentIp), false)
	},
)

func init() {
	ContainerShellCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
	ContainerShellCmd.Flags().BoolVarP(&shellMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerShellCmd.Flags().BoolVarP(&shellUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerShellCmd.Flags().BoolVar(&shellIp, "ip", false, "Passthrough usbmuxd for iPhone")

	ContainerSilentCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerSilentCmd.Flags().BoolVarP(&silentUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerSilentCmd.Flags().BoolVar(&silentIp, "ip", false, "Passthrough usbmuxd for iPhone")
}
