package handlers

import (
	"fmt"
	"groupie-tracker/internal"
	"html/template"
	"net/http"
)

func StartServer(artistsData *[]internal.Artist) {
	/*
		This function is used to start the server and set the routes for the server
	*/
	fmt.Println("(http://localhost:8080) - Server started on port 8080")
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			ErrorPage(w, http.StatusNotFound, "Page not found")
			return
		}
		Index(w, r, artistsData)
	})

	http.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		Artist(w, r, artistsData)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		SearchHandler(w, r, *artistsData)
	})

	http.HandleFunc("/filters", func(w http.ResponseWriter, r *http.Request) {
		FiltersHandler(w, r, *artistsData)
	})

	http.HandleFunc("/reset-filters", ResetFiltersHandler())

	http.HandleFunc("/api/concerts", GetConcertsDataHandler(artistsData))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	/*
		This function is used to render the templates using the data sent
	*/
	t, err := template.ParseFiles("pkg/templates/" + tmpl + ".gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
