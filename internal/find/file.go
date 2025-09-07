package find

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var FileCmd = &cobra.Command{
    Use:   "file [delete] <path> <terms...>",
    Short: "Search for files by name (find -type f -name)",
    Args:  cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        deleteMode := false
        if args[0] == "delete" {
            deleteMode = true
            args = args[1:]
        }
        path := args[0]
        terms := args[1:]

        for _, term := range terms {
            fmt.Printf("📄 Searching for files with '%s' in %s\n", term, path)
            cmd := exec.Command("find", path, "-type", "f", "-name", fmt.Sprintf("*%s*", term))
            matches, err := RunCommand(cmd)
            if err != nil {
                fmt.Printf("Error finding files with '%s': %v\n", term, err)
                continue
            }
            for _, match := range matches {
                fmt.Println(match)
            }

            if deleteMode && len(matches) > 0 && Confirm("⚠️  Delete matched files? (y/N): ") {
                for _, match := range matches {
                    DeletePath(match, false)
                }
            }
        }
    },
}

