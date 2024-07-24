package config

import (
	"fmt"
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

func (server ServerConfig) GetServerAddress() string {
	return fmt.Sprintf(":%s", strconv.Itoa(server.Port))
}

func (server ServerConfig) GetRendererBaseUrl() string {
	return server.BaseUrl + server.GetServerAddress()
}
