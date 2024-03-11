package main

import (
	ex1 "JAL/Ex1"
	ex2 "JAL/Ex2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	exercise := "none"
	for exercise != "exit" && exercise != "-1" {
		switch exercise {

		case "1":
			fmt.Printf("Enter the full path of the file: ")
			var filePath string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			filePath = scanner.Text()
			ex1.Ex1(filePath)

		case "2":
			fmt.Printf("Enter the full path of the file: ")
			var filePath string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			filePath = scanner.Text()
			ex2.Ex2(filePath)

		/*case "3":
			Ex3()
		case "4":
			Ex4()
		case "5":
			Ex5()
		case "6":
			Ex6()
		case "7":
			Ex7()
		case "8":
			Ex8()
		case "9":
			Ex9()
		case "10":
			Ex10()*/
		case "big":
			//ExBig.ExBig("kek.txt")
		default:
			fmt.Println("Invalid exercise number or command")
		}

		fmt.Println("\nType here an exercise number or 'big' for the big one ('exit' or '-1' to exit):")
		fmt.Scanln(&exercise)
	}
}
