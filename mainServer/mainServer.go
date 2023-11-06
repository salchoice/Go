package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNum string = ":8080"

// Handler Functions 

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info Page")
}

func main(){
	log.Println("Starting our simple web server")

	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("Started on port",portNum)

	err := http.ListenAndServe(portNum, nil)
	if err != nil{
		log.Fatal(err)
	}
}

