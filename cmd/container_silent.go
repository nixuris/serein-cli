package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var silentMount bool
var silentUsb bool
var silentIp bool

var containerSilentCmd = &cobra.Command{
	Use:   "silent [name]",
	Short: "Run a container in the background",
	Long:  `Run a container in the background.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		var podmanArgs []string
		podmanArgs = []string{"run", "-d"}

		if silentMount {
			podmanArgs = append(podmanArgs, "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")))
		}

		if silentUsb {
			podmanArgs = append(podmanArgs, "--device", "/dev/bus/usb")
		}

		if silentIp {
			podmanArgs = append(podmanArgs, "-v", "/var/run/usbmuxd:/var/run/usbmuxd")
		}

		podmanArgs = append(podmanArgs, name)

		podmanCmd := exec.Command("podman", podmanArgs...)
		podmanCmd.Stdout = os.Stdout
		podmanCmd.Stderr = os.Stderr

		if err := podmanCmd.Run(); err != nil {
			fmt.Println("Error running container in background:", err)
			os.Exit(1)
		}
	},
}

func init() {
	containerSilentCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	containerSilentCmd.Flags().BoolVarP(&silentUsb, "usb", "", false, "Passthrough USB devices")
	containerSilentCmd.Flags().BoolVarP(&silentIp, "ip", "", false, "Passthrough usbmuxd for iPhone")
	containerCmd.AddCommand(containerSilentCmd)
}
