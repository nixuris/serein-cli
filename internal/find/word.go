package find

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var WordCmd = &cobra.Command{
    Use:   "word <path> <terms...>",
    Short: "Search for words inside files",
    Args:  cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
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
        }
    },
}

var WordDeleteCmd = &cobra.Command{
    Use:   "delete <path> <terms...>",
    Short: "Delete files containing matching words",
    Args:  cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
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

            if Confirm("⚠️  Delete matched files? (y/N): ") {
                DeleteGrepMatches(path, term)
            }
        }
    },
}

func init() {
    WordCmd.AddCommand(WordDeleteCmd)
}

