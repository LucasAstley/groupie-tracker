package main

import (
	"fmt"
	"groupie-tracker/internal"
	"groupie-tracker/internal/handlers"
)

func main() {
	/*
		Main function of the project
	*/
	apiUrl := internal.GetConfigUrl("api_url_artists")

	request, err := internal.PingApi(apiUrl)
	if err != nil && request != "200 OK" {
		fmt.Println("API is not available")
		return
	}

	apiData, err := internal.GetApiData(apiUrl)
	if err == nil {
		artistsData, err := internal.CreateArtistsStruct(apiData)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := range artistsData {
			artistsData[i].Concerts, err = internal.CreateConcertsStruct(&artistsData[i])

			if err != nil {
				fmt.Println(err)
				return
			}
		}

		handlers.StartServer(&artistsData)
	}
}
