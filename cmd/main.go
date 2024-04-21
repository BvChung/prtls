package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil{
		fmt.Print(err)
		os.Exit(1)
	}
}
