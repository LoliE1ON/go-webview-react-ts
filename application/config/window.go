package config

import "github.com/jchv/go-webview2"

type WindowConfig struct {
	Title   string
	Width   int
	Height  int
	WebView webview2.WebView
}
