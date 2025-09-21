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

*   **Rename a container:**
    ```bash
    serein box rename <old_name> <new_name>
    ```

*   **Stop a container:**
    ```bash
    serein box stop <name>
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
    *   `--name` or `-n`: Assign a name to the imported container.

*   **Start an interactive shell in a container:**
    *   **Create a new container and run a shell:**
        ```bash
        serein box shell create [flags] <image>
        ```
        *   `--name` or `-n`: Assign a name to the container.
        *   `--temp` or `-t`: Create a temporary container that is automatically removed on exit.
        *   `--mount` or `-m`: Mount the current directory into `/mnt` inside the container.
        *   `--usb` or `-u`: Pass through host USB devices to the container.
        *   `--ip`: Pass through `usbmuxd` for iOS device communication.
        *   `--shell` or `-s`: Specify the shell to use (default: `sh`).

    *   **Resume a shell in an existing container:**
        ```bash
        serein box shell resume <container>
        ```
        *   `--shell` or `-s`: Specify the shell to use (default: `sh`).

*   **Run a container in the background (detached mode):**
    *   **Create a new container in the background:**
        ```bash
        serein box silent create [flags] <image>
        ```
        *   `--name` or `-n`: Assign a name to the container.
        *   `--mount` or `-m`: Mount the current directory into `/mnt` inside the container.
        *   `--usb` or `-u`: Pass through host USB devices to the container.
        *   `--ip`: Pass through `usbmuxd` for iOS device communication.

    *   **Resume an existing container in the background:**
        ```bash
        serein box silent resume <container>
        ```

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

*   **Rename an image:**
    ```bash
    serein box images rename <old_name> <new_name>
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

*   **Run a temporary shell in an Ubuntu container and name it `my-ubuntu`:**
    ```bash
    serein box shell create -t -m -n my-ubuntu ubuntu
    ```

*   **Force the use of docker for a command:**
    ```bash
    serein box --docker shell alpine
    ```
