## Archive Operations

### Flags and CLI Input

*   **Zip files:**
    `serein archive zip [archive-name] [target-to-archive...]`

*   **Zip files with password:**
    `serein archive zip password [archive-name] [target-to-archive...]`

*   **Unzip files:**
    `serein archive unzip [target-to-unarchive]`

*   **Unzip files with password:**
    `serein archive unzip password [password] [target-to-unarchive]`

### Examples

*   **Create a zip archive from multiple files:**
    ```bash
    serein archive zip myarchive.zip file1.txt file2.jpg
    ```

*   **Create a 7z archive from a folder:**
    ```bash
    serein archive zip myarchive.7z myfolder/
    ```

*   **Create a password-protected zip archive:**
    ```bash
    serein archive zip password myarchive.zip sensitive_data.txt
    ```

*   **Extract an archive:**
    ```bash
    serein archive unzip myarchive.zip
    ```

*   **Extract a password-protected archive:**
    ```bash
    serein archive unzip password mypassword encrypted_archive.7z
    ```