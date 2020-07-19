package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthParam struct {
	Email    string `json:email`
	Password string `json:password`
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	var auth AuthParam
	err = json.Unmarshal(body, &auth)
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	responseJson, err := json.Marshal(auth)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	w.Write(responseJson)
}
