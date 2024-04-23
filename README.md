# File Directory Tree Generator

This command-line interface (CLI) terminal user interface allows users to dynamically view files in directories and generate a visual representation of the directory tree structure for the current directory to an output text file. It's a handy utility for quickly visualizing the layout of your project's files and directories.

# Tech
- [Go](https://go.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## Example
```bash
./prtls -p . -o file.txt
```
### Directory UI
![dirs](https://github.com/BvChung/prtls/blob/main/demo/dirs.png)

### Tree Viewer UI
![tree](https://github.com/BvChung/prtls/blob/main/demo/tree.png)

### Directory Navigation
![Navigation](https://github.com/BvChung/prtls/blob/main/demo/demo.gif)

### Tree Creation
![Tree Gif](https://github.com/BvChung/prtls/blob/main/demo/treegen.gif)

## File Structure

```bash
.                            
├── README.md                
├── bin                      
│   └── prtls.exe            
├── cmd                      
│   └── main.go              
├── demo                     
│   ├── filetree.png         
│   └── treegen.gif          
├── file.txt                 
├── fstraversal              
│   ├── model.go             
│   ├── output.go            
│   ├── traversal.go         
│   └── traversal_test.go    
├── go.mod                   
├── go.sum                   
├── prt                      
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

## Development
```bash
go run ./cmd . -[flags]
```

## Compiling to executable file
```bash
go build -o ./{path}/{executable name} ./cmd

Usage:
  ./prtls - [flags]

Flags:
  -p           Initial directory path
  -o           Output file path
```


