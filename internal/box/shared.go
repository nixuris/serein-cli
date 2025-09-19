package box

import (
	"fmt"
	"os"

	"serein/internal/shared"
)

var ContainerEngine string

func SetContainerEngine(useDocker, usePodman bool) {
	if useDocker {
		ContainerEngine = "docker"
	} else if usePodman {
		ContainerEngine = "podman"
	} else {
		ContainerEngine = DetectContainerEngine()
	}

	if ContainerEngine == "" {
		fmt.Println("Error: Neither docker nor podman found in your PATH. Please install one of them.")
		os.Exit(1)
	}
}

func RunContainerCommand(args []string, useStdin bool) {
	if useStdin {
		shared.CheckErr(shared.ExecuteCommandWithStdin(ContainerEngine, args...))
	} else {
		shared.CheckErr(shared.ExecuteCommand(ContainerEngine, args...))
	}
}

func MountCurrentDir() string {
	return fmt.Sprintf("%s:/mnt/", os.Getenv("PWD"))
}

func AppendMountFlags(base []string, mount, usb, ip bool) []string {
	if mount {
		base = append(base, "-v", MountCurrentDir())
	}
	if usb {
		base = append(base, "--device", "/dev/bus/usb")
	}
	if ip {
		base = append(base, "-v", "/var/run/usbmuxd:/var/run/usbmuxd")
	}
	return base
}

func BuildShellArgs(image string, temp, mount, usb, ip bool) []string {
	args := []string{"run"}
	if temp {
		args = append(args, "--rm")
	}
	args = append(args, "-it")
	args = AppendMountFlags(args, mount, usb, ip)
	args = append(args, image, "/bin/sh")
	return args
}

func BuildDetachedArgs(image string, mount, usb, ip bool) []string {
	args := []string{"run", "-d"}
	args = AppendMountFlags(args, mount, usb, ip)
	args = append(args, image)
	return args
}

func BuildIOSArgs(image string, pair bool) []string {
	args := []string{"run", "--rm", "-it", "-v", MountCurrentDir(), "-v", "/var/run/usbmuxd:/var/run/usbmuxd"}
	if pair {
		args = append(args, "--device", "/dev/bus/usb")
	}
	args = append(args, image)
	return args
}
