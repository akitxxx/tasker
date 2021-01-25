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

func SignUp(w http.ResponseWriter, r *http.Request) {
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
	email := jsonBody["email"]
	password := jsonBody["password"]

	// validate
	if email == "" || password == "" {
		renderError(w, err, http.StatusBadRequest)
		return
	}

	// register user
	user, err := models.RegistUser(email, password)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// create token
	token, err := CreateToken(user)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusInternalServerError)
	}

	renderResponse(w, token, http.StatusOK)
}

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

func FindUserById(w http.ResponseWriter, r *http.Request) {
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

	user, err := models.FindUserById(userId)
	if err != nil {
		log.Println(err)
		renderError(w, err, http.StatusBadRequest)
		return
	}

	renderResponse(w, user, http.StatusOK)
}
