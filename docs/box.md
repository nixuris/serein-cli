## Box Management

This module provides a set of convenient wrappers around `podman` or `docker` for common container management tasks.

### Usage

#### Containers

*   **List all containers (running and stopped):**
    ```bash
    serein box list
    ```

*   **Delete a stopped container:**
    ```bash
    serein box delete [container-name]
    ```

*   **Export a container to a tarball:**
    ```bash
    serein box export <name>
    # Creates <name>.tar
    ```

*   **Import a container from a tarball:**
    ```bash
    serein box import <path-to-tar>
    # Creates a container with a name derived from the tar file
    ```

*   **Start an interactive shell in a container:**
    ```bash
    serein box shell [image-name]
    ```
    *   `--temp` or `-t`: Create a temporary container that is automatically removed on exit.
    *   `--mount` or `-m`: Mount the current directory into `/mnt` inside the container.
    *   `--usb` or `-u`: Pass through host USB devices to the container.
    *   `--ip`: Pass through `usbmuxd` for iOS device communication.

*   **Run a container in the background (detached mode):**
    ```bash
    serein box silent [image-name]
    ```
    *   Accepts the same `--mount`, `--usb`, and `--ip` flags as the `shell` command.

#### Container Images

*   **Build a container image from a Dockerfile:**
    ```bash
    serein box build [image-name] [path/to/dockerfile]
    ```

*   **List all images:**
    ```bash
    serein box images list
    ```

*   **Delete an image:**
    ```bash
    serein box images delete [image-id]
    ```

*   **Export an image to a tarball:**
    ```bash
    serein box images export <image-name>
    # Creates <name>.tar
    ```

*   **Import an image from a tarball:**
    ```bash
    serein box images import <path-to-tar>
    ```

#### iOS Helpers

*   **Manage iOS helper containers:**
    ```bash
    serein box ios [args...]
    ```
    You can use flags (`--sidestore`, `--pair`).

### Examples

*   **Build a new image named `my-app` from the current directory:**
    ```bash
    serein box build my-app .
    ```

*   **Export the `my-app` image to `my-app.tar`:**
    ```bash
    serein box images export my-app
    ```

*   **Import an image from a file:**
    ```bash
    serein box images import my-app.tar
    ```

*   **Run a temporary shell in an Ubuntu container:**
    ```bash
    serein box shell -t -m ubuntu
    ```

*   **Force the use of docker for a command:**
    ```bash
    serein box --docker shell alpine
    ```
