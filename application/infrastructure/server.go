package infrastructure

import (
	"github/e1on/go-webview.git/application/config"
	"net/http"
)

func CreateServer(application config.Application) {
	fs := http.FileServer(http.Dir(application.Server.RendererPath))
	http.Handle(application.Server.RendererUrl, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		fs.ServeHTTP(w, r)
	}))
	err := http.ListenAndServe(application.Server.GetServerAddress(), nil)
	if err != nil {
		return
	}
}
