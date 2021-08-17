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

	if r.Method == "POST" {
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

		log.Println(user, 1)

	}
}
