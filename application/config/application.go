package config

type Application struct {
	Server          ServerConfig
	Window          WindowConfig
	EventHandlerMap map[string]EventHandler
}
