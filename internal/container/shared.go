package container

import (
	"fmt"
	"os"

	"serein/internal/execute"
)

func RunPodman(args []string, useStdin bool, failMsg string) {
	if useStdin {
		execute.ExecuteCommandWithStdin("podman", args...)
	} else {
		execute.ExecuteCommand("podman", args...)
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

func BuildShellArgs(image string, temp bool) []string {
	args := []string{"run"}
	if temp {
		args = append(args, "--rm")
	}
	args = append(args, "-it", "-v", MountCurrentDir(), image, "/bin/bash")
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
