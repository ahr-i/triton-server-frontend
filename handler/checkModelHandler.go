package handler

import (
	"io"
	"net/http"
	"runtime"

	"github.com/ahr-i/triton-server-front-end/models"
	"github.com/ahr-i/triton-server-front-end/setting"
	"github.com/ahr-i/triton-server-front-end/src/errController"
	"github.com/gorilla/mux"
)

func (h *Handler) checkModelHandler(w http.ResponseWriter, r *http.Request) {
	_, fp, _, _ := runtime.Caller(1)

	vars := mux.Vars(r)
	model := vars["name"]

	modelMap := models.GetModelList()
	version, err := modelMap[model]
	if !err {
		rend.JSON(w, http.StatusOK, nil)
	}

	url := "http://" + setting.TritonUrl + "/v2/models/" + model + "/versions/" + version + "/ready"
	response, err_ := http.Get(url)
	errController.ErrorCheck(err_, "MODEL CHECK ERROR", fp)
	defer response.Body.Close()

	io.Copy(w, response.Body)
}
