package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePageHandler)
	/*
		http.Handle("/static/",
			http.StripPrefix("/static/",
				http.FileServer(http.Dir("static"))))
	*/
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
