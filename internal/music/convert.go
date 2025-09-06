package music

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	ConvertCmd.AddCommand(convertMp3NewCmd)
	ConvertCmd.AddCommand(convertPlaylistCmd)
}

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Music related conversion utilities",
	Long:  `Music related conversion utilities.`,
}

var convertMp3NewCmd = &cobra.Command{
	Use:   "mp3 [directory]",
	Short: "Convert opus/flac to mp3",
	Long:  `Convert opus/flac to mp3 using ffmpeg.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		logPath := filepath.Join(dir, "conversion_errors.log")

		// create log file
		logFile, err := CreateFile(logPath)
		if err != nil {
			os.Exit(1)
		}
		defer CloseFile(logFile)

		// walk and convert
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if !info.IsDir() && (ext == ".flac" || ext == ".opus") {
				out := strings.TrimSuffix(path, ext) + ".mp3"

				if _, err := os.Stat(out); err == nil {
					fmt.Println("Skipping (exists):", out)
					return nil
				}

				fmt.Println("Converting + embedding cover:", path, "→", out)
				ffmpeg := exec.Command(
					"ffmpeg",
					"-nostdin",
					"-i", path,
					"-map", "0:a",
					"-map", "0:v?",
					"-c:a", "libmp3lame",
					"-q:a", "0",
					"-id3v2_version", "3",
					"-metadata:s:v", "title=Album cover",
					"-metadata:s:v", "comment=Cover (front)",
					out,
				)

				stderr, convErr := RunCommandWithStderr(ffmpeg)
				if convErr != nil {
					LogError(logFile, fmt.Sprintf(
						"Conversion error for %s: %v\nFFmpeg Output:\n%s\n",
						path, convErr, stderr,
					))
					fmt.Printf("Conversion error for %s: %v\nFFmpeg Output:\n%s\n",
						path, convErr, stderr,
					)
					return nil
				}

				fi, err := os.Stat(out)
				if err == nil && fi.Size() > 0 {
					fmt.Println("Converted:", out)
					fmt.Println("Deleting source:", path)
					_ = os.Remove(path)
				} else {
					LogError(logFile, fmt.Sprintf(
						"Conversion failed (zero-size or missing output): %s\n", path,
					))
					fmt.Printf("Conversion failed (zero-size or missing output): %s\n", path)
				}
			}
			return nil
		})

		fmt.Println("All done! Check conversion_errors.log for any errors.")
	},
}

var convertPlaylistCmd = &cobra.Command{
	Use:   "playlist [path/to/.m3u]",
	Short: "Format a playlist",
	Long:  `Format a playlist to be Winamp/Ruizu-safe.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		playlist := args[0]

		f, err := OpenFile(playlist)
		if err != nil {
			os.Exit(1)
		}
		defer CloseFile(f)

		var lines []string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		formatted := FormatPlaylistLines(lines)
		output := strings.Join(formatted, "\n")

		if err := os.WriteFile(playlist, []byte(output), 0o644); err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		fmt.Printf("✔️  '%s' is now Winamp/Ruizu-safe\n", playlist)
	},
}
