package handler

import (
	"net/http"

	"github.com/ahr-i/triton-server-front-end/urls"
	"github.com/gorilla/mux"
)

func (h *Handler) getUrlHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlName := vars["name"]

	urlMap := urls.GetUrlList()

	url, err := urlMap[urlName]
	if !err {
		rend.JSON(w, http.StatusOK, nil)
	}

	responseData := map[string]string{
		"url": url,
	}

	rend.JSON(w, http.StatusOK, responseData)
}
