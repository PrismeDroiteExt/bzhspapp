# BreizhSport Backend API

This is the backend API for BreizhSport, a sports equipment e-commerce platform.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Technologies Used](#technologies-used)
3. [Prerequisites](#prerequisites)
4. [Getting Started](#getting-started)
5. [API Endpoints](#api-endpoints)
6. [Making Requests](#making-requests)
7. [Development](#development)
8. [Contributing](#contributing)
9. [License](#license)

## Project Structure

The project follows a typical Go application structure that you can find in [arch.md](./app/docs/arch.md) file.

## Technologies Used

- Go 1.23
- Gin Web Framework
- GORM (Go Object Relational Mapper)
- PostgreSQL
- Docker & Docker Compose

## Prerequisites

- Docker
- Docker Compose
- Go 1.23 (for local development)

## Getting Started

1. Clone the repository:

   ```bash
   git clone --recurse-submodules git@github.com:PrismeDroiteExt/bzhspapp.git
   cd bzhspapp
   ```

2. Create a `.env` file in the root directory. You can use the `env-create.sh` script to help you set up the environment variables:

   ```bash
   ./env-create.sh
   ```

3. Build and start the Docker containers:
   ```bash
   docker compose -f docker-compose.dev.local.yml up --build
   ```

The `product service` should now be running and accessible at `http://localhost:8081`.

The `auth service` should now be running and accessible at `http://localhost:8082`.

4. To test the services, you can access the `swagger` UI at `http://localhost:8081/swagger/index.html` for the product service and `http://localhost:8082/swagger/index.html` for the auth service.



## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
