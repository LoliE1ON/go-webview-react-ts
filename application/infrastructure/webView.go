package infrastructure

import (
	"encoding/json"
	"github.com/jchv/go-webview2"
	"github.com/lxn/win"
	"github/e1on/go-webview.git/application/config"
)

func CreateWebView() webview2.WebView {
	var webView = webview2.New(true)
	defer webView.Destroy()

	return webView
}

func ConfigureWebView(application *config.Application, webView webview2.WebView) {
	webView.SetTitle(application.Window.Title)
	webView.SetSize(application.Window.Width, application.Window.Height, webview2.HintNone)
	webView.Navigate(application.Server.GetRendererBaseUrl())
	webView.Bind("sendMessage", func(msg string) string {
		var message = jsonToMessage(msg)

		handler, exists := application.EventHandlerMap[message.Event]
		if !exists {
			return `{"error": "Unknown event"}`
		}

		response := handler(application, message.Data)
		return messageToJson(response)
	})

	hideWindowPanel(webView)

	webView.Run()
}

func messageToJson(message interface{}) string {
	result, _ := json.Marshal(message)

	return string(result)
}

func jsonToMessage(asd string) config.Message {
	var message config.Message
	json.Unmarshal([]byte(asd), &message)

	return message
}

func hideWindowPanel(w webview2.WebView) {
	hwnd := win.HWND(w.Window())

	oldStyle := win.GetWindowLong(hwnd, win.GWL_STYLE)
	newStyle := oldStyle &^ (win.WS_CAPTION | win.WS_THICKFRAME)
	win.SetWindowLong(hwnd, win.GWL_STYLE, newStyle)
	win.SetWindowPos(hwnd, 0, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED)
}
