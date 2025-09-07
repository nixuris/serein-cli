## Find Operations

The `find` command provides a convenient wrapper around standard `grep` and `find` utilities, allowing you to search for words within files, or locate files and directories by name. It also includes an optional `delete` mode for removing matched items.

### Flags and CLI Input

*   **Search for words in files:**
    `serein find word [delete] <path> <terms...>`
    - `[delete]`: Optional. If present, prompts to delete files where terms are found.
    - `<path>`: The directory path to start the search from.
    - `<terms...>`: One or more words to search for.

*   **Search for files by name:**
    `serein find file [delete] <path> <terms...>`
    - `[delete]`: Optional. If present, prompts to delete matched files.
    - `<path>`: The directory path to start the search from.
    - `<terms...>`: One or more file names (or parts of names) to search for. Wildcards are automatically applied.

*   **Search for directories by name:**
    `serein find dir [delete] <path> <terms...>`
    - `[delete]`: Optional. If present, prompts to delete matched directories.
    - `<path>`: The directory path to start the search from.
    - `<terms...>`: One or more directory names (or parts of names) to search for. Wildcards are automatically applied.

### Examples

*   **Search for the word "error" in the current directory:**
    ```bash
    serein find word . "error"
    ```

*   **Search for "TODO" and "FIXME" in the `src` directory and delete files containing them:**
    ```bash
    serein find word delete ./src "TODO" "FIXME"
    ```

*   **Find files named "config" (or containing "config") in the current directory:**
    ```bash
    serein find file . "config"
    ```

*   **Find and delete all files containing "temp" in their name within the `/tmp` directory:**
    ```bash
    serein find file delete /tmp "temp"
    ```

*   **Find directories named "build" in the current directory:**
    ```bash
    serein find dir . "build"
    ```

*   **Find and delete directories containing "old" in their name within the `/var/log` directory:**
    ```bash
    serein find dir delete /var/log "old"
    ```
