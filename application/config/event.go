package config

type EventHandler func(data interface{}) Message

var EventHandlers = map[string]EventHandler{
	"getData":      handleGetData,
	"anotherEvent": handleAnotherEvent,
}

func handleGetData(data interface{}) Message {
	return Message{Event: "responseData", Data: "Hello from Go"}
}

func handleAnotherEvent(data interface{}) Message {
	return Message{Event: "anotherResponse", Data: "Another event"}
}
