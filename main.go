package main

// ~/go/bin/air

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var videos = []string{
	"https://www.youtube.com/embed/k59lX70pZuk?si=KShJlrcikNHxHUbN",
	"https://www.youtube.com/embed/6N6hmz5dd7o?si=h9QN7YRS894BnBjA",
	"https://www.youtube.com/embed/tkQt6GoE36o",
	"https://www.youtube.com/embed/9eHqGEshB2A",
	"https://www.youtube.com/embed/TBMEBSfnJbQ",
	"https://www.youtube.com/embed/TybreaCetEA",
	"https://www.youtube.com/embed/OcpRc_LU-e0",
	"https://www.youtube.com/embed/buUa9jiD9os",
	"https://www.youtube.com/embed/GSU93sFdRls",
}

type Website struct {
	ImagePath string
	Link      string
}

var websites = []Website{
	{ImagePath: "static/websites/site1.png", Link: "/"},
	{ImagePath: "static/websites/vic-records.png", Link: "https://vicrecords.club/"},
	{ImagePath: "static/websites/site3.png", Link: "https://www.wec.education/"},
	{ImagePath: "static/websites/site2.png", Link: "https://www.cinebrief.com/"},
}

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
	http.HandleFunc("/websites", websitesHandler)

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
	data := struct {
		Videos []string
	}{
		Videos: videos,
	}
	tmpl.Execute(w, data)
}

func websitesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/websites.html"))
	data := struct {
		Websites []Website
	}{
		Websites: websites,
	}
	tmpl.Execute(w, data)
}
