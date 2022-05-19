package CREATE

import (
	"encoding/json"
	"net/http"
)

type event struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
}

type allEvents []event

var Events = allEvents{
	{
		ID:    "1",
		Title: "Golang",
	},
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	json.NewDecoder(r.Body).Decode(&newEvent)
	Events = append(Events, newEvent)
	json.NewEncoder(w).Encode(newEvent)
}
