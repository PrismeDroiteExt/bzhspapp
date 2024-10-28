# Project structure

This is the structure of the project:

```txt
app/
│
├── cmd/
│ └── web/
│ └── main.go # Entrypoint for the application
│
├── config/
│ └── config.yaml # Configuration file (e.g., DB connection details)
│
├── docs/ # Documentation files (e.g., Swagger)
│ └── arch.md # Project structure documentation
│
├── internal/ # Private application and library code
│ ├── api/
│ │ └── v1/
│ │ ├── routes.go # API route definitions
│ │ ├── brand_routes.go # Brand-specific routes
│ │ └── product_routes.go # Product-specific routes
│ ├── controllers/
│ │ ├── brand_controller.go # Brand controller
│ │ └── product_controller.go # Product controller
│ ├── database/
│ │ └── db.go # Database initialization and connection
│ ├── dto/
│ │ ├── brand_dto.go # Data Transfer Objects for Brand
│ │ └── product_dto.go # Data Transfer Objects for Product
│ ├── models/
│ │ ├── brand.go # Brand model
│ │ ├── cart.go # Cart model
│ │ ├── cartItem.go # CartItem model
│ │ ├── category.go # Category model
│ │ ├── order.go # Order model
│ │ ├── payment.go # Payment model
│ │ ├── product.go # Product model
│ │ └── user.go # User model
│ ├── repository/
│ │ ├── brand_repository.go # Brand repository
│ │ └── product_repository.go # Product repository
│ └── services/
│ ├── brand_service.go # Brand service
│ └── product_service.go # Product service
│
├── scripts/
│ └── dev-entrypoint.sh # Build automation script
│
├── static/ # Static files (e.g., images)
│
├── go.mod # Go module file
└── go.sum # Go module checksum file
```


This structure reflects a typical Go project layout with a clear separation of concerns:

- `cmd/`: Contains the main application entry points.
- `config/`: Holds configuration files.
- `docs/`: Stores documentation files.
- `internal/`: Houses the core application code, including API routes, controllers, models, repositories, and services.
- `scripts/`: Contains utility scripts for development and deployment.
- `static/`: Stores static assets.

The `internal/` directory is particularly important as it contains the bulk of the application logic, organized into subdirectories for different components of the system (API, controllers, models, etc.).
