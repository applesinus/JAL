package main

import "fmt"

func main() {
	exercise := "none"
	for exercise != "exit" && exercise != "-1" {
		switch exercise {
		case "1":
			Ex1()
		/*case "2":
			Ex2()
		case "3":
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
			ExBig()
		default:
			fmt.Println("Invalid exercise number or command")
		}
		fmt.Println("Type here an exercise number or 'big' for the big one ('exit' or '-1' to exit):")
		fmt.Scan(&exercise)
	}
}
