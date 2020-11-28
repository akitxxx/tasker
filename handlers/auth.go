package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/lelouch99v/tasker/models"
)

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

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
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = models.FindByEmailAndPassword(user.Email, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// no rows
			renderError(w, err, http.StatusUnauthorized)
			return
		}

		renderError(w, err, http.StatusInternalServerError)
		return
	}

	// create token
	token, err := createToken(user)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)

	}

	renderResponse(w, token, http.StatusOK)
}

func createToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "__init__",
	})

	key := "secret"
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}
