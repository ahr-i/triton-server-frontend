package urls

import (
	"encoding/json"
	"io/ioutil"
	"runtime"

	"github.com/ahr-i/triton-server-front-end/src/errController"
)

type Url struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var urlMap map[string]string

func Init(filePath string) {
	urlMap = make(map[string]string)

	readUrls(filePath)
}

func readUrls(filePath string) {
	_, fp, _, _ := runtime.Caller(1)

	file, err := ioutil.ReadFile(filePath)
	errController.ErrorCheck(err, "JSON READ ERROR: NO FILE", fp)

	var urls []Url

	err = json.Unmarshal(file, &urls)
	errController.ErrorCheck(err, "JSON READ ERROR", fp)

	for _, url := range urls {
		urlMap[url.Name] = url.Url
	}
}

func GetUrlList() map[string]string {
	return urlMap
}
