package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aphrem-thomas/go-api/api"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", api.About)

	http.HandleFunc("/json", api.GetJson)

	serverErr := http.ListenAndServe(":8080", nil)
	if serverErr != nil {
		log.Println("error in starting the server", serverErr)
	}
}
