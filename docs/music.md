## Music Utilities

This module provides tools for downloading and converting audio files.

### Usage

*   **Download Audio from YouTube:**
    ```bash
    serein music download <youtube-url>
    ```
    *   This uses `yt-dlp` to download the audio track from a YouTube URL, embedding the thumbnail as cover art and including metadata.

*   **Convert Audio Files to MP3:**
    ```bash
    serein music convert mp3 <directory1> [directory2...]
    ```
    *   This command searches the specified directories for `.flac` and `.opus` files, converts them to high-quality MP3, embeds cover art, and deletes the source files upon successful conversion.

*   **Format a Playlist:**
    ```bash
    serein music convert playlist <path1.m3u> [path2.m3u...]
    ```
    *   This formats one or more `.m3u` playlist files to use Windows-style backslashes (`\`) and line endings, required for some hardware audio players.

### Examples

*   **Download audio from a YouTube video:**
    ```bash
    serein music download "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
    ```

*   **Convert audio files in multiple directories to MP3:**
    ```bash
    serein music convert mp3 ./flac-albums/ ./opus-downloads/
    ```

*   **Fix multiple playlists for a portable music player:**
    ```bash
    serein music convert playlist /media/my-sd-card/rock.m3u /media/my-sd-card/pop.m3u
    ```