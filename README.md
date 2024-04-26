# File Directory Tree Generator

This command-line interface (CLI) terminal user interface allows users to dynamically view files in directories and generate a visual representation of the directory tree structure for the current directory to an output text file. It's a handy utility for quickly visualizing the layout of your project's files and directories.

# Tech
- [Go](https://go.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## Demo
```bash
./prtls -p . -o file.txt
```
<div align="center">
  <video src="https://github.com/BvChung/prtls/assets/88690065/b8b31ff5-37b1-4517-a67f-b5301551de8e"></video>
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
├── prt.exe                  
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


