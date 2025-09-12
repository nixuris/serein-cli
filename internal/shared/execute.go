package shared

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string, args ...string) {
	if DryRun {
		fmt.Printf("Executing: %s %s\n", command, strings.Join(args, " "))
		return
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing '%s %s': %v\n", command, strings.Join(args, " "), err)
		os.Exit(1)
	}
}

func ExecuteCommandWithStdin(command string, args ...string) {
	if DryRun {
		fmt.Printf("Executing: %s %s\n", command, strings.Join(args, " "))
		return
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing '%s %s': %v\n", command, strings.Join(args, " "), err)
		os.Exit(1)
	}
}

func ExecuteCommandWithStderr(command string, args ...string) (string, error) {
	if DryRun {
		fmt.Printf("Executing: %s %s\n", command, strings.Join(args, " "))
		return "", nil
	}

	cmd := exec.Command(command, args...)
	var stderr strings.Builder
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stderr.String(), err
}

func ExecuteCommandWithOutput(command string, args ...string) (string, error) {
	if DryRun {
		fmt.Printf("Executing: %s %s\n", command, strings.Join(args, " "))
		return "", nil
	}

	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	return string(output), err
}

func RunCommand(command string, args ...string) ([]string, error) {
	output, err := ExecuteCommandWithOutput(command, args...)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var cleaned []string
	for _, line := range lines {
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}
	return cleaned, nil
}
