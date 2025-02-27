package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func PingApi(url string) (string, error) {
	/*
		This function sends a GET request to the given URL and returns the status of the response
	*/
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}

func GetApiData(url string) (string, error) {
	/*
		This function sends a GET request to the given URL and returns the body of the response
	*/
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CreateArtistsStruct(apiData string) ([]Artist, error) {
	/*
		This function creates a slice of Artist structs from the given JSON data
	*/
	var artists []Artist
	err := json.Unmarshal([]byte(apiData), &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func CreateConcertsStruct(artistData *Artist) ([]Concert, error) {
	/*
		This function creates a slice of Concert structs from the given Artist struct
	*/
	apiUrl := GetConfigUrl("api_url_relation") + "/" + strconv.Itoa(artistData.Id)
	apiData, err := GetApiData(apiUrl)
	if err != nil {
		return nil, err
	}

	var relationData struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	err = json.Unmarshal([]byte(apiData), &relationData)
	if err != nil {
		return nil, err
	}

	var concerts []Concert
	for place, dates := range relationData.DatesLocations {
		concert := Concert{
			ArtistId: relationData.Id,
			Place:    place,
			Dates:    dates,
		}
		concerts = append(concerts, concert)
	}

	return concerts, nil
}
