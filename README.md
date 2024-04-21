# File Directory Tree Generator

This command-line interface (CLI) tool generates a visual representation of the directory tree structure for the current directory. It's a handy utility for quickly visualizing the layout of your project's files and directories.

## Features

- Generates a tree structure of the current directory.
- Includes both files and directories in the output.

## File Structure

```bash
.
├── README.md
├── assets
│   └── extension.go
├── bin
│   └── prtls.exe
├── cmd
│   ├── main.go
│   └── run.go
├── demo
│   ├── filetree.png
│   └── treegen.gif
├── fstraversal
│   ├── traversal.go
│   ├── traversal_test.go
│   └── types.go
├── go.mod
├── go.sum
└── testdata
    └── mockDirectory
        └── b
            └── c
                ├── cc.py
                └── d
                    ├── a.txt
                    └── b.txt
```

## Usage

To use this tool, navigate to the directory you want to visualize in your terminal and run the command:

```bash
go run main.go
```

# Compile to executable file
```bash
go build -o ./{path}/{executable name} ./cmd

Usage:
  ./prtls - [flags]

Flags:
  No flags     List all files in current directory
  -a, --all    Include hidden files (files prefixed with .)
  -h, --help   Help command
  -t, --tree   Tree view of the directory
```

## Example
```bash
./prtls -t . 
```
![Tree](https://github.com/BvChung/prtls/blob/main/demo/filetree.png)

```bash
./prtls -t . -a
```
![Complex Tree](https://github.com/BvChung/prtls/blob/main/demo/treegen.gif)


