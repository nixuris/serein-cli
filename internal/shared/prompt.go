package shared

import (
	"fmt"
	"strings"
)

func Confirm(prompt string) bool {
	fmt.Print(prompt)
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		fmt.Printf("⚠️  Failed to read input: %v\n", err)
		return false
	}
	return strings.ToLower(response) == "y"
}
