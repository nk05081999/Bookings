package main

import "fmt"

func main() {
	var myString string
	myString = "Neeraj"
	fmt.Println("My string is set to ", myString)
	changeUsingPointer(&myString)
	fmt.Println("After func calll my string is set to", myString)
}
func changeUsingPointer(s *string) {
	fmt.Println("The value of s is set to", s)
	newValue := "Red"
	*s = newValue
}
