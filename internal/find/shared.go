package find

import (
	"fmt"
	"os"
	"strings"

	"serein/internal/execute"
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

func RunCommand(command string, args ...string) ([]string, error) {
	output, err := execute.ExecuteCommandWithOutput(command, args...)
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
	matches, err := RunCommand("grep", "-rlE", term, basePath)
	if err != nil {
		fmt.Printf("Error finding files to delete: %v\n", err)
		return
	}
	for _, match := range matches {
		DeletePath(match, false)
	}
}
