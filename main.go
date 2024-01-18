package main

import (
	"net/http"

	"github.com/ahr-i/triton-server-front-end/handler"
	"github.com/ahr-i/triton-server-front-end/models"
	"github.com/ahr-i/triton-server-front-end/setting"
	"github.com/ahr-i/triton-server-front-end/src/corsController"
	"github.com/ahr-i/triton-server-front-end/urls"
	"github.com/urfave/negroni"
)

func Init() {
	models.Init(setting.ModelPath)
	urls.Init(setting.UrlPath)
}

func main() {
	Init()

	mux := handler.CreateHandler()
	handler := negroni.Classic()

	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	http.ListenAndServe(":"+setting.ServerPort, handler)
}
