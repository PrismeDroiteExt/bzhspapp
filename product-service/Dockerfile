FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/web/main.go

EXPOSE 8081

# Vérification des variables d'environnement requises
ENV PRODUCT_DB_HOST=db-product \
    PRODUCT_DB_USER=product_user \
    PRODUCT_DB_NAME=product_db \
    DB_PORT=5432 \
    DB_SSLMODE=disable \
    GIN_MODE=release

CMD ["./main"]