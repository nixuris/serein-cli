package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"serein/internal/git"
)
func init() {

	rootCmd.AddCommand(gitCmd)
	git.RegisterBasicGitCommands(gitCmd)
        
	gitCmd.AddCommand(git.TagCmd)
        gitCmd.AddCommand(git.BranchCmd)
        gitCmd.AddCommand(git.CommitCmd)

}

func runGitCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Git helper commands",
	Long:  `A set of helper commands to simplify common Git operations.`,
}
