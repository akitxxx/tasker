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

	lanes, err := models.SelectLaneList()
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
	}

	tasks, err := models.SelectTaskList()
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// marge lane and task
	for i := 0; i < len(lanes); i++ {
		for j := 0; j < len(tasks); j++ {
			if lanes[i].ID == tasks[j].LaneId {
				lanes[i].TaskList = append(lanes[i].TaskList, tasks[j])
			}
		}
	}

	renderResponse(w, lanes, http.StatusOK)
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
	newTask := models.Task{}
	err = json.Unmarshal(body, &newTask)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// validate
	if newTask.Title == "" {
		// title is required
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// create task
	task, err := models.CreateTask(&newTask)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, task, http.StatusOK)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
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
	if taskInput.ID == 0 {
		// ID is required
		// TODO render error message
		renderError(w, err, http.StatusBadRequest)
		return
	}

	targetTask, err := models.FindTaskById(taskInput.ID)
	if err != nil {
		renderError(w, err, http.StatusBadRequest)
	}

	if taskInput.Title != "" {
		targetTask.Title = taskInput.Title
	}
	if taskInput.Content != "" {
		targetTask.Content = taskInput.Content
	}

	// update task
	task, err := models.UpdateTask(targetTask)

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

func UpdateIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.NotFound(w, r)
		return
	}

	// get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderError(w, err, http.StatusBadRequest)
	}

	type TaskIndexData struct {
		TaskId     uint64 `json:"task_id"`
		SrcLaneId  uint64 `json:"src_lane_id"`
		SrcIndex   uint64 `json:"src_index"`
		DestLaneId uint64 `json:"dest_lane_id"`
		DestIndex  uint64 `json:"dest_index"`
	}

	indexDataInput := TaskIndexData{}
	if err = json.Unmarshal(body, &indexDataInput); err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	task, err := models.FindTaskById(indexDataInput.TaskId)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
	}

	// Update tasks in destination lane
	destTaskList, err := models.SelectTaskListAfterTargetIndex(indexDataInput.DestLaneId, indexDataInput.DestIndex)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}
	for _, t := range destTaskList {
		t.IndexNum++
		models.UpdateTask(&t)
	}

	// Update task index
	task.LaneId = indexDataInput.DestLaneId
	task.IndexNum = indexDataInput.DestIndex
	models.UpdateTask(task)

	// Update tasks in source lane
	srcTaskList, err := models.SelectTaskListAfterTargetIndex(indexDataInput.SrcLaneId, indexDataInput.SrcIndex)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}
	for _, t := range srcTaskList {
		t.IndexNum--
		models.UpdateTask(&t)
	}
}
