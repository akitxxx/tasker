package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lelouch99v/tasker/models"
)

func CreateLane(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	// get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// json parse
	newLane := models.Lane{}
	err = json.Unmarshal(body, &newLane)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	lane, err := models.CreateLane(&newLane)
}
