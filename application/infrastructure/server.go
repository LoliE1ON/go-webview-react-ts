package infrastructure

import (
	"github/e1on/go-webview.git/application/config"
	"net/http"
)

func CreateServer(application *config.Application) {
	fs := http.FileServer(http.Dir(application.Server.RendererPath))
	http.Handle(application.Server.RendererUrl, fs)
	err := http.ListenAndServe(application.Server.GetServerAddress(), nil)
	if err != nil {
		return
	}
}
