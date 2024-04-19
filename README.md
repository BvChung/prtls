# File Directory Tree Generator

This command-line interface (CLI) tool generates a visual representation of the directory tree structure for the current directory. It's a handy utility for quickly visualizing the layout of your project's files and directories.

## Features

- Generates a tree structure of the current directory.
- Includes both files and directories in the output.
- Uses Unicode box-drawing characters for a clean, readable output.

## Usage

To use this tool, navigate to the directory you want to visualize in your terminal and run the command:

```bash
go run main.go
```

```bash
Compile to executable file with go build

Usage:
  ./prtls - [flags]

Flags:
  -a, --all    List all files (including hidden files), directories
  -h, --help   help for prtls
  -t, --tree   Tree view of the directory
```

## Example
```bash
./prtls -t . 
```

**Tree Structure**

![Tree](https://github.com/BvChung/prtls/blob/main/images/filetree.png)
