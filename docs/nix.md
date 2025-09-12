## Nix System Management

### Flags and CLI Input

*   **Clean up Nix store:**
    `serein nix clean`

*   **Search nixpkgs:**
    `serein nix search [package name]`

*   **Build a Home Manager configuration:**
    `serein nix home build [path/to/flake]`

*   **List Home Manager generations:**
    `serein nix home gen`

*   **Delete Home Manager generations:**
    `serein nix home gen delete [number]`

    `serein nix home gen delete [number1]-[number2]` Delete range from home manager generation number1 to number2.

*   **Build a NixOS system:**
    `serein nix sys build [path/to/flake]`

*   **List system generations:**
    `serein nix sys gen`

*   **Delete system generations:**
    `serein nix sys gen delete [number]`
    
    `serein nix sys gen delete [number1]-[number2]` Delete range of system generations from number1 to number2.

*   **Update Nix flakes:**
    `serein nix update`

*   **Format Alejandra style:**
    `serein nix lint`

### Examples

*   **Clean up the Nix store:**
    ```bash
    serein nix clean
    ```

*   **Search nixpkgs:**
    ```bash
    serein nix search blender
    ```

*   **Build a Home Manager configuration from a flake:**
    ```bash
    serein nix home build /etc/nixos/#home-manager
    ```

*   **List all Home Manager generations:**
    ```bash
    serein nix home gen
    ```

*   **Delete Home Manager generations (e.g., generation 5 to 10):**
    ```bash
    serein nix home gen delete 5-10
    ```

*   **Build a NixOS system from a flake:**
    ```bash
    serein nix sys build /etc/nixos/#nixos-config
    ```

*   **Update all Nix flakes:**
    ```bash
    serein nix update
    ```
