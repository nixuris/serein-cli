package nix

import (
	"fmt"
	"os"
	"os/exec"
)

func runNixCommand(command string, args ...string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error executing '%s %v': %v\n", command, args, err)
        os.Exit(1)
    }
}
