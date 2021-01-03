package handlers

import (
	"log"
	"net/http"

	"github.com/lelouch99v/tasker/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
}

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
