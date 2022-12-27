package main

import (
	"net/http"
)

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Home.page.html") //calling function
	//pass responseWritter , name of template which i wanna parse
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "About.page.html")
}
