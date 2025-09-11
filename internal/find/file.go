package find

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FileCmd = &cobra.Command{
	Use:   "file <path> <terms...>",
	Short: "Search for files by name",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		terms := args[1:]

		for _, term := range terms {
			fmt.Printf("📄 Searching for files with '%s' in %s\n", term, path)
			matches, err := RunCommand("find", path, "-type", "f", "-name", fmt.Sprintf("*%s*", term))
			if err != nil {
				fmt.Printf("Error finding files with '%s': %v\n", term, err)
				continue
			}
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	},
}

var FileDeleteCmd = &cobra.Command{
	Use:   "delete <path> <terms...>",
	Short: "Delete files by name",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		terms := args[1:]

		for _, term := range terms {
			fmt.Printf("📄 Searching for files with '%s' in %s\n", term, path)
			matches, err := RunCommand("find", path, "-type", "f", "-name", fmt.Sprintf("*%s*", term))
			if err != nil {
				fmt.Printf("Error finding files with '%s': %v\n", term, err)
				continue
			}
			for _, match := range matches {
				fmt.Println(match)
			}

			if len(matches) > 0 && Confirm("⚠️  Delete matched files? (y/N): ") {
				for _, match := range matches {
					DeletePath(match, false)
				}
			}
		}
	},
}

func init() {
	FileCmd.AddCommand(FileDeleteCmd)
}
