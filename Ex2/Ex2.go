package ex1

// test file on Linux: /home/dapar/Desktop/GitRepos/JAL/Ex2/Ex2_code.ex2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func (f function) execute() string {
	var result string
	return fmt.Sprintf("%v", result)
}

// Save_var returns true if there was an error
func save_var(line_num int, line string, pile map[string]variable) {
	opIndex := strings.Index(line, "(")
	name := line[0:opIndex]
	value_str := line[opIndex+4 : len(line)-1]

	if strings.Contains(line, "(i)=") {
		// if it's int
		value, err := strconv.Atoi(value_str)
		if err != nil {
			fmt.Printf("[line %v] Couldn't convert %v's value (%v) to int: %v\n",
				line_num, name, value_str, err)
			return
		} else {
			pile[name] = variable{"int", value}
			return
		}
	} else if strings.Contains(line, "(f)=") {
		// if it's float
		value, err := strconv.ParseFloat(value_str, 64)
		if err != nil {
			fmt.Printf("[line %v] Couldn't convert %v's value (%v) to float: %v\n",
				line_num, name, value_str, err)
			return
		} else {
			pile[name] = variable{"float", value}
			return
		}
	} else {
		// if it's unknown
		fmt.Printf("[line %v] Unknown variable type: %v\n", line_num, line)
		return
	}
}

// Show_vars prints all the variables from the pile
func show_vars(line_num int, pile map[string]variable) {
	fmt.Printf("\n\n\n[line %v] ACTUAL VARIABLES:\n\n", line_num)
	fmt.Printf("NAME\tTYPE\tVALUE\n-------------------------\n")
	for name, v := range pile {
		fmt.Printf("|%v\t|%v\t|%v\t|\n", name, v.JALtype, v.value)
	}
	fmt.Println("-------------------------")
}

// Saves the function
//
// [!] Doesn't check if it's valid
func function_saver(line_num int, line string, functions map[string]function) {
	parIndex := strings.Index(line, "(")
	bodyIndex := strings.Index(line, ":")
	name_str := line[0:parIndex]
	parameters_str := strings.Join(strings.Fields(line[parIndex+1:bodyIndex-1]), " ")
	body_str := line[bodyIndex+1:]

	functions[name_str] = function{strings.Split(parameters_str, ","), body_str}
	fmt.Printf("[line %v] Function \"%v\" created:\n%v\n", line_num, name_str, functions[name_str])
}

// Gets value from the pile or from the literal.
// Can return nil if there was an error or int, float values
//
// TESTED FOR LITERAL, INT, FLOAT
func get_value(line_num int, char_num int, expression string, pile map[string]variable) string {
	var value interface{}
	var err error
	// Try to parse it as an int
	value, err = strconv.Atoi(expression)
	if err != nil {
		// If it's not an int, try to parse it as a float
		value, err = strconv.ParseFloat(expression, 64)
		if err != nil {
			// If it's not a float, try to get it from the pile
			vrbl, exists := pile[expression]
			if !exists {
				// If it's not in the pile, print an error
				fmt.Printf("[line %v, char %v] Variable \"%v\" not found\n", line_num, char_num, expression)
				value = nil
			} else {
				// If it's in the pile, get it's value
				value = vrbl.value
			}
		}
	}
	return fmt.Sprintf("%v", value)
}

func parse_word(expression string) string {
	var result string

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
			} else { // TODO
				// result is a function
				result = funct.execute()
			}
		} else {
			// result is a variable
			result = fmt.Sprintf("%v", vrbl.value)
		}
	} else {
		// result is a literal
		result = fmt.Sprintf("%v", num)
	}

	return result
}

func parse_to_reverse_polish_notation(expression string, notation_type string) string {
	switch notation_type {

	case "postfix":
		return expression

	case "intfix":
		// getting rid of spaces (harder to parse, but more universal)
		expression = strings.ReplaceAll(expression, " ", "")

		// getting rid of unary minuses (only for ex2-ex3)
		for i := 0; i < len(expression); i++ {
			if expression[i] == '-' && (i == 0 ||
				expression[i-1] == '(' || expression[i-1] == ')' ||
				expression[i-1] == '+' || expression[i-1] == '-' ||
				expression[i-1] == '*' || expression[i-1] == '/') {

				expression = expression[0:i] + "~" + expression[i:len(expression)]
			}
		}

		// parsing
		stack := make([]string, 0)
		var word_builder strings.Builder
		stack = append(stack, "(")

		for _, c := range expression {
			switch c {

			case '-', '+', '*', '/', '(', ')':
				word_on_left := word_builder.String()
				word_builder.Reset()
				word := parse_word(word_on_left)
				if word != "ERROR" {
					stack = append(stack, word_on_left)
				} else {
					return "ERROR"
				}

			default:
				word_builder.Grow(1)
				word_builder.WriteRune(c)
			}
		}
	}

}

// Assigns the value to the variable. Can be recursive.
func assign(line_num int, line string, pile map[string]variable, functions map[string]function) (string, bool) {
	return "", false
}

func Ex2(filePath string) {
	//file, err := os.Open(filePath)
	file, err := os.Open("/home/dapar/Desktop/GitRepos/JAL/Ex2/Ex2_code.ex2")

	if err != nil {
		fmt.Printf("Couldn't find the file \"%v\": %v\n", filePath, err)
	} else {
		pile = make(map[string]variable)
		functions = make(map[string]function)
		defer file.Close()
		line_num := 0

		fmt.Printf("\nINTERPRETING FILE %v\n\n", filePath)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			line_num++

			if strings.Contains(line, "(i)=") || strings.Contains(line, "(f)=") {
				save_var(line_num, line, pile)
			} else if strings.Contains(line, "print") {
				show_vars(line_num, pile)
			} else if strings.Contains(line, ":") {
				function_saver(line_num, line, functions)
			} else if strings.Contains(line, "=") {
				var save_line strings.Builder
				save_line.WriteString(line[:strings.Index(line, "=")+1])
				assignation, err := assign(line_num, line[strings.Index(line, "=")+1:], pile, functions)
				if err {
					fmt.Printf("[line %v] Couldn't assign value\n", line_num)
				} else {
					save_line.WriteString(assignation)
					save_var(line_num, save_line.String(), pile)
				}
			}
		}
	}
}
