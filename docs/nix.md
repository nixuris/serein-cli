## Nix System Management

### Flags and CLI Input

*   **Clean up Nix store:**
    `serein nix clean`

*   **Build a Home Manager configuration:**
    `serein nix home build [path/to/flake]`

*   **List Home Manager generations:**
    `serein nix home gen`

*   **Delete Home Manager generations:**
    `serein nix home gen delete [number]`

*   **Build a NixOS system:**
    `serein nix sys build [path/to/flake]`

*   **List system generations:**
    `serein nix sys gen`

*   **Delete system generations:**
    `serein nix sys gen delete [number]`

*   **Update Nix flakes:**
    `serein nix update`

*   **Format Alejandra style:**
    `serein nix lint`

### Examples

*   **Clean up the Nix store:**
    ```bash
    serein nix clean
    ```

*   **Build a Home Manager configuration from a flake:**
    ```bash
    serein nix home build /etc/nixos/#home-manager
    ```

*   **List all Home Manager generations:**
    ```bash
    serein nix home gen
    ```

*   **Delete a specific Home Manager generation (e.g., generation 5):**
    ```bash
    serein nix home gen delete 5
    ```

*   **Build a NixOS system from a flake:**
    ```bash
    serein nix sys build /etc/nixos/#nixos-config
    ```

*   **Update all Nix flakes:**
    ```bash
    serein nix update
    ```
