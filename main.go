package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id    int
	Title string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Title: "HTMX and GO"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	addTodosHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		templ := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Title: title}
		data["Todos"] = append(data["Todos"], todo)

		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)

	http.HandleFunc("/add-todo", addTodosHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
