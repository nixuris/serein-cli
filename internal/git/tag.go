package git

import (
	"github.com/spf13/cobra"
)

func init() {
    TagCmd.AddCommand(tagDeleteLocalCmd)
    TagCmd.AddCommand(tagCreateCmd)
    TagCmd.AddCommand(tagWipeCmd)
    TagCmd.AddCommand(tagDeleteRemoteCmd)
}

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Git tag commands",
}

var tagCreateCmd = &cobra.Command{
	Use:   "create [SHA] [name] [message]",
	Short: "Alias for git tag -a <name> -m \"<msg>\" <SHA>",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("tag", "-a", args[1], "-m", args[2], args[0])
	},
}

var tagDeleteLocalCmd = &cobra.Command{
	Use:   "local [name]",
	Short: "Alias for git tag -d <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("tag", "-d", args[0])
	},
}

var tagDeleteRemoteCmd = &cobra.Command{
	Use:   "remote [name]",
	Short: "Alias for git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("push", "origin", "--delete", args[0])
	},
}

var tagWipeCmd = &cobra.Command{
	Use:   "wipe [name]",
	Short: "Alias for git tag -d <name> && git push origin --delete <name>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("tag", "-d", args[0])
		runGitCommand("push", "origin", "--delete", args[0])
	},
}
