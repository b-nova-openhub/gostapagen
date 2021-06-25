all: test vet fmt mod build run

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

mod:
	go mod tidy
	go mod vendor

build:
	go build -o bin/stapagen cmd/stapagen/main.go

run:
	./bin/stapagen

install:
	mod
	go install -v ./...
