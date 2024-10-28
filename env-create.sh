#!/bin/bash

# Prompt for database details
read -p "Please enter your Database root password: " DB_ROOT_PASSWORD
read -p "Please enter your Database username: " DB_USER
read -p "Please enter your Database user password: " DB_USER_PASSWORD
read -p "Please enter your Database name: " DB_NAME
read -p "Please enter the name of your Database host (container name or ip address): " DB_HOST

read -p "What DB SSLMODE to use? (default disable) " DB_SSLMODE
if [ -z "${DB_SSLMODE}" ]
    then
    DB_SSLMODE="'disable'"
    else
    if [ "${DB_SSLMODE}" == "disable" ] || [ "${DB_SSLMODE}" == "enable" ]
        then
        DB_SSLMODE="'${DB_SSLMODE}'"
        else
        echo "Invalid DB SSLMODE. Please enter either the word disable or enable."
        exit 1
    fi
fi

read -p "What is the Database exposed port? (default 5432) " DB_PORT
if [ -z "${DB_PORT}" ]
    then
    DB_PORT=5432
    else
    # Validate that DB_PORT is a number between 1024 and 65535
    if ! [[ "${DB_PORT}" =~ ^[0-9]+$ ]] || [ "${DB_PORT}" -lt 1024 ] || [ "${DB_PORT}" -gt 65535 ]; then
        echo "Invalid port number. Please enter a number between 1024 and 65535."
        exit 1
    fi
fi

# read -p "What is the name of your CloudBeaver admin user? (default administrator) " CB_ADMIN_NAME
# if [ -z "${CB_ADMIN_NAME}" ]
#     then
#     CB_ADMIN_NAME="'administrator'"
# fi

# read -p "What is the password of your CloudBeaver admin user? (default S0mePazzworD) " CB_ADMIN_PASSWORD
# if [ -z "${CB_ADMIN_PASSWORD}" ]
#     then
#     CB_ADMIN_PASSWORD="'S0mePazzworD'"
# fi

read -p "What Timezone to use? (default Europe/Paris) " TIMEZONE
if [ -z "${TIMEZONE}" ]
  then
    TIMEZONE="'Europe/Paris'"
fi


read -p "What is the mode of your GIN application? (default debug) " GIN_MODE
if [ -z "${GIN_MODE}" ]
    then
    GIN_MODE="'debug'"
    else
    if [ "${GIN_MODE}" == "debug" ] || [ "${GIN_MODE}" == "release" ]
        then
        GIN_MODE="'${GIN_MODE}'"
        else
        echo "Invalid GIN mode. Please enter either the word debug or release."
        exit 1
    fi
fi

read -p "Enter JWT secret key for authentication: " JWT_SECRET
if [ -z "${JWT_SECRET}" ]
then
    JWT_SECRET=$(openssl rand -base64 32)
fi

# Create .env file
cat > .env <<EOL
DB_ROOT_PASSWORD=${DB_ROOT_PASSWORD}
DB_HOST=${DB_HOST}    
DB_USER=${DB_USER}
DB_PASSWORD=${DB_USER_PASSWORD}
HOSTNAME=${HOSTNAME}
DB_NAME=${DB_NAME}
TIMEZONE=${TIMEZONE}
DB_SSLMODE=${DB_SSLMODE}
DB_PORT=${DB_PORT}
# CB_ADMIN_NAME=${CB_ADMIN_NAME}
# CB_ADMIN_PASSWORD=${CB_ADMIN_PASSWORD}
GIN_MODE=${GIN_MODE}
JWT_SECRET=${JWT_SECRET}
EOL