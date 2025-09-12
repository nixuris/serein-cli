package shared

import (
	"fmt"
	"os"
)

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
