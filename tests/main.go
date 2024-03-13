package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IT WORKS!!!

func assign(line_num int, line string) (string, bool) {
	fmt.Printf("[%v] Assigning: %v\n", line_num, line)
	var brackets_line_bldr, final_bldr strings.Builder
	in_brackets := 0

	// calculating in-brackets values
	for i, c := range line {
		switch c {
		case '(':
			in_brackets++
			if in_brackets != 1 {
				brackets_line_bldr.Grow(1)
				brackets_line_bldr.WriteRune(c)
			}
		case ')':
			in_brackets--
			if in_brackets == 0 {
				fmt.Printf("[%v:%v] Brackets line (%v): %v\n", line_num, i, in_brackets, brackets_line_bldr.String())
				value, err := assign(line_num, brackets_line_bldr.String())
				if err {
					return "", true
				}
				final_bldr.WriteString(fmt.Sprintf("%v", value))
				brackets_line_bldr.Reset()
			} else {
				brackets_line_bldr.Grow(1)
				brackets_line_bldr.WriteRune(c)
			}
		default:
			if in_brackets > 0 {
				brackets_line_bldr.Grow(1)
				brackets_line_bldr.WriteRune(c)
			} else {
				final_bldr.Grow(1)
				final_bldr.WriteRune(c)
			}
		}
	}
	just_line := final_bldr.String()
	if len(just_line) == len(line) {
		line = just_line
	} else {
		line = final_bldr.String()
	}
	// saving a line without brackets
	fmt.Printf("[%v] Excluded brackets: %v\n", line_num, line)
	return line, false
}

func main() {
	string_to_test := "(bg*25+(6*(myvar-10)))"

	line, err := assign(0, string_to_test)
	fmt.Printf("%v => %v\n", line, err)

	fmt.Printf("\n\nPress any key to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
