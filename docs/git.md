## Git Commands

### Flags and CLI Input

*   **Sync a branch:**
    `serein git sync [branch]`

*   **Stage changes:**
    `serein git stage [path]`

*   **Unstage changes:**
    `serein git unstage [path]`

*   **Undo changes:**
    `serein git undo [path]`

*   **Show changes:**
    `serein git changes [path]`

*   **Show status:**
    `serein git status`

*   **Commit changes:**
    `serein git commit "[message]"`

*   **Push commit:**
    `serein git commit push [branch]`

*   **List commits:**
    `serein git commit list`

*   **Undo a commit:**
    `serein git commit undo [SHA]`

*   **Delete a commit:**
    `serein git commit delete [number]`

*   **Show commit changes:**
    `serein git commit changes [SHA]`

*   **Compare commits:**
    `serein git commit compare [SHA1] [SHA2]`

*   **List branches:**
    `serein git branch list`

*   **Create a branch:**
    `serein git branch create [name]`

*   **Switch branch:**
    `serein git branch switch [name]`

*   **Delete local branch:**
    `serein git branch delete local [name]`

*   **Delete remote branch:**
    `serein git branch delete remote [name]`

*   **Create a tag:**
    `serein git tag create [SHA] [name] "[message]"`

*   **Delete local tag:**
    `serein git tag delete local [name]`

*   **Delete remote tag:**
    `serein git tag delete remote [name]`

*   **Wipe tag (local and remote):**
    `serein git tag wipe [name]`

### Examples

*   **Sync the current branch with origin/main:**
    ```bash
    serein git sync main
    ```

*   **Stage a specific file:**
    ```bash
    serein git stage src/main.go
    ```

*   **Commit changes with a message:**
    ```bash
    serein git commit "feat: Add new feature"
    ```

*   **Push the current branch to origin:**
    ```bash
    serein git commit push main
    ```

*   **List all commits:**
    ```bash
    serein git commit list
    ```

*   **Switch to a new branch:**
    ```bash
    serein git branch switch develop
    ```

*   **Create a new tag:**
    ```bash
    serein git tag create abcdef1 v1.0.0 "Release version 1.0.0"
    ```