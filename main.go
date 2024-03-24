package main

import (
	"net/http"
)

func main() {
	// Define the directory paths for serving static files
	htmlDir := "./html"
	cssDir := "./css"
	imagesDir := "./images"

	// Create file servers for each directory
	htmlFileServer := http.FileServer(http.Dir(htmlDir))
	cssFileServer := http.FileServer(http.Dir(cssDir))
	imagesFileServer := http.FileServer(http.Dir(imagesDir))

	// Handle requests for static files by serving from the respective directories
	http.Handle("/html/", http.StripPrefix("/html/", htmlFileServer))
	http.Handle("/css/", http.StripPrefix("/css/", cssFileServer))
	http.Handle("/images/", http.StripPrefix("/images/", imagesFileServer))

	// Handler for serving index.html on root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, htmlDir+"/index.html")
	})

	// Start the HTTP server on port 8080
	port := ":8080"
	println("Server listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
