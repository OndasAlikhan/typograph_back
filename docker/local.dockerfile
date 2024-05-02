FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0 \
    && go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN #go mod tidy

#ENTRYPOINT ["air", "-c", ".air.toml"]
# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]