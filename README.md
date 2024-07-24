# Application

### Desktop application built with Go and React using WebView and message bus

## Development

Execute the command 
`go run main.go`

## Build

To build execute the command:
`go build -ldflags="-s -w" -o go-webview.git.exe main.go`

For compress the binary:
`./bin/upx.exe --best --lzma go-webview.git.exe`