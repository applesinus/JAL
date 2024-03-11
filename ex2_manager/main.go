package main

import (
	ex2 "JAL/Ex2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Couldn't find a file path\n")
	} else {
		ex2.Ex2(os.Args[1])
	}

	fmt.Printf("\n\nPress any key to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
