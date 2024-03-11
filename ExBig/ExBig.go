package exbig

import (
	"bufio"
	"fmt"
	"os"
)

func ExBig(filePath string) {
	fmt.Printf("Opening file: %v\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't find the file \"%v\": %v\n", filePath, err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//line := scanner.Text()

	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Couldn't find a file path\n")
		return
	}
	ExBig(os.Args[1])
}
