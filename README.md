# Serein CLI

Serein is an opinionated command-line interface (CLI) tool designed to streamline various system management and utility tasks.

## Features

*   **Container Management:** Easily build, delete, list, and run Podman containers, including options for interactive shells, background execution, volume mounts, and device passthrough.
*   **Music Conversion:** Convert audio files (FLAC, Opus) to MP3 format with embedded cover art, and format M3U playlists for compatibility.
*   **Nix System Management:** Manage NixOS system and Home Manager configurations, including building, listing, and deleting generations, and updating flakes.
*   **Archive Operations:** Compress and extract files using 7z, with support for password protection.
*   **YouTube Audio Download:** Download audio from YouTube URLs using `yt-dlp` with embedded thumbnails and metadata.
*   **Git Helper:** A collection of aliases for common Git commands.

## Installation

Serein is built using Nix flakes, providing a reproducible and easy way to install the tool.

### Quick Try (Run without Installation)

If you want to quickly try `serein` without installing it permanently:

1.  **Ensure Nix is installed** on your system with flake support enabled.
2.  **Run Serein directly:**

    ```bash
    nix run github:nixuris/serein-cli -- [args]
    ```
    (Replace `[args]` with any `serein` command and its arguments, e.g., `nix run github:nixuris/serein-cli -- music convert mp3 /path/to/dir`)

### For NixOS/Home Manager Configurations

If you manage your system or user environment with NixOS or Home Manager flakes, you can add `serein-cli` as an input to your configuration:

1.  **Add `serein-cli` as an input in your `flake.nix`:**

    ```nix
    # In your flake.nix (e.g., /etc/nixos/flake.nix or ~/.config/home-manager/flake.nix)
    {
      description = "Your personal NixOS/Home Manager configuration";

      inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; # Or whatever nixpkgs they use
        home-manager.url = "github:nix-community/home-manager";
        home-manager.inputs.nixpkgs.follows = "nixpkgs";

        # Add serein-cli flake as an input
        serein-cli.url = "github:nixuris/serein-cli";
        serein-cli.inputs.nixpkgs.follows = "nixpkgs"; # Ensure consistent nixpkgs
      };

      outputs = { self, nixpkgs, home-manager, serein-cli, ... } @ inputs: {
        # ... your existing outputs
      };
    }
    ```

2.  **Install `serein` in your NixOS or Home Manager configuration:**

    **Option A: Install System-Wide (NixOS Configuration)**

    ```nix
    # In your configuration.nix (or a NixOS module)
    { config, pkgs, lib, ... }:

    {
      environment.systemPackages = with pkgs; [
        # Reference serein from the serein-cli flake input
        inputs.serein-cli.packages.${pkgs.system}.default
      ];

      # ... other system configurations
    }
    ```

    **Option B: Install via Home Manager (User-Specific)**

    ```nix
    # In your Home Manager configuration (e.g., ~/.config/home-manager/home.nix)
    { config, pkgs, ... }:

    {
      home.packages = [
        # Reference serein from the serein-cli flake input
        inputs.serein-cli.packages.${pkgs.system}.default
      ];

      # ... other Home Manager options
    }
    ```

### Binary Distribution (For Non-Nix Users)

For users not using Nix, `serein` can be downloaded as a single executable binary.

1.  **Download the latest release:**
    Visit the [GitHub Releases page](https://github.com/nixuris/serein-cli/releases) and download the appropriate binary for your operating system and architecture.

2.  **Make the binary executable:**
    ```bash
    chmod +x ./serein
    ```

3.  **Move the binary to your PATH (optional but recommended):**
    ```bash
    sudo mv ./serein /usr/local/bin/
    ```

**Important:** Regardless of how you install `serein`, you must ensure that all [Prerequisites](#prerequisites) are installed and available in your system's PATH for `serein`'s commands to function correctly.

## Usage

Serein provides a set of subcommands for different functionalities. For detailed usage, including all flags and examples, please refer to the [full documentation](docs/docs.md).

### General Commands

*   **Display help:**
    ```bash
    serein --help
    serein [command] --help
    ```

## Prerequisites

Serein currently leverages on several external tools being available in your system's PATH (whether having these tools or not is not mandatory):

*   **`podman`**: For container management commands.
*   **`ffmpeg`**: For audio conversion.
*   **`yt-dlp`**: For YouTube audio downloads.
*   **`7z` (p7zip)**: For archive operations.
*   **`git`**: For git commands.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the GPL-3.0 License - see the LICENSE file for details.
