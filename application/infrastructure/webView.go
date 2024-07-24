package infrastructure

import (
	"encoding/json"
	"github.com/jchv/go-webview2"
	"github.com/lxn/win"
	"github/e1on/go-webview.git/application/config"
)

func CreateWebView() {
	w := webview2.New(config.Server.Debug)
	defer w.Destroy()

	w.SetTitle(config.Window.Title)
	w.SetSize(config.Window.Width, config.Window.Height, webview2.HintNone)
	w.Navigate(config.Server.GetRendererBaseUrl())
	err := w.Bind("sendMessage", func(msg string) string {
		var message = jsonToMessage(msg)

		handler, exists := config.EventHandlers[message.Event]
		if !exists {
			return `{"error": "Unknown event"}`
		}

		response := handler(message.Data)
		return messageToJson(response)
	})
	if err != nil {
		return
	}

	hideWindowPanel(w)

	w.Run()
}

func messageToJson(message config.Message) string {
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
