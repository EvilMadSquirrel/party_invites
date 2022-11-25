package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for idx, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name+".html")
		if err != nil {
			panic(err)
		}
		templates[name] = t
		fmt.Println("Loaded template", idx, name)
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	err := templates["welcome"].Execute(writer, nil)
	if err != nil {
		panic(err)
	}
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	err := templates["list"].Execute(writer, responses)
	if err != nil {
		panic(err)
	}
}
