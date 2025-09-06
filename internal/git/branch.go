package git

import (
	"github.com/spf13/cobra"
)

func init() {
	BranchCmd.AddCommand(gitBranchListCmd)
	BranchCmd.AddCommand(gitBranchCreateCmd)
	BranchCmd.AddCommand(gitBranchSwitchCmd)
	BranchCmd.AddCommand(gitBranchDeleteLocalCmd)
	BranchCmd.AddCommand(gitBranchDeleteRemoteCmd)
}

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Git branch commands",
}

var gitBranchListCmd = &cobra.Command{
	Use:   "list",
	Short: "Alias for git branch -avv",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("branch", "-avv")
	},
}

var gitBranchCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Alias for git switch -c <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("switch", "-c", args[0])
	},
}

var gitBranchSwitchCmd = &cobra.Command{
	Use:   "switch [name]",
	Short: "Alias for git switch <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("switch", args[0])
	},
}

var gitBranchDeleteLocalCmd = &cobra.Command{
	Use:   "local [name]",
	Short: "Alias for git branch -D <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("branch", "-D", args[0])
	},
}

var gitBranchDeleteRemoteCmd = &cobra.Command{
	Use:   "remote [name]",
	Short: "Alias for git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("push", "origin", "--delete", args[0])
	},
}
