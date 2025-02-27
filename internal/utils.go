package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfigUrl(requestedUrl string) string {
	/*
		This function reads the config.json file and returns the URL for the requested key
	*/
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer jsonFile.Close()

	var config map[string]string
	if err := json.NewDecoder(jsonFile).Decode(&config); err != nil {
		fmt.Println(err)
		return ""
	}

	return config[requestedUrl]
}
