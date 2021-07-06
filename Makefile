all: mod test vet fmt build run

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
	go build -o bin/gostapagen cmd/gostapagen/main.go

run:
	chmod +x gostapagen.sh
	./gostapagen.sh

install:
	go install -v ./...
