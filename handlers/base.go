package handlers

import (
	"net/http"
	"log"
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
	mux.HandleFunc("/",  redirectToSite)
	return mux
}
