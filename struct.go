package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneMumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName:   "Neeraj",
		LastName:    "Kumar",
		PhoneMumber: "9149790991,8571046266",
		Age:         23,
	}
	log.Println("First name is ", user.FirstName)
	log.Println("Lirst name is ", user.LastName)
	log.Println("Phone number  is ", user.PhoneMumber)
	log.Println("Age  is ", user.Age)
	log.Println("Birthday date is ", user.BirthDate)

}
