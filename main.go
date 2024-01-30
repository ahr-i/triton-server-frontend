package main

import (
	"net/http"

	"github.com/ahr-i/triton-server-frontend/handler"
	"github.com/ahr-i/triton-server-frontend/models"
	"github.com/ahr-i/triton-server-frontend/setting"
	"github.com/ahr-i/triton-server-frontend/src/corsController"
	"github.com/ahr-i/triton-server-frontend/urls"
	"github.com/urfave/negroni"
)

/* Server Setting */
func Init() {
	models.Init(setting.ModelPath) // Model List Init
	urls.Init(setting.UrlPath)     // URL List Init
}

/* Main */
func main() {
	Init()

	mux := handler.CreateHandler()
	handler := negroni.Classic()

	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	http.ListenAndServe(":"+setting.ServerPort, handler)
}
