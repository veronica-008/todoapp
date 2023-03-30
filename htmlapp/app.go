package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}
type Pagedata struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := Pagedata{
		Title: "TO DO LIST",
		Todos: []Todo{
			{Item: "Set uo Go", Done: true},
			{Item: "Play with the cats", Done: true},
			{Item: "Learn Go", Done: false}},
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()

	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/todoapp", todo)

	log.Fatal(http.ListenAndServe(":9000", mux))
}
