package handlers

import (
	"encoding/json"
	"groupie-tracker/internal"
	"net/http"
	"strconv"
)

func GetConcerts(artistId int, artistsData *[]internal.Artist) []internal.Concert {
	/*
		This function returns the concerts of a given artist
	*/
	concertsArray := []internal.Concert{}
	for _, artist := range *artistsData {
		if artist.Id == artistId {
			concertsArray = artist.Concerts
			break
		}
	}
	return concertsArray
}

func GetConcertsDataHandler(artistsData *[]internal.Artist) http.HandlerFunc {
	/*
		This function handle the request to get the concerts of a given artist
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		artistIdStr := r.URL.Query().Get("artistId")
		if artistIdStr == "" {
			http.Error(w, "artistId is required", http.StatusBadRequest)
			return
		}

		artistId, err := strconv.Atoi(artistIdStr)
		if err != nil {
			http.Error(w, "invalid artistId", http.StatusBadRequest)
			return
		}

		concerts := GetConcerts(artistId, artistsData)

		json.NewEncoder(w).Encode(concerts)
	}
}
