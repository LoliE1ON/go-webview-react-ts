package main

import (
	"github/e1on/go-webview.git/application/config"
	"github/e1on/go-webview.git/application/infrastructure"
	"github/e1on/go-webview.git/application/message"
)

func main() {
	var application = config.Application{
		Server: config.ServerConfig{
			Port:         8080,
			BaseUrl:      "http://localhost",
			RendererPath: "renderer/dist",
			RendererUrl:  "/",
			Debug:        true,
		},
		Window: config.WindowConfig{
			Title:   "Test application.go",
			Width:   800,
			Height:  600,
			WebView: infrastructure.CreateWebView(),
		},
		EventHandlerMap: map[string]config.EventHandler{
			"closeApplication": message.CloseApplication,
		},
	}

	go infrastructure.CreateServer(application)
	infrastructure.ConfigureWebView(application)
}
