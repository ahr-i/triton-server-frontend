package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rend *render.Render = render.New()

type Handler struct {
	http.Handler
}

func CreateHandler() *Handler {
	mux := mux.NewRouter()
	handler := &Handler{
		Handler: mux,
	}

	mux.HandleFunc("/", handler.HomeHandler).Methods("GET")
	mux.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	mux.HandleFunc("/get/url/{name:[a-z-_]+}", handler.getUrlHandler).Methods("GET")
	mux.HandleFunc("/get/model-list", handler.getModelListHandler).Methods("GET")
	mux.HandleFunc("/model/{name:[a-z-_]+}/ready", handler.checkModelHandler).Methods("GET")
	mux.HandleFunc("/model/{name:[a-z-_]+}/infer", handler.inferHandler).Methods("POST")

	return handler
}
