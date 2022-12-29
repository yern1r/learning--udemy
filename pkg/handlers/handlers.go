package handlers

import (
	"learning--udemy/pkg/render"
	"net/http"
)

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	//we changed location of our function
	//then called these function from their packages
	render.RenderTemplate(w, "Home.page.tmpl") //calling function
	//pass responseWritter , name of template which i wanna parse
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "About.page.tmpl")
}
