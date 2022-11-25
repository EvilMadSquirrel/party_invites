package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

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

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		err := templates["form"].Execute(writer, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})
		if err != nil {
			panic(err)
		}
	} else if request.Method == http.MethodPost {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		responseData := Rsvp{
			Email:      request.Form["email"][0],
			Name:       request.Form["name"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name.")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email.")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number.")
		}
		if len(errors) > 0 {
			err := templates["form"].Execute(writer, formData{Rsvp: &responseData, Errors: errors})
			if err != nil {
				panic(err)
			}
		} else {
			responses = append(responses, &responseData)

			if responseData.WillAttend {
				err := templates["thanks"].Execute(writer, responseData.Name)
				if err != nil {
					panic(err)
				}
			} else {
				err := templates["sorry"].Execute(writer, responseData.Name)
				if err != nil {
					panic(err)
				}
			}
		}

	}
}
