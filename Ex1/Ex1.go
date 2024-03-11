package ex1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type variable struct {
	Name            string
	Value           int
	VisibilityLevel int
}

func variableSaver(line string, visibilityLevel int) variable {
	var newVar variable

	var_index := strings.Index(line, "var")
	equals_index := strings.Index(line, "=")

	name := line[var_index:equals_index]
	newVar.Name = name

	val, err := strconv.Atoi(line[equals_index+1 : len(line)-1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	newVar.Value = val

	newVar.VisibilityLevel = visibilityLevel

	return newVar
}

func Ex1(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Couldn't find the file \"%v\": %v\n", filePath, err)
	} else {
		visibilityLevel := 0
		stack := make([]map[string]variable, 0)
		stack = append(stack, make(map[string]variable))
		defer file.Close()

		fmt.Printf("\nINTERPRETING FILE %v\n\n", filePath)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, "{") {
				// Open new level
				visibilityLevel++
				stack = append(stack, make(map[string]variable))

			} else if strings.Contains(line, "}") {
				// Close level and drop all the variables
				stack = stack[:len(stack)-1]
				if visibilityLevel > 0 {
					visibilityLevel--
				}

			} else if strings.Contains(line, "=") {
				// Save variable
				newVar := variableSaver(line, visibilityLevel)
				stack[visibilityLevel][newVar.Name] = newVar

			} else if strings.Contains(line, "ShowVar") {
				// Show variables
				to_print := make(map[string]int)
				if len(stack) > 0 {
					for _, visLv := range stack {
						if len(visLv) > 0 {
							for _, some_var := range visLv {
								to_print[some_var.Name] = some_var.Value
							}
						}
					}
				}

				if len(to_print) > 0 {
					fmt.Printf("%v\n", to_print)
				} else {
					fmt.Printf("No variables\n")
				}
			}
		}
	}
}
