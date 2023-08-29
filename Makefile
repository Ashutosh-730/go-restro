build:
	@echo "Current Directory: $(shell pwd)"
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./main main.go
	@echo "Executable Created: ./main"
	zip main.zip main

clean:
	rm -rf ./main
	rm -rf ./main.zip