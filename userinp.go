package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("Your are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
