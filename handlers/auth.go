package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(res)
}
