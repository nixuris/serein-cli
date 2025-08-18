package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", args...)
	},
}

var gitSyncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Alias for git pull origin <branch>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "pull", "origin", args[0])
	},
}

var gitStageCmd = &cobra.Command{
	Use:   "stage [path]",
	Short: "Alias for git add <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "add", args[0])
	},
}

var gitUnstageCmd = &cobra.Command{
	Use:   "unstage [path]",
	Short: "Alias for git reset <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "reset", args[0])
	},
}

var gitUndoCmd = &cobra.Command{
	Use:   "undo [path]",
	Short: "Alias for git restore <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "restore", args[0])
	},
}

var gitChangesCmd = &cobra.Command{
	Use:   "changes [path]",
	Short: "Alias for git diff <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "diff", args[0])
	},
}

var gitStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Alias for git status",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "status")
	},
}

var gitTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Git tag commands",
}

var gitTagCreateCmd = &cobra.Command{
	Use:   "create [SHA] [name] [message]",
	Short: "Alias for git tag -a <name> -m \"<msg>\" <SHA>",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "tag", "-a", args[1], "-m", args[2], args[0])
	},
}

var gitTagDeleteLocalCmd = &cobra.Command{
	Use:   "local [name]",
	Short: "Alias for git tag -d <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "tag", "-d", args[0])
	},
}

var gitTagDeleteRemoteCmd = &cobra.Command{
	Use:   "remote [name]",
	Short: "Alias for git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "push", "origin", "--delete", args[0])
	},
}

var gitTagWipeCmd = &cobra.Command{
	Use:   "wipe [name]",
	Short: "Alias for git tag -d <name> && git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "tag", "-d", args[0])
		runGitCommand("git", "push", "origin", "--delete", args[0])
	},
}

var gitCommitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Alias for git commit -m \"<commit msg>\"",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "commit", "-m", args[0])
	},
}

var gitCommitPushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Alias for git push origin <branch>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "push", "origin", args[0])
	},
}

var gitCommitListCmd = &cobra.Command{
	Use:   "list",
	Short: "Alias for git log --oneline",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "log", "--oneline")
	},
}

var gitCommitUndoCmd = &cobra.Command{
	Use:   "undo [SHA]",
	Short: "Alias for git revert <SHA>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "revert", args[0])
	},
}

var gitCommitDeleteCmd = &cobra.Command{
	Use:   "delete [number]",
	Short: "Alias for git reset --hard HEAD~<numb>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "reset", "--hard", "HEAD~"+args[0])
	},
}

var gitCommitChangesCmd = &cobra.Command{
	Use:   "changes [SHA]",
	Short: "Alias for git show <SHA>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "show", args[0])
	},
}

var gitCommitCompareCmd = &cobra.Command{
	Use:   "compare [SHA1] [SHA2]",
	Short: "Alias for git diff <SHA1> <SHA2>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "diff", args[0], args[1])
	},
}

var gitBranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Git branch commands",
}

var gitBranchListCmd = &cobra.Command{
	Use:   "list",
	Short: "Alias for git branch -avv",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "branch", "-avv")
	},
}

var gitBranchCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Alias for git switch -c <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "switch", "-c", args[0])
	},
}

var gitBranchSwitchCmd = &cobra.Command{
	Use:   "switch [name]",
	Short: "Alias for git switch <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "switch", args[0])
	},
}

var gitBranchDeleteLocalCmd = &cobra.Command{
	Use:   "local [name]",
	Short: "Alias for git branch -D <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "branch", "-D", args[0])
	},
}

var gitBranchDeleteRemoteCmd = &cobra.Command{
	Use:   "remote [name]",
	Short: "Alias for git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("git", "push", "origin", "--delete", args[0])
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	gitCmd.AddCommand(gitSyncCmd)
	gitCmd.AddCommand(gitStageCmd)
	gitCmd.AddCommand(gitUnstageCmd)
	gitCmd.AddCommand(gitUndoCmd)
	gitCmd.AddCommand(gitChangesCmd)
	gitCmd.AddCommand(gitStatusCmd)

	gitCmd.AddCommand(gitTagCmd)
	gitTagCmd.AddCommand(gitTagCreateCmd)
	gitTagCmd.AddCommand(gitTagDeleteLocalCmd)
	gitTagCmd.AddCommand(gitTagDeleteRemoteCmd)
	gitTagCmd.AddCommand(gitTagWipeCmd)

	gitCmd.AddCommand(gitCommitCmd)
	gitCommitCmd.AddCommand(gitCommitPushCmd)
	gitCommitCmd.AddCommand(gitCommitListCmd)
	gitCommitCmd.AddCommand(gitCommitUndoCmd)
	gitCommitCmd.AddCommand(gitCommitDeleteCmd)
	gitCommitCmd.AddCommand(gitCommitChangesCmd)
	gitCommitCmd.AddCommand(gitCommitCompareCmd)

	gitCmd.AddCommand(gitBranchCmd)
	gitBranchCmd.AddCommand(gitBranchListCmd)
	gitBranchCmd.AddCommand(gitBranchCreateCmd)
	gitBranchCmd.AddCommand(gitBranchSwitchCmd)
	gitBranchCmd.AddCommand(gitBranchDeleteLocalCmd)
	gitBranchCmd.AddCommand(gitBranchDeleteRemoteCmd)
}
