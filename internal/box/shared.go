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

func BuildShellCreateArgs(image string, temp, mount, usb, ip bool, name, shell string) []string {
	args := []string{"run"}
	if temp {
		args = append(args, "--rm")
	}
	if name != "" {
		args = append(args, "--name", name)
	}
	args = append(args, "-it")
	args = AppendMountFlags(args, mount, usb, ip)
	args = append(args, image, "/bin/"+shell)
	return args
}

func RunShellWithWatcher(container string, shell string) {
	fmt.Printf("Starting shell with auto-stop watcher...\n")

	// Start the exec command
	execArgs := BuildShellResumeArgs(container, shell)

	// Run the exec command and wait for it to finish
	err := shared.ExecuteCommandWithStdin(ContainerEngine, execArgs...)

	// Debug: always try to stop, regardless of error
	fmt.Printf("Shell exited, stopping container %s...\n", container)
	stopArgs := []string{"stop", container}
	stopErr := shared.ExecuteCommand(ContainerEngine, stopArgs...)

	if stopErr != nil {
		fmt.Printf("Warning: failed to stop container: %v\n", stopErr)
	} else {
		fmt.Printf("Container %s stopped successfully\n", container)
	}

	// Check the original exec error
	shared.CheckErr(err)
}

func BuildShellResumeArgs(container string, shell string) []string {
	// For exec mode: exec into running container
	return []string{"exec", "-it", container, "/bin/" + shell}
}

func BuildShellAttachArgs(container string) []string {
	// For default mode: attach to container (stops when you exit)
	return []string{"start", "-ai", container}
}

func BuildDetachedCreateArgs(image string, mount, usb, ip bool, name string) []string {
	args := []string{"run", "-d"}
	if name != "" {
		args = append(args, "--name", name)
	}
	args = AppendMountFlags(args, mount, usb, ip)
	args = append(args, image)
	return args
}

func BuildDetachedResumeArgs(container string) []string {
	return []string{"start", container}
}

func BuildIOSArgs(image string, pair bool) []string {
	args := []string{"run", "--rm", "-it", "-v", MountCurrentDir(), "-v", "/var/run/usbmuxd:/var/run/usbmuxd"}
	if pair {
		args = append(args, "--device", "/dev/bus/usb")
	}
	args = append(args, image)
	return args
}
