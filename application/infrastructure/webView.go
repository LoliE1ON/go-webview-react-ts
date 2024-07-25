package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/jchv/go-webview2"
	"github.com/lxn/win"
	"github/e1on/go-webview.git/application/config"
)

func CreateWebView() webview2.WebView {
	return webview2.New(true)
}

func ConfigureWebView(application config.Application) {
	defer application.Window.WebView.Destroy()

	application.Window.WebView.SetTitle(application.Window.Title)
	application.Window.WebView.SetSize(application.Window.Width, application.Window.Height, webview2.HintNone)
	application.Window.WebView.Navigate(application.Server.GetRendererBaseUrl())

	err := application.Window.WebView.Bind("sendMessageToGo", func(json string) string {
		fmt.Println("Received message from JavaScript:", json)

		var message = jsonToMessage(json)

		handler, exists := application.EventHandlerMap[message.Event]
		if !exists {
			return `{"error": "Unknown event"}`
		}

		response := handler(application, message.Data)
		return messageToJson(response)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	hideWindowPanel(application.Window.WebView)

	application.Window.WebView.Run()

	fmt.Println("WebView is set up and running")
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
