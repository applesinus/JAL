package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Global pile of variables
var pile map[string]variable

// Global pile of functions
var functions map[string]function

// JALtype may be only:
//
// "int" for integers
//
// "float" for floats
type variable struct {
	JALtype string
	value   interface{}
}

// Body executes in the .execute() method
type function struct {
	parameters []string
	body       string
}

// Executes the function
//
// Returns the value of the function. Nil if there was an error
func (f function) execute() interface{} {
	return nil
}

func parse_word(expression string) string {
	// Try to parse literal
	num, err := strconv.ParseFloat(expression, 64)

	// If it's not a literal, try to get a variable from the pile
	if err != nil {
		vrbl, exists := pile[expression]

		// If it's not a variable, try to get a function
		if !exists {
			funct, exists := functions[expression]

			// If it's not a function, return an error
			if !exists {
				return "ERROR"
			} else {
				// TODO
				num = funct.execute().(float64)
			}
		} else {
			num = vrbl.value.(float64)
		}
	}
	return fmt.Sprintf("%v", num)
}

func main() {
	pile = make(map[string]variable)
	pile["a"] = variable{JALtype: "int", value: 10}
	string_to_test := "a"

	line := parse_word(string_to_test)
	fmt.Printf("%v => %v\n", string_to_test, line)

	fmt.Printf("\n\nPress any key to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
