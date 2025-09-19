package box

import (
	"os/exec"
)

// DetectContainerEngine checks for docker and podman in the PATH and returns the preferred engine.
func DetectContainerEngine() string {
	if _, err := exec.LookPath("docker"); err == nil {
		return "docker"
	}
	if _, err := exec.LookPath("podman"); err == nil {
		return "podman"
	}
	return ""
}
