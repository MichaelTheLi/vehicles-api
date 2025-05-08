# Vehicles API

A GraphQL API for managing vehicles, built with Go, following Domain-Driven Design and Clean Architecture principles.

## Features

- GraphQL API with Echo framework
- MySQL database integration
- Docker and docker-compose support
- Clean Architecture with DDD principles

## Prerequisites

- Go 1.21 or later
- Docker and docker-compose
- MySQL 8.0

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate GraphQL code:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```

4. Start the application with docker-compose:
   ```bash
   docker-compose up --build
   ```

## API Documentation

The GraphQL Playground is available at http://localhost:8080

### Example Query

```graphql
query {
  vehicles {
    id
    make
    model
    year
    vin
    createdAt
    updatedAt
  }
}
```

## Project Structure

- `api/` - API layer (GraphQL, HTTP handlers)
- `domain/` - Domain models and interfaces
- `adapters/` - Adapters for external services (MySQL)
- `data/` - Data access layer 