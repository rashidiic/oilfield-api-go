# Oilfield API (Week 1-2)
A Go backend project built to learn and implement a clean API flow using **Gin**, **GORM**, **SQLite**, and **Swagger**.
- **Week 1:** Project setup + Swagger + SQLite + Mock CRUD
- **Week 2:** Real relational schema for oilfield data + migrations + deterministic seed command

---

## Tech Stack
- **Go**
- **Gin** (HTTP server)
- **GORM** (ORM)
- **SQLite** (local database)
- **Swaggo** (Swagger / OpenAPI docs)

---

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
    go run ./cmd/api/main.go
    ```


## Project Structure

```
.
├── cmd/          # Main application entry point and server setup
│   ├── api/
│   ├── main.go
│   ├── seed/
│   │   └── main.go
│   └── data/
│       └── app.db    # SQLite database file
├── docs/             # Swagger documentation files (auto-generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/         # Internal application logic
│   ├── db/           # Database initialization and migration
│   │   └── db.go
│   ├── mock/         # Mock item feature module
│   │   ├── dto.go    # Data Transfer Objects
│   │   ├── handler.go # Request handlers (controller logic)
│   │   ├── model.go   # GORM database model
│   │   └── routes.go  # Route registration
│   │  
│   └── models/
│        ├── oilfield.go 
│        ├── production_reading.go
│        ├── sensor.go
│        └── well.go
│
├── go.mod            # Go module dependencies
└── go.sum            # Dependency checksums
<<<<<<< HEAD

```

---

## Week 1: Mock CRUD
Week 1 implements a simple CRUD module using the `MockItem` model to learn the full flow:
- GORM model + AutoMigrate
- Request/response DTOs
- Handlers + routes
- Swagger documentation

### Available endpoints
- `GET /api/health` → `{ "status": "ok" }`
- Mock CRUD endpoints are available in Swagger UI.

---

## Week 2: Oilfield Relational Schema + Seed

### Schema Overview
The real oilfield schema is implemented with GORM relations:

- **OilField** (1) → (*) **Well**
- **Well** (1) → (*) **Sensor**
- **Sensor** (1) → (*) **ProductionReading**

Entities:
- **OilField:** name, location, operator company, start date
- **Well:** belongs to an oilfield, name, status (active/shut-in/abandoned), drill date, depth
- **Sensor:** belongs to a well, type (pressure/temperature/flowrate), install date, isActive
- **ProductionReading:** belongs to a sensor, timestamp, value, unit

### Migrations
Database schema is applied automatically via GORM **AutoMigrate**.

---

## Getting Started

### 1) Install dependencies
```bash
go mod tidy
```

# Run the API server

```bash
go run ./cmd/api
```

**Server runs on:**

`http://localhost:8080`

**Health check:**

`GET http://localhost:8080/api/health`

**Swagger UI:**

`http://localhost:8080/swagger/index.html`

## Seeding the Database (Week 2)

```bash
go run ./cmd/seed
```

**This will:**

*   **Ensure migrations are applied**
*   **Clear existing oilfield-related tables**
*   **Insert deterministic mock data (oilfields → wells → sensors → readings)**

**SQLite database file:**

`data/app.db`
