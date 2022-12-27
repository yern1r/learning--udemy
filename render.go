package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// render - give
// function allows take ResponseWritter and name of template which i wanna parse
// parse it and show it on browser
func renderTemplate(w http.ResponseWriter, html string) { //take 2 arguments = w ,  named template
	//parse template
	parsedTemplate, _ := template.ParseFiles("./templates/" + html) //path + name of template
	//this loads file from the root level of the application(. this dot means application)
	//inside, there is folder named templates
	//and we gonna take the argument of template that we wanna parse as the string
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

//go module is management of packages
