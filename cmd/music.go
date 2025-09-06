package cmd

import (
	"github.com/spf13/cobra"
	"serein/internal/music"
)

var musicCmd = &cobra.Command{
	Use:   "music",
	Short: "Music related utilities",
	Long:  `Music related utilities.`,
}

func init() {
	rootCmd.AddCommand(musicCmd)
	musicCmd.AddCommand(music.YTMusicDownloadCmd)
	musicCmd.AddCommand(music.ConvertCmd)
}
