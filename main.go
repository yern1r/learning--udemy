package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function3
func main() {

	http.HandleFunc("/", Home)       //gonna response to Home func , if there is "/"
	http.HandleFunc("/about", About) //gonna response to About func , if there is "/about"

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//if there is error - does mean nothing --> _
	_ = http.ListenAndServe(portNumber, nil)

}
