package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var iosSidestore bool
var iosPair bool

var containerIosCmd = &cobra.Command{
	Use:   "ios",
	Short: "Manage iOS devices with containers",
	Long:  `Manage iOS devices with containers for sidestore and pairing.`,
	Run: func(cmd *cobra.Command, args []string) {
		if iosSidestore {
			podmanCmd := exec.Command("podman", "run", "--rm", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), "-v", "/var/run/usbmuxd:/var/run/usbmuxd", "ghcr.io/sidestore/altcon")
			podmanCmd.Stdout = os.Stdout
			podmanCmd.Stderr = os.Stderr
			podmanCmd.Stdin = os.Stdin

			if err := podmanCmd.Run(); err != nil {
				fmt.Println("Error running sidestore container:", err)
				os.Exit(1)
			}
		} else if iosPair {
			podmanCmd := exec.Command("podman", "run", "--rm", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), "-v", "/var/run/usbmuxd:/var/run/usbmuxd", "--device", "/dev/bus/usb", "ipairing")
			podmanCmd.Stdout = os.Stdout
			podmanCmd.Stderr = os.Stderr
			podmanCmd.Stdin = os.Stdin

			if err := podmanCmd.Run(); err != nil {
				fmt.Println("Error running ipairing container:", err)
				os.Exit(1)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	containerIosCmd.Flags().BoolVarP(&iosSidestore, "sidestore", "s", false, "Run sidestore container")
	containerIosCmd.Flags().BoolVarP(&iosPair, "pair", "p", false, "Run ipairing container")
	containerCmd.AddCommand(containerIosCmd)
}
