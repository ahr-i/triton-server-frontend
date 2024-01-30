package models

import (
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/ahr-i/triton-server-frontend/src/errController"
)

/* model_list.json Struct */
type Model struct {
	ModelName    string `json:"model_name"`
	ModelVersion string `json:"model_version"`
}

var modelMap map[string]string

/* Init: modelMap Setting */
func Init(filePath string) {
	modelMap = make(map[string]string)

	readModels(filePath)
}

/* Read JSON File */
func readModels(filePath string) {
	_, fp, _, _ := runtime.Caller(1)

	// File Read
	file, err := ioutil.ReadFile(filePath)
	errController.ErrorCheck(err, "JSON READ ERROR: NO FILE", fp)

	var models []Model

	// Decode
	err = json.Unmarshal(file, &models)
	errController.ErrorCheck(err, "JSON READ ERROR", fp)

	// Save Values In The modelMap
	for _, model := range models {
		modelMap[model.ModelName] = model.ModelVersion
	}
}

/* Get modelMap */
func GetModelList() map[string]string {
	return modelMap
}
