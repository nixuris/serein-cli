package box

import (
	"strings"

	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var tempShell bool
var shellMount, shellUsb, shellIp bool
var shellName, shell string

var silentMount, silentUsb, silentIp bool
var silentName string

// Resume command variables (separate from create to avoid conflicts)
var resumeShell string

var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Run a shell in a container",
	Long:  `Run a shell in a container.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var SilentCmd = &cobra.Command{
	Use:   "silent",
	Short: "Run a container in the background",
	Long:  `Run a container in the background.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var ShellCreateCmd = shared.NewCommand(
	"create [flags] <image>",
	"Create a shell in a new container. See flags for options.",
	cobra.MinimumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		imageName := args[len(args)-1]
		for _, arg := range args[:len(args)-1] {
			switch strings.ToLower(arg) {
			case "temp":
				tempShell = true
			case "mount":
				shellMount = true
			case "usb":
				shellUsb = true
			case "ip":
				shellIp = true
			}
		}
		RunContainerCommand(BuildShellCreateArgs(imageName, tempShell, shellMount, shellUsb, shellIp, shellName, shell), true)
	},
)

var ShellResumeCmd = shared.NewCommand(
	"resume <container>",
	"Resume a shell in an existing container.",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		RunContainerCommand(BuildShellResumeArgs(containerName, resumeShell), true)
	},
)

var SilentCreateCmd = shared.NewCommand(
	"create [flags] <image>",
	"Create a container in the background. See flags for options.",
	cobra.MinimumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		imageName := args[len(args)-1]
		for _, arg := range args[:len(args)-1] {
			switch strings.ToLower(arg) {
			case "mount":
				silentMount = true
			case "usb":
				silentUsb = true
			case "ip":
				silentIp = true
			}
		}
		RunContainerCommand(BuildDetachedCreateArgs(imageName, silentMount, silentUsb, silentIp, silentName), false)
	},
)

var SilentResumeCmd = shared.NewCommand(
	"resume <container>",
	"Resume an existing container in the background.",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		RunContainerCommand(BuildDetachedResumeArgs(containerName), false)
	},
)

func init() {
	ShellCmd.AddCommand(ShellCreateCmd)
	ShellCmd.AddCommand(ShellResumeCmd)
	SilentCmd.AddCommand(SilentCreateCmd)
	SilentCmd.AddCommand(SilentResumeCmd)

	// Shell Create flags
	ShellCreateCmd.Flags().BoolVarP(&tempShell, "temp", "t", false, "Use a temporary container")
	ShellCreateCmd.Flags().BoolVarP(&shellMount, "mount", "m", false, "Mount current directory to /mnt")
	ShellCreateCmd.Flags().BoolVarP(&shellUsb, "usb", "u", false, "Passthrough USB devices")
	ShellCreateCmd.Flags().BoolVar(&shellIp, "ip", false, "Passthrough usbmuxd for iPhone")
	ShellCreateCmd.Flags().StringVarP(&shellName, "name", "n", "", "Assign a name to the container")
	ShellCreateCmd.Flags().StringVarP(&shell, "shell", "s", "sh", "Specify the shell to use")

	// Shell Resume flags (only shell flag)
	ShellResumeCmd.Flags().StringVarP(&resumeShell, "shell", "s", "sh", "Specify the shell to use")

	// Silent Create flags
	SilentCreateCmd.Flags().BoolVarP(&silentMount, "mount", "m", false, "Mount current directory to /mnt")
	SilentCreateCmd.Flags().BoolVarP(&silentUsb, "usb", "u", false, "Passthrough USB devices")
	SilentCreateCmd.Flags().BoolVar(&silentIp, "ip", false, "Passthrough usbmuxd for iPhone")
	SilentCreateCmd.Flags().StringVarP(&silentName, "name", "n", "", "Assign a name to the container")

	// Silent Resume has no flags (containers resume with original settings)
}
