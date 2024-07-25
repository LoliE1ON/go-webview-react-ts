package config

import (
	"fmt"
	"strconv"
)

type ServerConfig struct {
	Port         int
	BaseUrl      string
	RendererPath string
	RendererUrl  string
	Debug        bool
}

func (server ServerConfig) GetServerAddress() string {
	return fmt.Sprintf(":%s", strconv.Itoa(server.Port))
}

func (server ServerConfig) GetRendererBaseUrl() string {
	return server.BaseUrl + server.GetServerAddress()
}
