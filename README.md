# Serein CLI

Serein is an opinionated command-line interface (CLI) tool designed to streamline various system management and utility tasks, particularly focusing on container management, music conversion, Nix system configurations, and archive operations.

## Features

*   **Container Management:** Easily build, delete, list, and run Podman containers, including options for interactive shells, background execution, volume mounts, and device passthrough.
*   **Music Conversion:** Convert audio files (FLAC, Opus) to MP3 format with embedded cover art, and format M3U playlists for compatibility.
*   **Nix System Management:** Manage NixOS system and Home Manager configurations, including building, listing, and deleting generations, and updating flakes.
*   **Archive Operations:** Compress and extract files using 7z, with support for password protection.
*   **YouTube Audio Download:** Download audio from YouTube URLs using `yt-dlp` with embedded thumbnails and metadata.

## Installation

Serein is built using Nix flakes, providing a reproducible and easy way to install the tool.

1.  **Ensure Nix is installed** on your system with flake support enabled.
2.  **Clone the repository:**

    ```bash
    git clone https://github.com/YOUR_USERNAME/serein-cli.git
    cd serein-cli
    ```

3.  **Build and install Serein:**

    ```bash
    nix profile install .#
    ```

    This command will build the `serein` executable and add it to your user's Nix profile, making it available in your PATH.

## Usage

Serein provides a set of subcommands for different functionalities. Here are some common usage examples:

### General Commands

*   **Display help:**
    ```bash
    serein --help
    serein [command] --help
    ```

### Container Management

*   **List all containers:**
    ```bash
    serein container list
    ```
*   **Build a container image:**
    ```bash
    serein container build [image-name] /path/to/Dockerfile
    ```
*   **Run a shell in a container:**
    ```bash
    serein container shell [container-name]
    serein container shell --temp [container-name] # Temporary container
    ```
*   **Run a container in the background (silent):**
    ```bash
    serein container silent [container-name]
    serein container silent --mount [container-name] # Mount current directory
    serein container silent --usb [container-name] # Passthrough USB devices
    serein container silent --ip [container-name] # Passthrough usbmuxd for iPhone
    ```
*   **Manage iOS devices with containers:**
    ```bash
    serein container ios --sidestore # Run sidestore container
    serein container ios --pair      # Run ipairing container
    ```

### Music Utilities

*   **Convert audio files to MP3:**
    ```bash
    serein music mp3 /path/to/music/directory
    ```
*   **Format M3U playlist:**
    ```bash
    serein music playlist /path/to/playlist.m3u
    ```
*   **Download YouTube audio:**
    ```bash
    serein music download "<youtube-url>"
    ```

### Nix System Management

*   **Clean up Nix store:**
    ```bash
    serein nix clean
    ```
*   **Build Home Manager configuration:**
    ```bash
    serein nix home build /path/to/your/flake#home-manager-config
    ```
*   **List Home Manager generations:**
    ```bash
    serein nix home gen
    ```
*   **Delete Home Manager generations:**
    ```bash
    serein nix home gen delete [generation-number]
    ```
*   **Build NixOS system:**
    ```bash
    serein nix sys build /path/to/your/flake#nixos-config
    ```
*   **List system generations:**
    ```bash
    serein nix sys gen
    ```
*   **Delete system generations:**
    ```bash
    serein nix sys gen delete [generation-number]
    ```
*   **Update Nix flakes:**
    ```bash
    serein nix update
    ```

### Archive Operations

*   **Zip files:**
    ```bash
    serein zip myarchive.zip file1.txt file2.jpg
    serein zip myarchive.7z myfolder/
    ```
*   **Zip files with password:**
    ```bash
    serein zip password myarchive.zip file.txt
    ```
*   **Unzip files:**
    ```bash
    serein unzip myarchive.zip
    ```
*   **Unzip files with password:**
    ```bash
    serein unzip password mypassword myarchive.zip
    ```

## Prerequisites

Serein relies on several external tools being available in your system's PATH:

*   **`podman`**: For container management commands.
*   **`ffmpeg`**: For audio conversion.
*   **`yt-dlp`**: For YouTube audio downloads.
*   **`7z` (p7zip)**: For archive operations.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
