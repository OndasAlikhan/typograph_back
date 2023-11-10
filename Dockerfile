FROM golang:1.21

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@1.49
RUN go install github.com/swaggo/swag/cmd/swag@1.16

COPY . .
RUN go mod tidy
