package handlers

import (
	"log"
	"net/http"
	"os"
)

func redirectToSite(w http.ResponseWriter, r *http.Request) {
	// Loop over header names
	log.Println("-------")
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			log.Println(name, value)
		}
	}
	log.Println("-------")
	http.Redirect(w, r, os.Getenv("REDIRECT_URL"), http.StatusSeeOther)
}

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("site/")))
	//mux.HandleFunc("/", redirectToSite)
	return mux
}
