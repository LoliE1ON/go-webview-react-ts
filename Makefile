## Outputs this help screen
help:
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## Development
dev:
	go run main.go

## Build
build:
	go build -ldflags="-s -w -H=windowsgui" -o go-webview.git.exe main.go
	./bin/upx.exe --best --lzma ../go-webview.git.exe

## Update dependencies
refresh:
	go mod tidy