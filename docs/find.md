## Find Operations

The `find` command provides convenient wrappers for `grep` and `find` to locate words, files, or directories. Each command includes an optional `delete` mode.

### Usage

*   **Find words within files:**
    ```bash
    serein find word <path> <terms...>
    ```

*   **Find and delete files containing specific words:**
    ```bash
    serein find word delete <path> <terms...>
    ```

*   **Find files by name:**
    ```bash
    serein find file <path> <terms...>
    ```

*   **Find and delete files by name:**
    ```bash
    serein find file delete <path> <terms...>
    ```

*   **Find directories by name:**
    ```bash
    serein find dir <path> <terms...>
    ```

*   **Find and delete directories by name:**
    ```bash
    serein find dir delete <path> <terms...>
    ```

### Examples

*   **Search for all occurrences of the word "error" in the current directory:**
    ```bash
    serein find word . "error"
    ```

*   **Find all files in `src/` containing "TODO" or "FIXME" and be prompted to delete them:**
    ```bash
    serein find word delete ./src "TODO" "FIXME"
    ```

*   **Find all files with "config" in their name:**
    ```bash
    serein find file . "config"
    ```

*   **Find and be prompted to delete all directories named `__pycache__`:**
    ```bash
    serein find dir delete . "__pycache__"
    ```