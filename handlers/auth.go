package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lelouch99v/tasker/models"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}
	var auth Auth
	err = json.Unmarshal(body, &auth)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	user, err := models.FindByEmailAndPassword(auth.Email, auth.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// no rows
			renderError(w, err, http.StatusUnauthorized)
			return
		}

		renderError(w, err, http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:  "ID",
		Value: user.Email,
	}
	http.SetCookie(w, cookie)
	renderResponse(w, nil, http.StatusOK)
}
