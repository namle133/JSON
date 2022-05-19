package main

import (
	"github.com/namle133/JSON.git/MainProgram/MainProgram/CREATE"
	"github.com/namle133/JSON.git/MainProgram/MainProgram/GETEVENT"
	"github.com/namle133/JSON.git/MainProgram/MainProgram/HOMEPAGE"
	"net/http"
)

func main() {
	http.HandleFunc("/", HOMEPAGE.HomeLink)
	http.HandleFunc("/event", CREATE.CreateEvent)
	http.HandleFunc("/events", GETEVENT.GetEvent)
	http.ListenAndServe(":8080", nil)

}
