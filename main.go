package main

import (
	"fmt"
	"net/http"
	"text/template"
)
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.Execute(w, nil)
}
func home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "templates/index.html")
}

func routeFunc() {
	http.HandleFunc("/", home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
}

func main() {
	routeFunc()
	http.ListenAndServe(":8080", nil)
}
