package find

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var DirCmd = &cobra.Command{
    Use:   "dir [delete] <path> <terms...>",
    Short: "Search for directories by name (find -type d -name)",
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
            fmt.Printf("📁 Searching for directories with '%s' in %s\n", term, path)
            cmd := exec.Command("find", path, "-type", "d", "-name", fmt.Sprintf("*%s*", term))
            matches, err := RunCommand(cmd)
            if err != nil {
                fmt.Printf("Error finding directories with '%s': %v\n", term, err)
                continue
            }
            for _, match := range matches {
                fmt.Println(match)
            }

            if deleteMode && len(matches) > 0 && Confirm("⚠️  Delete matched directories? (y/N): ") {
                for _, match := range matches {
                    DeletePath(match, true)
                }
            }
        }
    },
}

