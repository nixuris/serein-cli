package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
    Use:   "container",
    Short: "Manage containers with podman aliases",
    Long:  `Manage containers with podman aliases.`,
    Run: func(cmd *cobra.Command, args []string) {
        if err := cmd.Help(); err != nil {
            fmt.Println("Error showing help:", err)
        }
    },
}

var containerBuildCmd = &cobra.Command{
    Use:   "build [name] [path/to/dockerfile]",
    Short: "Build a container image",
    Long:  `Build a container image using podman.`,
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        dockerfilePath := args[1]

        podmanCmd := exec.Command("podman", "build", "-t", name, dockerfilePath)
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error building container:", err)
            os.Exit(1)
        }
    },
}

var containerDeleteCmd = &cobra.Command{
    Use:   "delete [name]",
    Short: "Delete a container",
    Long:  `Delete a container using podman rm.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]

        podmanCmd := exec.Command("podman", "rm", name)
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error deleting container:", err)
            os.Exit(1)
        }
    },
}

var containerImagesCmd = &cobra.Command{
    Use:   "images",
    Short: "Manage container images",
    Long:  `Manage container images.`,
    Run: func(cmd *cobra.Command, args []string) {
        if err := cmd.Help(); err != nil {
            fmt.Println("Error showing help:", err)
        }
    },
}

var containerImagesDeleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Delete a container image",
    Long:  `Delete a container image using podman rmi.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0]

        podmanCmd := exec.Command("podman", "rmi", id)
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error deleting image:", err)
            os.Exit(1)
        }
    },
}

var containerImagesListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all container images",
    Long:  `List all container images using podman images.`,
    Run: func(cmd *cobra.Command, args []string) {
        podmanCmd := exec.Command("podman", "images")
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error listing images:", err)
            os.Exit(1)
        }
    },
}

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
            if err := cmd.Help(); err != nil {
                fmt.Println("Error showing help:", err)
            }
        }
    },
}

var containerListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all containers",
    Long:  `List all containers using podman ps -a.`,
    Run: func(cmd *cobra.Command, args []string) {
        podmanCmd := exec.Command("podman", "ps", "-a")
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error listing containers:", err)
            os.Exit(1)
        }
    },
}

var tempShell bool

var containerShellCmd = &cobra.Command{
    Use:   "shell [name]",
    Short: "Start a shell in a container",
    Long:  `Start a shell in a container.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]

        var podmanArgs []string
        if tempShell {
            podmanArgs = []string{"run", "--rm", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), name, "/bin/bash"}
        } else {
            podmanArgs = []string{"run", "-it", "-v", fmt.Sprintf("%s:/mnt/", os.Getenv("PWD")), name, "/bin/bash"}
        }

        podmanCmd := exec.Command("podman", podmanArgs...)
        podmanCmd.Stdout = os.Stdout
        podmanCmd.Stderr = os.Stderr
        podmanCmd.Stdin = os.Stdin

        if err := podmanCmd.Run(); err != nil {
            fmt.Println("Error starting shell:", err)
            os.Exit(1)
        }
    },
}

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
    rootCmd.AddCommand(containerCmd)
    containerCmd.AddCommand(containerBuildCmd)
    containerCmd.AddCommand(containerDeleteCmd)
    containerCmd.AddCommand(containerImagesCmd)
    containerImagesCmd.AddCommand(containerImagesDeleteCmd)
    containerImagesCmd.AddCommand(containerImagesListCmd)
    containerCmd.AddCommand(containerIosCmd)
    containerIosCmd.Flags().BoolVarP(&iosSidestore, "sidestore", "s", false, "Run sidestore container")
    containerIosCmd.Flags().BoolVarP(&iosPair, "pair", "p", false, "Run ipairing container")
    containerCmd.AddCommand(containerListCmd)
    containerCmd.AddCommand(containerShellCmd)
    containerShellCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
    containerCmd.AddCommand(containerSilentCmd)
    containerSilentCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
    containerSilentCmd.Flags().BoolVarP(&silentUsb, "usb", "", false, "Passthrough USB devices")
    containerSilentCmd.Flags().BoolVarP(&silentIp, "ip", "", false, "Passthrough usbmuxd for iPhone")
}
