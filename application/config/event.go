package config

type EventHandler func(application *Application, data interface{}) interface{}
