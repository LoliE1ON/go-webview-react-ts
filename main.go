package main

import (
	"github/e1on/go-webview.git/application/infrastructure"
)

func main() {
	go infrastructure.CreateServer()

	infrastructure.CreateWebView()
}
