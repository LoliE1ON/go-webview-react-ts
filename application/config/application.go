package config

type ApplicationConfig struct {
	Title  string
	Width  int
	Height int
}

var Application = ApplicationConfig{
	Title:  "Test application",
	Width:  800,
	Height: 600,
}
