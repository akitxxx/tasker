package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	var auth Auth
	err = json.Unmarshal(body, &auth)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	res, err := json.Marshal(auth)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(res)
}
