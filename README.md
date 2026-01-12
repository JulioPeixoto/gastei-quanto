# Gastei Quanto

REST API for parsing and analyzing bank transaction CSV files.

## Features

- User authentication with JWT
- Expense management (CRUD operations)
- CSV file upload and parsing
- Transaction analysis and grouping
- Categorization by expense type
- Spending summaries by category and description
- Total income, expenses, and net balance calculation
- SQLite database for data persistence

## Tech Stack

- Go 1.25.4
- Gin Web Framework
- SQLite (with support for other SQL databases)
- JWT Authentication
- Swagger/OpenAPI documentation

## Getting Started

### Prerequisites

- Go 1.25.4 or higher
- swag CLI for generating API docs

### Installation

Install swag CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Clone and setup:

```bash
git clone <repository-url>
cd gastei-quanto
go mod download
```

### Environment Variables

Create a `.env` file based on `.env.example`:

```bash
JWT_SECRET=your-secret-key-change-in-production
DB_DRIVER=sqlite
DB_DSN=./gastei-quanto.db
```

Available database drivers:
- `sqlite` - SQLite database (default)
- More drivers can be easily added by implementing the `database.Database` interface

### Generate API Documentation

```bash
make swagger
```

or

```bash
swag init -g src/cmd/api/main.go --parseInternal=true
```

### Run the API

```bash
make run
```

or

```bash
go run src/cmd/api/main.go
```

The API will be available at http://localhost:8080

Swagger documentation: http://localhost:8080/swagger/index.html

## Database

The application uses a flexible database layer that allows easy switching between different SQL databases.

### Current Implementation

- SQLite (default)
- Automatic migrations on startup
- Foreign key constraints
- Indexed queries for performance

### Adding a New Database

To add support for a new SQL database (e.g., PostgreSQL, MySQL):

1. Create a new implementation of the `database.Database` interface in `src/pkg/database/`
2. Implement the required methods: `GetDB()`, `Close()`, `Migrate()`
3. Update `main.go` to support the new driver

Example:

```go
case "postgres":
    db, err = database.NewPostgresDatabase(dbDSN)
```

## API Endpoints

### Authentication

**POST /api/v1/auth/register**

Register a new user.

**POST /api/v1/auth/login**

Login and receive JWT token.

**GET /api/v1/auth/me**

Get current user information (requires authentication).

### Expenses

**POST /api/v1/expenses**

Create a new expense (requires authentication).

**GET /api/v1/expenses**

List expenses with optional filters (requires authentication).

Query parameters:
- `start_date` - Filter by start date (YYYY-MM-DD)
- `end_date` - Filter by end date (YYYY-MM-DD)
- `category` - Filter by category
- `type` - Filter by type (income/expense)
- `min_amount` - Minimum amount
- `max_amount` - Maximum amount
- `description` - Search in description

**GET /api/v1/expenses/stats**

Get expense statistics (requires authentication).

**GET /api/v1/expenses/:id**

Get a specific expense (requires authentication).

**PUT /api/v1/expenses/:id**

Update an expense (requires authentication).

**DELETE /api/v1/expenses/:id**

Delete an expense (requires authentication).

**POST /api/v1/expenses/import**

Import transactions from parser (requires authentication).

### Parser

**POST /api/v1/parser/upload/csv**

Upload a CSV file with transactions. Expected format:

```csv
date,title,amount
2025-09-01,Store Name,14.60
```

### Analysis

**POST /api/v1/analysis/transactions**

Analyze transactions and get spending summaries grouped by category and description.

Request body:
```json
{
  "transactions": [
    {
      "date": "2025-09-01T00:00:00Z",
      "description": "Store Name",
      "category": "",
      "amount": 14.60
    }
  ]
}
```

Response includes:
- Total spent
- Total income
- Net balance
- Transaction count
- Breakdown by category (with averages)
- Breakdown by description

## Development

### Build

```bash
make build
```

### Run with live reload

```bash
make dev
```

## Project Structure

```
src/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── middleware.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── repository_sql.go
│   │   ├── routes.go
│   │   └── model.go
│   ├── expense/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── repository_sql.go
│   │   ├── routes.go
│   │   └── model.go
│   ├── parser/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── routes.go
│   │   └── model.go
│   └── analysis/
│       ├── handler.go
│       ├── service.go
│       ├── routes.go
│       └── model.go
└── pkg/
    ├── database/
    │   ├── database.go
    │   └── sqlite.go
    └── response/
```

## License

MIT

