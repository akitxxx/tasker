package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, lane, http.StatusOK)
}

func DeleteLane(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.NotFound(w, r)
		return
	}

	// localhost:xxxx/xxxx/{id}
	idStr := strings.SplitN(r.URL.Path, "/", 3)[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// delete lane
	err = models.DeleteLane(id)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
	}

	renderResponse(w, nil, http.StatusOK)
}
