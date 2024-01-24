package handler

import (
	"net/http"

	"github.com/ahr-i/triton-server-front-end/models"
)

/* Model List 반환 */
func (h *Handler) getModelListHandler(w http.ResponseWriter, r *http.Request) {
	modelMap := models.GetModelList()
	responseData := make(map[string]string)

	for key := range modelMap {
		responseData["name"] = key
	}

	rend.JSON(w, http.StatusOK, responseData)
}
