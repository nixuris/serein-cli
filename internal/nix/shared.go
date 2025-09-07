package nix

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runNixCommand(command string, args ...string) {
	fmt.Printf("Executing: %s %v\n", command, args)
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing '%s %v': %v\n", command, args, err)
		os.Exit(1)
	}
}

// parseGenerations takes a slice of strings, which can contain numbers or ranges (e.g., "1-5"),
// and returns a slice of strings with all the individual numbers.
func parseGenerations(args []string) []string {
	var generations []string
	for _, arg := range args {
		// Check if the argument contains a hyphen, indicating a range.
		if strings.Contains(arg, "-") {
			parts := strings.Split(arg, "-")
			if len(parts) != 2 {
				fmt.Printf("Invalid range format: %s\n", arg)
				os.Exit(1)
			}
			// Convert the start and end of the range to integers.
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Printf("Invalid range numbers: %s\n", arg)
				os.Exit(1)
			}
			// Ensure the start of the range is not greater than the end.
			if start > end {
				fmt.Printf("Start of range cannot be greater than end: %s\n", arg)
				os.Exit(1)
			}
			// Append all numbers in the range to the generations slice.
			for i := start; i <= end; i++ {
				generations = append(generations, strconv.Itoa(i))
			}
		} else {
			// If it's not a range, just append the number.
			generations = append(generations, arg)
		}
	}
	return generations
}
