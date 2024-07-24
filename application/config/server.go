package config

import (
	"strconv"
)

type ServerConfig struct {
	Title        string
	Port         int
	BaseUrl      string
	RendererPath string
	RendererUrl  string
	Debug        bool
}

var Server = ServerConfig{
	Title:        "Application",
	Port:         8080,
	BaseUrl:      "http://localhost",
	RendererPath: "renderer/dist",
	RendererUrl:  "/",
	Debug:        true,
}

func GetServerAddress() string {
	return ":" + strconv.Itoa(Server.Port)
}

func GetRendererBaseUrl() string {
	return Server.BaseUrl + GetServerAddress()
}
