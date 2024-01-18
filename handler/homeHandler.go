package handler

import "net/http"

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "index.html", http.StatusTemporaryRedirect)
}
