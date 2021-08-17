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

func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/index.html")
}

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/create", h.CreateUser)
	mux.HandleFunc("/", h.IndexParse)
	// http.HandleFunc("/update", UpdateUserById)
	// http.HandleFunc("/get", GetUserById)
	log.Println("created routers")
	return mux
}
