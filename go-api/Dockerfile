# FROM golang:latest as build
FROM golang:1.8

WORKDIR /go/src/github.com/ccdle12/Blocksage/go-api
COPY ./ ./

RUN go get -u github.com/gorilla/mux \
    && go get github.com/pilu/fresh 

# Fresh runs the API with hot reload / when deploying live should use go run main.go
CMD fresh