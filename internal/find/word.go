package find

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var WordCmd = &cobra.Command{
    Use:   "word [delete] <path> <terms...>",
    Short: "Search for words inside files (grep -rE)",
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
            fmt.Printf("🔍 Searching for '%s' in %s\n", term, path)
            cmd := exec.Command("grep", "-rE", term, path)
            output, err := RunCommand(cmd)
            if err != nil {
                fmt.Printf("Error searching for '%s': %v\n", term, err)
                continue
            }
            for _, line := range output {
                fmt.Println(line)
            }

            if deleteMode && Confirm("⚠️  Delete matched files? (y/N): ") {
                DeleteGrepMatches(path, term)
            }
        }
    },
}

