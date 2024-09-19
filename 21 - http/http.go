package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", getHome)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/api/user", getJsonData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHome(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "user.html", "Marcelo")
}

func getJsonData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name": "Marcelo"}`))
}
