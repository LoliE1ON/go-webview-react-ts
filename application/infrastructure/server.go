package infrastructure

import (
	"github/e1on/go-webview.git/application/config"
	"net/http"
)

func CreateServer() {
	fs := http.FileServer(http.Dir(config.Server.RendererPath))
	http.Handle(config.Server.RendererUrl, fs)
	err := http.ListenAndServe(config.Server.GetServerAddress(), nil)
	if err != nil {
		return
	}
}
