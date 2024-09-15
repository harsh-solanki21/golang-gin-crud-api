# Golang Gin CRUD APIs

A CRUD (Create, Read, Update, Delete) API built with Go and the Gin web framework.

## Introduction

This project serves as a comprehensive example of building a CRUD API using Go and Gin. It's designed to be a valuable resource for developers looking to learn essential concepts in Go web development, as there's a scarcity of quality learning materials for Go and Gin on the internet.

<br />

## Folder Structure

This structure is more common in web application projects, especially those inspired by MVC (Model-View-Controller) patterns.

```
golang-gin-crud-api/
├── configs/
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   ├── user_controller.go
│   └── product_controller.go
├── docs/
│   └── (Postman collection)
├── middlewares/
│   ├── authenticate.go
│   ├── authorize.go
│   └── error.go
├── models/
│   ├── user.go
│   └── product.go
├── repositories/
│   ├── user_repository.go
│   └── product_repository.go
├── routes/
│   └── routes.go
├── services/
│   ├── auth_service.go
│   ├── product_service.go
│   └── user_service.go
├── tests/
│   ├── product_test.go
│   └── user_test.go
├── utils/
│   ├── jwt.go
│   ├── pagination.go
│   ├── password.go
│   └── response.go
├── validations/
│   ├── product_validator.go
│   └── user_validator.go
├── .github/
│   └── workflows/
│       └── go.yaml
├── .air.toml
├── Dockerfile
├── .dockerignore
├── main.go
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

- `configs/`: Contains database configuration (MongoDB connection config).
- `controllers/`: HTTP request handlers for authentication, users, and products.
- `docs/`: Contains the Postman collection for API documentation.
- `middlewares/`: Custom middleware for authentication, authorization, and error handling.
- `models/`: Data structures for users and products.
- `repositories/`: Data access layer for users and products.
- `routes/`: API route definitions.
- `services/`: Business logic for authentication, users, and products.
- `tests/`: Unit tests for products and users.
- `utils/`: Utility functions for JWT, pagination, password hashing, and response formatting.
- `validations/`: Input validation for products and users.
- `.github/workflows/`: Contains GitHub Actions configuration for CI/CD.
- `.air.toml`: Configuration for hot reloading during development.
- `Dockerfile` and `.dockerignore`: For containerizing the application.
- `main.go`: The entry point of the application.
- `Makefile`: Contains useful commands for building and running the project.

This architecture allows for better testing, easier maintenance, and improved scalability.

<br />

## Folder Structure (Recommended for Go)

This structure is often referred to as the "Standard Go Project Layout" or "Go Project Layout". It's more suitable for:

- Larger, more complex projects
- Projects that may be used as libraries or have reusable components
- Projects that follow a more strict separation between public and private code

- Key features:

  - `cmd/`: Contains main applications for the project.
  - `internal/`: Houses private application and library code.
  - `pkg/`: Contains code that's ok for use by external applications.
  - Clear separation between public (pkg) and private (internal) code.

```
golang-gin-crud-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── item_handler.go
│   ├── models/
│   │   └── item.go
│   ├── repository/
│   │   └── item_repository.go
│   └── service/
│       └── item_service.go
├── pkg/
│   └── db/
│       └── database.go
├── config/
│   └── config.go
├── go.mod
├── go.sum
└── README.md
```

- `cmd/`: Contains the main application entry points.
- `internal/`: Houses the core application code, not meant to be imported by other projects.
  - `handlers/`: HTTP request handlers.
  - `models/`: Data structures and domain models.
  - `repository/`: Data access layer.
  - `service/`: Business logic layer.
- `pkg/`: Reusable packages that can be imported by other projects.
- `config/`: Configuration-related code.

<br />

## Steps to Run

Follow these steps to run the project:

1. Ensure you have Go installed on your system

2. Clone the repository:

   ```
   git clone https://github.com/harsh-solanki21/golang-gin-crud-api
   cd golang-gin-crud-api
   ```

3. Install dependencies:

   ```
   go mod download
   ```

4. Set up your MongoDB database and update the connection string in `.env`.

5. Use the following Makefile commands to run, test, or build the project:

   - Run tests:

     ```
     make test
     ```

   - Run with hot reloading (requires Air to be installed):

     ```
     make air
     ```

   - Run in development mode:

     ```
     make dev
     ```

   - Build the project:

     ```
     make build
     ```

   - Run the built binary:
     ```
     make run
     ```

6. The API should now be running on `http://localhost:5000` (or the port specified in your configuration).

## API Documentation

The API documentation is available as a Postman collection in the `docs/` folder. Import this collection into Postman to explore and test the available endpoints.

## Learning Go and Gin

This project serves as a practical example for learning Go and Gin, covering essential concepts such as:

- Setting up a Go project with modules
- Using the Gin web framework
- Implementing CRUD operations
- Structuring a Go application
- Working with MongoDB in Go
- Authentication and authorization
- Error handling and logging
- Input validation
- API documentation
- Testing in Go
- CI/CD with GitHub Actions
- Containerization with Docker

By exploring this codebase, developers can gain hands-on experience with these concepts and best practices in Go web development.
