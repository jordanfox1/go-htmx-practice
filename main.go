package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	port = ":8080"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{
					Title:    "Die Hard",
					Director: "John McTiernan",
				},
				{
					Title:    "Mad Max: Fury Road",
					Director: "George Miller",
				},
				{
					Title:    "Inception",
					Director: "Christopher Nolan",
				},
				{
					Title:    "The Shawshank Redemption",
					Director: "Frank Darabont",
				},
				{
					Title:    "Forrest Gump",
					Director: "Robert Zemeckis",
				},
				{
					Title:    "The Godfather",
					Director: "Francis Ford Coppola",
				},
			},
		}
		templ.Execute(w, films)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film/", handler2)
	fmt.Printf("server trying to start on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
