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
PRODUCT_DB_USER=${PRODUCT_DB_USER:-product_user}

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

# Création du fichier .env
cat << EOF > .env
# Common
TIMEZONE=${TIMEZONE}
GIN_MODE=${GIN_MODE}
DB_PORT=5432
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
EOF