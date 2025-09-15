## Container Management

This module provides a set of convenient wrappers around `podman` for common container management tasks.

### Usage

#### Containers

*   **List all containers (running and stopped):**
    ```bash
    serein container list
    ```

*   **Delete a stopped container:**
    ```bash
    serein container delete [container-name]
    ```

*   **Export a container to a tarball:**
    ```bash
    serein container export <name>
    # Creates <name>.tar
    ```

*   **Import a container from a tarball:**
    ```bash
    serein container import <path-to-tar>
    # Creates a container with a name derived from the tar file
    ```

*   **Start an interactive shell in a container:**
    ```bash
    serein container shell [image-name]
    ```
    *   `--temp` or `-t`: Create a temporary container that is automatically removed on exit.
    *   `--mount` or `-m`: Mount the current directory into `/mnt` inside the container.
    *   `--usb` or `-u`: Pass through host USB devices to the container.
    *   `--ip`: Pass through `usbmuxd` for iOS device communication.

*   **Run a container in the background (detached mode):**
    ```bash
    serein container silent [image-name]
    ```
    *   Accepts the same `--mount`, `--usb`, and `--ip` flags as the `shell` command.

#### Container Images

*   **Build a container image from a Dockerfile:**
    ```bash
    serein container build [image-name] [path/to/dockerfile]
    ```

*   **List all images:**
    ```bash
    serein container images list
    ```

*   **Delete an image:**
    ```bash
    serein container images delete [image-id]
    ```

*   **Export an image to a tarball:**
    ```bash
    serein container images export <image-name>
    # Creates <name>.tar
    ```

*   **Import an image from a tarball:**
    ```bash
    serein container images import <path-to-tar>
    ```

#### iOS Helpers

*   **Manage iOS helper containers:**
    ```bash
    serein container ios
    ```
    *   `--sidestore` or `-s`: Run the SideStore helper container.
    *   `--pair` or `-p`: Run a container to pair your iOS device.

### Examples

*   **Build a new image named `my-app` from the current directory:**
    ```bash
    serein container build my-app .
    ```

*   **Export the `my-app` image to `my-app.tar`:**
    ```bash
    serein container images export my-app
    ```

*   **Import an image from a file:**
    ```bash
    serein container images import my-app.tar
    ```

*   **Run a temporary shell in an Ubuntu container with your current code mounted:**
    ```bash
    serein container shell --temp --mount ubuntu
    ```