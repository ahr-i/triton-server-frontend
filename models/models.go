package models

import (
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/ahr-i/triton-server-front-end/src/errController"
)

type Model struct {
	ModelName    string `json:"model_name"`
	ModelVersion string `json:"model_version"`
}

var modelMap map[string]string

func Init(filePath string) {
	modelMap = make(map[string]string)

	readModels(filePath)
}

func readModels(filePath string) {
	_, fp, _, _ := runtime.Caller(1)

	file, err := ioutil.ReadFile(filePath)
	errController.ErrorCheck(err, "JSON READ ERROR: NO FILE", fp)

	var models []Model

	err = json.Unmarshal(file, &models)
	errController.ErrorCheck(err, "JSON READ ERROR", fp)

	for _, model := range models {
		modelMap[model.ModelName] = model.ModelVersion
	}
}

func GetModelList() map[string]string {
	return modelMap
}
