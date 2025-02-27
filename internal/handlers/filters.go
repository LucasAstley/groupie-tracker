package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request, artistsData []internal.Artist) {
	/*
		This function is used to get the filters from the search form
	*/
	creationDateStart, _ := strconv.Atoi(r.URL.Query().Get("creation-date-start"))
	creationDateEnd, _ := strconv.Atoi(r.URL.Query().Get("creation-date-end"))
	firstAlbumDateStart, _ := strconv.Atoi(r.URL.Query().Get("first-album-date-start"))
	firstAlbumDateEnd, _ := strconv.Atoi(r.URL.Query().Get("first-album-date-end"))
	concertLocation := r.URL.Query().Get("locations")
	membersParams := r.URL.Query()["members"]
	var numberOfMembers []int
	for _, member := range membersParams {
		num, _ := strconv.Atoi(member)
		numberOfMembers = append(numberOfMembers, num)
	}

	filteredArtists := filterArtists(artistsData, creationDateStart, creationDateEnd, firstAlbumDateStart, firstAlbumDateEnd, numberOfMembers, concertLocation)

	renderTemplate(w, "index", filteredArtists)
}

func filterArtists(artistsData []internal.Artist, creationDateStart, creationDateEnd, firstAlbumDateStart, firstAlbumDateEnd int, numberOfMembers []int, concertLocation string) []internal.Artist {
	/*
		This function is used to filter the artists based on the filters provided
	*/
	var filteredArtists []internal.Artist

	for _, artist := range artistsData {
		isValid := true

		if creationDateStart != 0 || creationDateEnd != 0 {
			artistCreationDate, _ := strconv.Atoi(artist.CreationDate.String())
			if artistCreationDate < creationDateStart || artistCreationDate > creationDateEnd {
				isValid = false
			}
		}

		if isValid && (firstAlbumDateStart != 0 || firstAlbumDateEnd != 0) {
			artistFirstAlbumDate, _ := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
			if artistFirstAlbumDate < firstAlbumDateStart || artistFirstAlbumDate > firstAlbumDateEnd {
				isValid = false
			}
		}

		if isValid && len(numberOfMembers) > 0 {
			memberMatch := false
			for _, member := range numberOfMembers {
				if len(artist.Members) == member {
					memberMatch = true
					break
				}
			}
			if !memberMatch {
				isValid = false
			}
		}

		if isValid && concertLocation != "" {
			locationMatch := false
			for _, concert := range artist.Concerts {
				if strings.Contains(strings.ToLower(concert.Place), strings.ToLower(concertLocation)) {
					locationMatch = true
					break
				}
			}
			if !locationMatch {
				isValid = false
			}
		}

		if isValid {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists
}

func ResetFiltersHandler() func(w http.ResponseWriter, r *http.Request) {
	/*
		This function is used to reset the filters
	*/
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer jsonFile.Close()

	var config map[string]string
	if err := json.NewDecoder(jsonFile).Decode(&config); err != nil {
		fmt.Println(err)
		return nil
	}

	apiUrl := config["api_url_artists"]
	request, err := internal.PingApi(apiUrl)
	if err != nil && request != "200 OK" {
		fmt.Println("API is not available")
		return nil
	}

	apiData, err := internal.GetApiData(apiUrl)
	if err == nil {
		artistsData, err := internal.CreateArtistsStruct(apiData)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		return func(w http.ResponseWriter, r *http.Request) {
			Index(w, r, &artistsData)
		}
	}

	return nil
}
