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
    serein box shell [args...] [image-name]
    ```
    You can use flags (`--temp`, `--mount`, `--usb`, `--ip`) or positional arguments (`temp`, `mount`, `usb`, `ip`).

*   **Run a container in the background (detached mode):**
    ```bash
    serein box silent [args...] [image-name]
    ```
    You can use flags (`--mount`, `--usb`, `--ip`) or positional arguments (`mount`, `usb`, `ip`).

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
    You can use flags (`--sidestore`, `--pair`) or positional arguments (`sidestore`, `pair`).

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

*   **Run a temporary shell in an Ubuntu container with your current code mounted (using positional arguments):**
    ```bash
    serein box shell temp mount ubuntu
    ```

*   **Run a temporary shell in an Ubuntu container with your current code mounted (using flags):**
    ```bash
    serein box shell --temp --mount ubuntu
    ```

*   **Force the use of docker for a command:**
    ```bash
    serein box --docker shell alpine
    ```