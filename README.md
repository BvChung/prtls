# File Directory Tree Generator

This terminal user interface application allows users to dynamically view files in directories and generate a visual representation of the directory tree structure for the current directory to an output text file. It's a handy utility for quickly visualizing the layout of your project's files and directories. Go to [releases](https://github.com/BvChung/prtls/releases) for the latest executable.

## Table of Contents
- [Technologies](#tech)
- [Usage](#usage)
- [Demo](#demo)
- [Development](#dev)
- [Issues](#issues)

## Tech <a name="tech"></a>
- [Go](https://go.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## Usage <a name="usage"></a>
```bash
Usage:
  ./prtls - [flags]

Flags:
  -h, --h, -help, --help        help command

  -o string
        -o, -o=, --o string     output file path (will create a new file as long as the path is valid)

  -p string
        -p, -p=, --p string     optional initial directory path, defaults to current directory (default ".")
```

## Demo <a name="demo"></a>
```bash

Build executable: go build -o ./prtls.exe ./cmd

./prtls.exe -p . -o file.txt
```
<div align="center">
  <video src="https://github.com/BvChung/prtls/assets/88690065/8acd1402-fc12-4f58-9091-dfb03df87062"></video>
</div>

## File Structure

```bash
.                            
├── README.md                
├── bin                      
│   └── prtls.exe            
├── cmd                      
│   └── main.go              
├── demo                     
│   ├── demo.gif             
│   ├── dirs.png             
│   ├── tree.png             
│   └── treegen.gif          
├── fstraversal              
│   ├── model.go             
│   ├── output.go            
│   ├── output_test.go       
│   ├── traversal.go         
│   └── traversal_test.go    
├── go.mod                   
├── go.sum                                   
├── prtls.exe                
└── testdata                 
    └── mockDirectory        
        └── b                
            └── c            
                ├── cc.py    
                └── d        
                    ├── a.txt
                    └── b.txt
```

To use this tool, navigate to the directory you want to visualize in your terminal and run the command:

## Development <a name="dev"></a>
```bash
go run ./cmd . -[flags]

Build to executable
go build -o {name} ./cmd
```

## Issues  <a name="issues"></a>
Create a new issue if there are any problems with this application or any suggestions you would like to make
