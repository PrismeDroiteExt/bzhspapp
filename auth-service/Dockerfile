FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/web/main.go

EXPOSE 8082

CMD ["./main"]