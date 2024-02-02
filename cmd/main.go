package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/aphrem-thomas/go-api/api"
	config "github.com/aphrem-thomas/go-api/configuration"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templ, e := template.ParseFiles("./cmd/web/templates/home.page.tmpl")
	if e != nil {
		log.Println(e)
	}
	templ.Execute(w, nil)
}

func main() {

	conf := config.New("mongo", "chi")
	aboutConfig := api.NewRepo(conf)
	api.NewHandler(aboutConfig)

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", api.Repo.About)

	http.HandleFunc("/json", api.GetJson)

	serverErr := http.ListenAndServe(":8080", nil)
	if serverErr != nil {
		log.Println("error in starting the server", serverErr)
	}
}
