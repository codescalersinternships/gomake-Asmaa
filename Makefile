build: test
	go build -o bin/gomake cmd/main.go

test: 
	go test ./...
