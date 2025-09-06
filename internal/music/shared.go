package music

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCommand executes a command and streams output to stdout/stderr
func RunCommand(cmd *exec.Cmd, failMsg string) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(failMsg, err)
		return err
	}
	return nil
}

// RunCommandWithStderr captures stderr output for logging
func RunCommandWithStderr(cmd *exec.Cmd) (string, error) {
	var stderr strings.Builder
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stderr.String(), err
}

// CreateFile safely creates a file and returns the handle
func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil, err
	}
	return file, nil
}

// OpenFile safely opens a file and returns the handle
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	return file, nil
}

// CloseFile safely closes a file and logs any error
func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println("Error closing file:", err)
	}
}

// LogError writes a message to a log file
func LogError(file *os.File, message string) {
	if _, err := fmt.Fprintln(file, message); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

// FormatPlaylistLines transforms playlist lines to Winamp/Ruizu-safe format
func FormatPlaylistLines(lines []string) []string {
	var formatted []string
	for _, line := range lines {
		line = strings.ReplaceAll(line, "/", "\\") + "\r"
		formatted = append(formatted, line)
	}
	return formatted
}
