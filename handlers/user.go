package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lelouch99v/tasker/models"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	users, err := models.SelectUserList()
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	renderResponse(w, users, http.StatusOK)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// localhost:xxxx/user/1 のようなパスによるuser id指定を想定
	userIdStr := strings.SplitN(r.URL.Path, "/", 3)[2]

	// userIDStrは文字列なのでuint64型に変換する。
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	user, err := models.FindById(userId)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	renderResponse(w, user, http.StatusOK)
}
