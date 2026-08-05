FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o gc-monitor ./cmd/server

EXPOSE 8080

CMD ["./gc-monitor"]