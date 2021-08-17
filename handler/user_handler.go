package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/models"
)

//controllers/handlers
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		log.Println(err)
	}
	h.Services.CreateUser(user)
	log.Println(user, "user create success")
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	var u models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(resBody, &u)
	if err != nil {
		log.Println(err)
	}
	user, status, err := h.Services.GetUserById(u.ID)
	if err != nil {
		log.Println(err)
	}
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, string(js), status)
		return
	}
	log.Println(user, status, "get user")
	w.Write(js)
}
func (h *Handler) UpdateUserByUUID(w http.ResponseWriter, r *http.Request) {
	var u models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(resBody, &u)
	if err != nil {
		log.Println(err)
	}
	status, err := h.Services.UpdateUserByUUID(&u)
	if err != nil {
		log.Println(err)
	}
	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, string(js), status)
		return
	}
	log.Println(status, "user updated")
	w.Write(js)
}
