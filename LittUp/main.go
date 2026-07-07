package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Tell the user the server is starting
	fmt.Println("Server starting on http://localhost:8080")

	// Serve all files in the "static" folder
	// When someone visits /styles.css, it looks in static/styles.css
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the homepage
	// When someone visits http://localhost:8080, show index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the root path
		if r.URL.Path != "/" {
			// If not, return a 404 Not Found
			http.NotFound(w, r)
			return
		}
		// Read the index.html file
		http.ServeFile(w, r, "index.html")
	})

	// Start the server on port 8080
	// If there's an error, log it and stop
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(" Server failed to start:", err)
	}
}