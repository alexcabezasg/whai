package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
    args := os.Args[1:]
    // Command interpreter

    fmt.Println("Executed subcommands: ", args)
}