package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var convertPlaylistCmd = &cobra.Command{
	Use:   "playlist [path/to/.m3u]",
	Short: "Format a playlist",
	Long:  `Format a playlist to be Winamp/Ruizu-safe.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		playlistPath := args[0]

		file, err := os.Open(playlistPath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.ReplaceAll(line, "/", "\\")
			line += "\r"
			lines = append(lines, line)
		}

		output := strings.Join(lines, "\n")
		err = os.WriteFile(playlistPath, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		fmt.Println("✔️  '" + playlistPath + "' is now Winamp/Ruizu-safe")
	},
}

func init() {
	convertCmd.AddCommand(convertPlaylistCmd)
}
