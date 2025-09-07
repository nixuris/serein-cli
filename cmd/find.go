package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "strings"

    "github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
    Use:   "find",
    Short: "Search for words, files, or directories",
    Long: `Simple wrapper for grep and find.

Examples:
  serein find word . "kitty"
  serein find word delete . "kitty"
  serein find file delete ./docs "readme"
  serein find dir delete ./assets "images"`,
}

func init() {
    rootCmd.AddCommand(findCmd)
    findCmd.AddCommand(findWordCmd)
    findCmd.AddCommand(findFileCmd)
    findCmd.AddCommand(findDirCmd)
}

var findWordCmd = &cobra.Command{
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
            grepCmd := exec.Command("grep", "-rE", term, path)
            output, err := grepCmd.Output()
            if err != nil {
                fmt.Printf("Error searching for '%s': %v\n", term, err)
                continue
            }
            fmt.Print(string(output))

            if deleteMode {
                fmt.Print("⚠️  Delete matched files? (y/N): ")
                var response string
                fmt.Scanln(&response)
                if strings.ToLower(response) == "y" {
                    deleteMatchedFilesFromGrep(path, term)
                }
            }
        }
    },
}

var findFileCmd = &cobra.Command{
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
            findCmd := exec.Command("find", path, "-type", "f", "-name", fmt.Sprintf("*%s*", term))
            output, err := findCmd.Output()
            if err != nil {
                fmt.Printf("Error finding files with '%s': %v\n", term, err)
                continue
            }
            matches := strings.Split(strings.TrimSpace(string(output)), "\n")
            for _, match := range matches {
                if match != "" {
                    fmt.Println(match)
                }
            }

            if deleteMode && len(matches) > 0 {
                fmt.Print("⚠️  Delete matched files? (y/N): ")
                var response string
                fmt.Scanln(&response)
                if strings.ToLower(response) == "y" {
                    for _, match := range matches {
                        if match != "" {
                            deletePath(match, false)
                        }
                    }
                }
            }
        }
    },
}

var findDirCmd = &cobra.Command{
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
            findCmd := exec.Command("find", path, "-type", "d", "-name", fmt.Sprintf("*%s*", term))
            output, err := findCmd.Output()
            if err != nil {
                fmt.Printf("Error finding directories with '%s': %v\n", term, err)
                continue
            }
            matches := strings.Split(strings.TrimSpace(string(output)), "\n")
            for _, match := range matches {
                if match != "" {
                    fmt.Println(match)
                }
            }

            if deleteMode && len(matches) > 0 {
                fmt.Print("⚠️  Delete matched directories? (y/N): ")
                var response string
                fmt.Scanln(&response)
                if strings.ToLower(response) == "y" {
                    for _, match := range matches {
                        if match != "" {
                            deletePath(match, true)
                        }
                    }
                }
            }
        }
    },
}

func deletePath(path string, isDir bool) {
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

func deleteMatchedFilesFromGrep(basePath, term string) {
    grepCmd := exec.Command("grep", "-rlE", term, basePath)
    output, err := grepCmd.Output()
    if err != nil {
        fmt.Printf("Error finding files to delete: %v\n", err)
        return
    }
    matches := strings.Split(strings.TrimSpace(string(output)), "\n")
    for _, match := range matches {
        if match != "" {
            deletePath(match, false)
        }
    }
}
