package main

import (
	ex1 "JAL/Ex1"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v", os.Args)
	if len(os.Args) <= 1 {
		fmt.Printf("Couldn't find a file path\n")
	} else {
		ex1.Ex1(os.Args[1])
	}

	fmt.Printf("\n\nPress any key to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
