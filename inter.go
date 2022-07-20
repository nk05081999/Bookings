package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}
type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German shefered",
	}
	PrintInfo(&dog)
	gorilla := Gorilla{
		Name:          "Jock",
		color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}
func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "UGh"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
