#!/bin/bash

# Variables communes
read -p "Enter timezone (default: Europe/Paris): " TIMEZONE
TIMEZONE=${TIMEZONE:-Europe/Paris}

read -p "Enter Gin mode (default: debug): " GIN_MODE
GIN_MODE=${GIN_MODE:-debug}

# Base de données Auth Service
read -p "Enter Auth DB host (default: db-auth): " AUTH_DB_HOST
AUTH_DB_HOST=${AUTH_DB_HOST:-db-auth}

read -p "Enter Auth DB user (default: auth_user): " AUTH_DB_USER
AUTH_DB_USER=${AUTH_DB_USER:-auth_user}

read -p "Enter Auth DB password: " AUTH_DB_PASSWORD
if [ -z "${AUTH_DB_PASSWORD}" ]
then
    AUTH_DB_PASSWORD=$(openssl rand -base64 32)
fi

read -p "Enter Auth DB name (default: auth_db): " AUTH_DB_NAME
AUTH_DB_NAME=${AUTH_DB_NAME:-auth_db}

# Base de données Product Service
read -p "Enter Product DB host (default: db-product): " PRODUCT_DB_HOST
PRODUCT_DB_HOST=${PRODUCT_DB_HOST:-db-product}

read -p "Enter Product DB user (default: product_user): " PRODUCT_DB_USER
PRODUCT_DB_USER="${PRODUCT_DB_USER:-product_user}"

read -p "Enter Product DB password: " PRODUCT_DB_PASSWORD
if [ -z "${PRODUCT_DB_PASSWORD}" ]
then
    PRODUCT_DB_PASSWORD=$(openssl rand -base64 32)
fi

read -p "Enter Product DB name (default: product_db): " PRODUCT_DB_NAME
PRODUCT_DB_NAME=${PRODUCT_DB_NAME:-product_db}

# JWT Secret
read -p "Enter JWT secret key: " JWT_SECRET
if [ -z "${JWT_SECRET}" ]
then
    JWT_SECRET=$(openssl rand -base64 32)
fi


read -p "What AUTH API URL to use? (default http://localhost:8081) " AUTH_API_URL
AUTH_API_URL="${AUTH_API_URL:-http://localhost:8081}"

read -p "What PRODUCT API URL to use? (default http://localhost:8082) " PRODUCT_API_URL
PRODUCT_API_URL="${PRODUCT_API_URL:-http://localhost:8082}"

while true; do
    read -p "What is the mode of your NODE application? (default development) " NODE_ENV
    NODE_ENV=${NODE_ENV:-development}
    if [ "${NODE_ENV}" == "production" ] || [ "${NODE_ENV}" == "development" ]
        then
        NODE_ENV=${NODE_ENV}
        break
        else
        echo "Invalid NODE_ENV. Please enter either the word production or development."
        exit 1
    fi
done


# Création du fichier .env
cat << EOF > .env
# Common
TIMEZONE=${TIMEZONE}
GIN_MODE=${GIN_MODE}
DB_AUTH_PORT=5433
DB_PRODUCT_PORT=5434
DB_SSLMODE=disable

# Auth Service
AUTH_DB_HOST=${AUTH_DB_HOST}
AUTH_DB_USER=${AUTH_DB_USER}
AUTH_DB_PASSWORD=${AUTH_DB_PASSWORD}
AUTH_DB_NAME=${AUTH_DB_NAME}
JWT_SECRET=${JWT_SECRET}

# Product Service
PRODUCT_DB_HOST=${PRODUCT_DB_HOST}
PRODUCT_DB_USER=${PRODUCT_DB_USER}
PRODUCT_DB_PASSWORD=${PRODUCT_DB_PASSWORD}
PRODUCT_DB_NAME=${PRODUCT_DB_NAME}

# Front Service
AUTH_API_URL=${AUTH_API_URL}
PRODUCT_API_URL=${PRODUCT_API_URL}
NODE_ENV=${NODE_ENV}
EOF

echo "File .env created successfully."

echo "You can now run docker compose up to start the services."
read -p "Run docker compose up? (y/n): " RUN_DOCKER_COMPOSE
if [ "${RUN_DOCKER_COMPOSE}" == "y" ]
then
    docker compose up
fi