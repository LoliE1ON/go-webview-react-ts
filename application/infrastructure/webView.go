package infrastructure

import (
	"encoding/json"
	"github.com/jchv/go-webview2"
	"github/e1on/go-webview.git/application/config"
)

func CreateWebView() {
	w := webview2.New(config.Server.Debug)
	defer w.Destroy()

	w.SetTitle(config.Server.Title)
	w.SetSize(config.Application.Width, config.Application.Height, webview2.HintNone)

	w.Navigate(config.GetRendererBaseUrl())
	w.Run()

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
