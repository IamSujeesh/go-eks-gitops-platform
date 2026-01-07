package main

import (
	"log"
	"net/http"
)

// aboutPage serves the About HTML page
func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

// contactPage serves the Contact HTML page
func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Serve About and Contact pages
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Default route redirects to /about
	http.HandleFunc("/", aboutPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}