package main

import (
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
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", about)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
