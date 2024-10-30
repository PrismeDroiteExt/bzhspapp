FROM golang:1.23.2-alpine

WORKDIR /app

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go mod download

# Generate swagger documentation
RUN swag init -g cmd/web/main.go --parseDependency --parseInternal

RUN go build -o main ./cmd/web/main.go

EXPOSE 8082

CMD ["./main"]