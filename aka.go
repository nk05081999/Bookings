package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template
var templ *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("form.html"))
	templ = template.Must(template.ParseFiles("khalipage.html"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/hi", basic)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, req *http.Request) {
	tmpl.ExecuteTemplate(w, "form.html", nil)

}
func basic(w http.ResponseWriter, req *http.Request) {
	templ.ExecuteTemplate(w, "khalipage.html", nil)

}
