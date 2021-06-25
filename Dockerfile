FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/stapagen

COPY . .

RUN go get -d -v ./... \
    && go install -v ./...

EXPOSE 8080

CMD ["stapagen"]