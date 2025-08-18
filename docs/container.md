## Container Management

### Flags and CLI Input

*   **List all containers:**
    `serein container list`

*   **Build a container image:**
    `serein container build [name] [path/to/dockerfile]`

*   **Delete a container:**
    `serein container delete [name]`

*   **Manage container images:**
    `serein container images`

*   **Delete a container image:**
    `serein container images delete [id]`

*   **List all container images:**
    `serein container images list`

*   **Manage iOS devices with containers:**
    `serein container ios [--sidestore | -s] [--pair | -p]`

*   **Start a shell in a container:**
    `serein container shell [name] [--temp | -t]`

*   **Run a container in the background (silent):**
    `serein container silent [name] [--mount | -m] [--usb] [--ip]`

### Examples

*   **List all containers:**
    ```bash
    serein container list
    ```

*   **Build a container image:**
    ```bash
    serein container build my-image /path/to/Dockerfile
    ```

*   **Run a shell in a temporary container:**
    ```bash
    serein container shell --temp my-container
    ```

*   **Run a container in the background, mounting the current directory and passing through USB devices:**
    ```bash
    serein container silent --mount --usb my-container
    ```

*   **Run sidestore container for iOS device management:**
    ```bash
    serein container ios --sidestore
    ```