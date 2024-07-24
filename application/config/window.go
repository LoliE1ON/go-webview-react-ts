package config

type WindowConfig struct {
	Title  string
	Width  int
	Height int
}

var Window = WindowConfig{
	Title:  "Test application",
	Width:  800,
	Height: 600,
}
