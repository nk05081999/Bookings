package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}
func about(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(4, 4)
	fmt.Fprintf(w, fmt.Sprintf("This is about home page and the sum of 4+4is %d", sum))
	fmt.Fprintf(w, "This is the About page")

}
func AddValues(x, y int) int {
	return x + y
}
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f,", 100.0, 10.0, f))
}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil

}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/Divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
