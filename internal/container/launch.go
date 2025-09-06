package container

import (
	"github.com/spf13/cobra"
)

var tempShell bool
var silentMount bool
var silentUsb bool
var silentIp bool

var ContainerShellCmd = &cobra.Command{
	Use:   "shell [name]",
	Short: "Start a shell in a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman(BuildShellArgs(args[0], tempShell), true, "Error starting shell:")
	},
}

var ContainerSilentCmd = &cobra.Command{
	Use:   "silent [name]",
	Short: "Run a container in the background",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		RunPodman(BuildDetachedArgs(args[0], silentMount, silentUsb, silentIp), false, "Error running container in background:")
	},
}

func init() {
	ContainerShellCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
	ContainerShellCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerShellCmd.Flags().BoolVarP(&silentUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerShellCmd.Flags().BoolVarP(&silentIp, "ip", "", false, "Passthrough usbmuxd for iPhone")
	ContainerSilentCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	ContainerSilentCmd.Flags().BoolVarP(&silentUsb, "usb", "u", false, "Passthrough USB devices")
	ContainerSilentCmd.Flags().BoolVarP(&silentIp, "ip", "", false, "Passthrough usbmuxd for iPhone")
}
