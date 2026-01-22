# Oilfield API Go
[![Ask DeepWiki](https://devin.ai/assets/askdeepwiki.png)](https://deepwiki.com/rashidiic/oilfield-api-go)

This repository contains a RESTful API built with Go, providing mock CRUD (Create, Read, Update, Delete) functionality for "items". It utilizes the Gin web framework for routing, GORM for database interaction with a SQLite backend, and Swagger for API documentation.

## Tech Stack

*   **Go**: The core programming language.
*   **Gin**: A high-performance HTTP web framework.
*   **GORM**: A developer-friendly ORM library for Go.
*   **SQLite**: A self-contained, serverless, transactional SQL database engine.
*   **Swagger**: For generating interactive API documentation.

## Features

*   RESTful API for managing mock items.
*   Full CRUD functionality: Create, List, Update, and Delete items.
*   Persistent data storage using SQLite.
*   Input validation for request payloads.
*   Automatic API documentation via Swagger UI.
*   A health check endpoint to monitor API status.

## Getting Started

### Prerequisites

*   Go (version 1.22 or later)

### Installation & Running

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/rashidiic/oilfield-api-go.git
    cd oilfield-api-go
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Run the application:**
    ```sh
    go run cmd/api/main.go
    ```

The API server will start and listen on port `8080`.

## API Endpoints

The API is accessible under the `/api` prefix.

| Method | Endpoint                  | Description                                |
| :----- | :------------------------ | :----------------------------------------- |
| `GET`  | `/api/health`             | Checks the health of the API.              |
| `POST` | `/api/mock-items`         | Creates a new mock item.                   |
| `GET`  | `/api/mock-items`         | Retrieves a list of all mock items.        |
| `PUT`  | `/api/mock-items/{id}`    | Updates an existing mock item by its ID.   |
| `DELETE` | `/api/mock-items/{id}`    | Deletes a mock item by its ID.             |
| `GET`  | `/swagger/index.html`     | Serves the Swagger API documentation UI.   |

### Example Payloads

**Create Item (`POST /api/mock-items`)**

```json
{
  "name": "My First Item",
  "description": "This is a detailed description of the item."
}
```

**Update Item (`PUT /api/mock-items/{id}`)**
```json
{
  "name": "Updated Item Name",
  "description": "This is an updated description."
}
```

## Project Structure

```
.
├── cmd/api/          # Main application entry point and server setup
│   ├── main.go
│   └── data/
│       └── app.db    # SQLite database file
├── docs/             # Swagger documentation files (auto-generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/         # Internal application logic
│   ├── db/           # Database initialization and migration
│   │   └── db.go
│   └── mock/         # Mock item feature module
│       ├── dto.go    # Data Transfer Objects
│       ├── handler.go # Request handlers (controller logic)
│       ├── model.go   # GORM database model
│       └── routes.go  # Route registration
├── go.mod            # Go module dependencies
└── go.sum            # Dependency checksums