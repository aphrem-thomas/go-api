package api

import (
	"fmt"
	"net/http"

	config "github.com/aphrem-thomas/go-api/configuration"
)

var Repo *Repository

type Repository struct {
	Conf *config.Configuartion
}

func NewRepo(c *config.Configuartion) *Repository {
	return &Repository{
		Conf: c,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("config is", rep.Conf)
	fmt.Fprintf(w, "About page...")
}
