package handler

import "net/http"

/* Ping Handler: Ping Check 용도 */
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
