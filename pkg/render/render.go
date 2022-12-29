package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// render - give
// function allows take ResponseWritter and name of template which i wanna parse
// parse it and show it on browser

// RenderTemplate renders templactes using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) { //take 2 arguments = w ,  named template

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	//parse template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl) //path + name of template
	//this loads file from the root level of the application(. this dot means application)
	//inside, there is folder named templates
	//and we gonna take the argument of template that we wanna parse as the string
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var functions = template.FuncMap{}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	//holds all my templates and create when started our application
	myCache := map[string]*template.Template{} // taking all pages in template package
	// to look up to value very quickly
	//map index is string type , and we are using for index the name of template

	pages, err := filepath.Glob("./templates/*.page.tmpl") //go to the templates folder
	//and get all files starts with anything but end with .page.tmpl

	if err != nil {
		return myCache, err
	}

	for _, page := range pages { //every page which we find
		//we gonna get index , and page
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}
	return myCache, nil
}
