package internal

import "encoding/json"

type Artist struct {
	Id           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate json.Number `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	Location     string      `json:"locations"`
	Concerts     []Concert
}

type Concert struct {
	ArtistId int      `json:"id"`
	Place    string   `json:""`
	Dates    []string `json:""`
}
