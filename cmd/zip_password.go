package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var zipPasswordCmd = &cobra.Command{
	Use:   "password [archive-name] [target-to-archive]",
	Short: "Archive files with 7z and a password",
	Long:  `Archive files with 7z and a password.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		archiveName := args[0]
		targets := args[1:]

		for i, target := range targets {
			info, err := os.Stat(target)
			if err == nil && info.IsDir() {
				targets[i] = filepath.Join(target, "/*")
			}
		}

		fileExt := filepath.Ext(archiveName)

		var sevenZipCmd *exec.Cmd

		cmdArgs := []string{"a", "-p"}
		if fileExt == ".zip" {
			cmdArgs = append(cmdArgs, "-tzip")
		}
		cmdArgs = append(cmdArgs, archiveName)
		cmdArgs = append(cmdArgs, targets...)

		sevenZipCmd = exec.Command("7z", cmdArgs...)

		sevenZipCmd.Stdout = os.Stdout
		sevenZipCmd.Stderr = os.Stderr

		if err := sevenZipCmd.Run(); err != nil {
			fmt.Println("Error archiving with password:", err)
			os.Exit(1)
		}
	},
}

func init() {
	zipCmd.AddCommand(zipPasswordCmd)
}
