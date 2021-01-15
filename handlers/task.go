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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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
	taskInput := models.Task{}
	err = json.Unmarshal(body, &taskInput)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// validate
	if taskInput.ID == 0 || taskInput.Title == "" {
		// ID and title are required
		// TODO render error message
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// update task
	task, err := models.UpdateTask(&taskInput)

	renderResponse(w, task, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.NotFound(w, r)
		return
	}

	// localhost:xxxx/xxxxx/id のようなパスによるid指定を想定
	idStr := strings.SplitN(r.URL.Path, "/", 3)[2]

	// userIDStrは文字列なのでuint64型に変換する。
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// delete task
	err = models.DeleteTask(id)

	renderResponse(w, nil, http.StatusOK)
}
