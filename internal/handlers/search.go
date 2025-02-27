package handlers

import (
	"groupie-tracker/internal"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, artistsData []internal.Artist) {
	/*
		This function handle the search query
	*/
	searchQuery := r.URL.Query().Get("search")

	filteredArtists := searchedArtists(artistsData, searchQuery)

	renderTemplate(w, "index", filteredArtists)
}

func searchedArtists(artistsData []internal.Artist, searchQuery string) []internal.Artist {
	/*
		This function filter the artists based on the search query
	*/
	var filteredArtists []internal.Artist
	for _, artist := range artistsData {
		searchLower := strings.ToLower(searchQuery)

		if strings.Contains(strings.ToLower(artist.Name), searchLower) ||
			strings.Contains(strings.ToLower(string(artist.CreationDate)), searchLower) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), searchLower) ||
			memberContains(artist.Members, searchLower) ||
			concertContains(artist.Concerts, searchLower) {

			if !contains(filteredArtists, artist) {
				filteredArtists = append(filteredArtists, artist)
			}

		}
	}
	return filteredArtists
}

func contains(array []internal.Artist, artist internal.Artist) bool {
	/*
		This function check if an artist is already in the array
	*/
	for _, a := range array {
		if a.Id == artist.Id {
			return true
		}
	}
	return false
}

func memberContains(members []string, search string) bool {
	/*
		This function check if a member is in the members list
	*/
	for _, member := range members {
		if strings.Contains(strings.ToLower(member), search) {
			return true
		}
	}
	return false
}

func concertContains(concerts []internal.Concert, search string) bool {
	/*
		This function check if a concert is in the concerts list
	*/
	for _, concert := range concerts {
		if strings.Contains(strings.ToLower(concert.Place), search) {
			return true
		}
	}
	return false
}
