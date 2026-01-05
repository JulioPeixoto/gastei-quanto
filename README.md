# Gastei Quanto

REST API for parsing and analyzing bank transaction CSV files.

## Features

- CSV file upload and parsing
- Transaction analysis and grouping
- Categorization by expense type
- Spending summaries by category and description
- Total income, expenses, and net balance calculation

## Tech Stack

- Go 1.25.4
- Gin Web Framework
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

## API Endpoints

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
    └── response/
```

## License

MIT

