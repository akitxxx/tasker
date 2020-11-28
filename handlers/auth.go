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

type Auth struct {
	Email    string `json"email"`
	Password string `json"password"`
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// JwtMiddleware check token
func JwtMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtMiddleware.Handler(handler).ServeHTTP(w, r)
	}
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	var auth Auth
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
		return
	}
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

	// create token
	token, err := createToken(user)
	if err != nil {
		renderError(w, err, http.StatusInternalServerError)
	}

	renderResponse(w, token, http.StatusOK)
}

func createToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
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
