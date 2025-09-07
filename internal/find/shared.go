package find

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DeletePath(path string, isDir bool) {
	var err error
	if isDir {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}
	if err != nil {
		fmt.Printf("❌ Failed to delete %s: %v\n", path, err)
	} else {
		fmt.Printf("✅ Deleted: %s\n", path)
	}
}

func Confirm(prompt string) bool {
	fmt.Print(prompt)
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		fmt.Printf("⚠️  Failed to read input: %v\n", err)
		return false
	}
	return strings.ToLower(response) == "y"
}

func RunCommand(cmd *exec.Cmd) ([]string, error) {
	output, err := cmd.Output()
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

func DeleteGrepMatches(basePath, term string) {
	cmd := exec.Command("grep", "-rlE", term, basePath)
	matches, err := RunCommand(cmd)
	if err != nil {
		fmt.Printf("Error finding files to delete: %v\n", err)
		return
	}
	for _, match := range matches {
		DeletePath(match, false)
	}
}
