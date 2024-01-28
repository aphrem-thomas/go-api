package api

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page...")
}
