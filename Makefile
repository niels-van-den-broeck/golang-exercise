format:
	go fmt ./...

test:
	go test ./...

build:
	go build -o ./bin/taskapp cmd/api/main.go

run:
	./bin/taskapp