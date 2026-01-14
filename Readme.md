# Zippy 

**Zippy** is a lightweight, intuitive CLI tool built in Go to solve a simple problem: **Zipping files on Windows sucks.** Native Windows tools are often slow, lack fine-grained control, and don't allow you to easily exclude clutter like `.git` or `node_modules`. Zippy gives you professional-grade compression and flexibility with simple commands.

## Example of zipping files on Windows on the command line
**From Google**

You can zip files on Windows CMD using the built-in
tar command (Windows 10/11+) with tar -a -c -f archive.zip folder_or_files for basic zipping, or use the more powerful 7z command from the third-party 7-Zip tool (recommended for better features) via 7z a -tzip archive.zip folder_or_files after adding it to your system's PATH. For PowerShell integration from CMD, use powershell Compress-Archive -Path ...

## Key Features

* **Real Compression:** Uses the `Deflate` algorithm to actually shrink your files, not just bundle them.
* **Selective Zipping:** Zip exactly the files you want by listing them, rather than being forced to zip a whole folder.
* **Smart Directory Walking:** Recursively zips folders while automatically ignoring "clutter" (like `.git`) and preventing the "infinite loop" bug when zipping into the current directory.
* **Integrity Verification:** Automatically verifies that the archive is structurally sound after creation.

---

## Installation

1. Clone the repository.
2. Open your terminal in the project folder.
3. Build the executable:
```bash
go build -o zippy.exe .

```



---

## Usage & Commands

### 1. Zipping Specific Files

Instead of zipping a whole folder, you can list exactly which files you want to include.

```bash
zippy -o project.zip main.go run.go zip.go README.md

```

* **`-o`**: Sets the output name (defaults to `output.zip` if omitted).
* **Files**: Any number of files listed after the flags will be compressed into the archive.

### 2. Zipping a Directory (Recursive)

Use the `-d` flag to zip an entire folder.

```bash
zippy -d -o backup.zip ./my-project

```

* **Exclude Clutter:** Zippy is hardcoded to skip `.git` and `node_modules` folders to keep your archives clean.
* **Self-Awareness:** If you create `backup.zip` inside the folder you are zipping, Zippy will detect it and skip it to prevent corruption.

### 3. Quick Zip (Defaults)

If you don't provide an output name, Zippy defaults to `output.zip`.

```bash
zippy -d .

```

*(Zips the current directory into `output.zip` while skipping hidden system folders.)*

---

## Technical Improvements

| Feature | Why we added it |
| --- | --- |
| **Deflate Algorithm** | Standard ZIP tools only "store" files by default. We added `zip.Deflate` so your archives actually save disk space. |
| **Path Normalization** | Windows uses `\`, but ZIPs require `/`. Zippy automatically converts paths so your files can be opened safely on Mac or Linux. |
| **Absolute Path Check** | Fixed the bug where the program would stall if it tried to zip its own output file. |
| **Flag Parsing** | Replaced rigid argument counting with the `flag` package, allowing for optional inputs and unlimited file lists. |

## Things to improve
- Update the help section for the new way to use zippy. 
- Add an unzip option
- Add a way to exclude specific folders that the use wants
- Create a way to auto install on windows so it doesn't require a build and added to the path. 
---

## Author:

Aarav Sibbal

## Licence

# Released under MIT License

Copyright (c) 2023 Aarav Sibbal.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---