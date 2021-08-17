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
	mux.HandleFunc("/create", h.CreateUser)
	// http.HandleFunc("/update", UpdateUserById)
	// http.HandleFunc("/get", GetUserById)
	log.Println("created routers")
	return mux
}
