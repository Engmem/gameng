FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/gameng/main.go

EXPOSE 5051

CMD ["./main"]