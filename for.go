package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
