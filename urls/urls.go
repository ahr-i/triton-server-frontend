package urls

import (
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/ahr-i/triton-server-frontend/src/errController"
)

/* url_list.json Struct */
type Url struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var urlMap map[string]string

/* Init: urlMap */
func Init(filePath string) {
	urlMap = make(map[string]string)

	readUrls(filePath)
}

/* Read Json File */
func readUrls(filePath string) {
	_, fp, _, _ := runtime.Caller(1)

	// File Read
	file, err := ioutil.ReadFile(filePath)
	errController.ErrorCheck(err, "JSON READ ERROR: NO FILE", fp)

	var urls []Url

	// Decode
	err = json.Unmarshal(file, &urls)
	errController.ErrorCheck(err, "JSON READ ERROR", fp)

	// Save Values In The urlMap
	for _, url := range urls {
		urlMap[url.Name] = url.Url
	}
}

/* Get urlMap */
func GetUrlList() map[string]string {
	return urlMap
}
