package main

import "fmt"

func main() {
	fmt.Println("Enter your Number:-")
	var a int
	fmt.Scanln(&a)
	switch a {
	case 10:
		fmt.Println("The number is 10")
	case 20:
		fmt.Println("The number is 20")
	case 30:
		fmt.Println("The number is 30")
	default:
		fmt.Println("It is not 10 , 20,30")
	}
}
