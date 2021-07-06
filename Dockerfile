FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/gostapagen

COPY . .

RUN go get -d -v ./... \
    && go build -o bin/gostapagen cmd/gostapagen/main.go \
    && go install -v ./... \
    && chmod +x gostapagen.sh

EXPOSE 8080

CMD ["sh", "gostapagen.sh"]