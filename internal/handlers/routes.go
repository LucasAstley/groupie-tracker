package handlers

import (
	"groupie-tracker/internal"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, artists *[]internal.Artist) {
	/*
		This function is used to render the index page
	*/
	renderTemplate(w, "index", artists)
}

func Artist(w http.ResponseWriter, r *http.Request, artists *[]internal.Artist) {
	/*
		This function is used to render the artist page
	*/
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		ErrorPage(w, http.StatusBadRequest, "Artist ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	var artist *internal.Artist
	for _, a := range *artists {
		if a.Id == id {
			artist = &a
			break
		}
	}

	if artist == nil {
		ErrorPage(w, http.StatusNotFound, "Artist not found")
		return
	}

	renderTemplate(w, "artist", artist)
}

func ErrorPage(w http.ResponseWriter, code int, message string) {
	/*
		This function is used to render the error page
	*/
	w.WriteHeader(code)

	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}

	renderTemplate(w, "error", data)
}
