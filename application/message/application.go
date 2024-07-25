package message

import "github/e1on/go-webview.git/application/config"

func CloseApplication(application config.Application, data interface{}) interface{} {
	application.Window.WebView.Terminate()

	return nil
}
