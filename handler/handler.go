package handler

import (
	"log"
	"net/http"

	service "github.com/devstackq/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", h.CreateUser)       //post:fields, fname, lname, email, age
	mux.HandleFunc("/get", h.GetUserById)         //post:id
	mux.HandleFunc("/update", h.UpdateUserByUUID) //put:field:  fname, lName,  email, age
	log.Println("created routers")
	return mux
}
