package main

// ~/go/bin/air

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	// Serve static files with proper MIME type handling
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ext := filepath.Ext(r.URL.Path); ext == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		}
		fs.ServeHTTP(w, r)
	})))

	// Define handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/videos", videosHandler)

	// Start the server
	fmt.Println("Starting server on :8001")
	http.ListenAndServe(":8001", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>Hello, HTMX!</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	tmpl.Execute(w, nil)
}

func videosHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/videos.html"))
	tmpl.Execute(w, nil)
}
