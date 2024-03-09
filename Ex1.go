package main

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

func Ex1() {
	fmt.Printf("==================================================\n\n\t\t   INFO:\n\n")
	fmt.Printf("This program is a Memory manager, \n")
	fmt.Printf("If you wanna stop the program, say 'n'\n")
	fmt.Printf("\n==================================================\n\n")

	var willContinue string

	for willContinue != "y" && willContinue != "n" {
		var (
			file            *os.File
			visibilityLevel int
		)

		stack := make([]map[string]variable, 0)

		fmt.Printf("=====\n\n")
		fmt.Printf("Enter the name of new file (with extension like .txt): ")
		var fileName string
		fmt.Scan(&fileName)

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			visibilityLevel = 0
			stack = append(stack, make(map[string]variable))
			defer file.Close()

			fmt.Printf("\n\tNEW FILE HAS BEEN OPENED\n\n")

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "{") {
					visibilityLevel++
					stack = append(stack, make(map[string]variable))
				} else if strings.Contains(line, "}") {
					stack = stack[:len(stack)-1]

					//to_delete := make([]string, 0)

					/*for _, visLv := range stack {
						for _, some_var := range visLv {
							if some_var.VisibilityLevel == visibilityLevel {
								//to_delete = append(to_delete, some_var.Name)
								delete(visLv, some_var.Name)
							}
						}
					}*/

					/*for i := range to_delete {
						delete(stack, to_delete[i])
					}*/

					if visibilityLevel > 0 {
						visibilityLevel--
					}
				} else if strings.Contains(line, "=") {
					newVar := variableSaver(line, visibilityLevel)
					stack[visibilityLevel][newVar.Name] = newVar
				} else if strings.Contains(line, "ShowVar") {
					to_print := make([]string, 0)
					if len(stack) > 0 {
						for _, visLv := range stack {
							if len(visLv) > 0 {
								for _, some_var := range visLv {
									to_print = append(to_print, some_var.Name)
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

			fmt.Printf("\nIs there any additional files?\nPlease type only 'y' for yes and 'n' for no.\nAnswer: ")
			fmt.Scan(&willContinue)

			for willContinue != "y" && willContinue != "n" {
				fmt.Printf("\nIncorrect Input\n")
				fmt.Printf("Is there any additional files?\nPlease type only 'y' for yes and 'n' for no.\nAnswer: ")
				fmt.Scan(&willContinue)
			}

			fmt.Printf("\n==================================================\n\n")
		}
	}
}
