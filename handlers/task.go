package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lelouch99v/tasker/models"
)

func GetTaskList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	tasks, err := models.SelectTaskList()
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	renderResponse(w, tasks, http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
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
	jsonBody := map[string]string{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}
	title := jsonBody["title"]
	content := jsonBody["content"]

	// validate
	if title == "" {
		// title is required
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// create task
	task, err := models.CreateTask(title, content)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, task, http.StatusOK)
}
