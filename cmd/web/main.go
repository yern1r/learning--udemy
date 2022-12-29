package main

import (
	"fmt"
	"learning--udemy/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function3
func main() {

	//we changed location of our function
	//then called these function from their packages
	http.HandleFunc("/", handlers.Home)       //gonna response to Home func , if there is "/"
	http.HandleFunc("/about", handlers.About) //gonna response to About func , if there is "/about"

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//if there is error - does mean nothing --> _
	_ = http.ListenAndServe(portNumber, nil)

}
