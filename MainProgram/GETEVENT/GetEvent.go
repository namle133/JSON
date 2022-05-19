package GETEVENT

import (
	"encoding/json"
	"github.com/namle133/JSON.git/MainProgram/MainProgram/CREATE"
	"net/http"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CREATE.Events)
}
