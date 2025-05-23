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
- `database.sql` (required: provides the initial schema and data for the MySQL database)

## Environment Setup

The application uses environment variables for configuration. Create a `.env` file in the project root:

```bash
# Database configuration
DB_PASSWORD=your_secure_password
MYSQL_ROOT_PASSWORD=your_secure_root_password

# API configuration
PORT=8080
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_NAME=vehicles
```

Replace `your_secure_password` and `your_secure_root_password` with strong passwords for your environment.

## Setup

You can use the provided `setup.sh` script to automate the Go environment setup, dependency installation, gqlgen installation, and code generation steps:

1. Clone the repository
2. Run the setup script:
   ```bash
   ./setup.sh
   ```
   This will verify your Go installation, download dependencies, install gqlgen, and generate the necessary GraphQL code.

Alternatively, you can perform the steps manually:

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Generate GraphQL code:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```

5. Start the application with docker compose:
   ```bash
   docker compose up --build
   ```

## Docker Usage

You can run the Vehicles API and its MySQL database entirely with Docker. This is the recommended way to get started quickly.

### Build and Run with Docker Compose

1. Ensure Docker and Docker Compose are installed on your system.
2. Build and start the services:
   ```bash
   docker compose up --build
   ```
   This will build the Go application image and start both the API and MySQL containers.

3. The API will be available at [http://localhost:8080](http://localhost:8080).

4. To stop the services, press `Ctrl+C` and then run:
   ```bash
   docker compose down
   ```

### Notes
- The MySQL database is exposed on port 3307 (host) mapped to 3306 (container).
- The database is initialized with `database.sql` (required) if it does not already exist. Make sure this file is present in the project root.
- Data is persisted in a Docker volume (`mysql_data`).
- You can modify environment variables in `docker-compose.yml` as needed.

## API Documentation

The GraphQL Playground is available at http://localhost:8080

### Example Query

```graphql
query {
  vehicles(first: 5) {
    edges {
      node {
        id
        make
        model
        year
        vin
        createdAt
        updatedAt
      }
      cursor
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
```

## Project Structure

- `api/` - API layer (GraphQL, HTTP handlers)
- `domain/` - Domain models and interfaces
- `adapters/` - Adapters for external services (MySQL)
- `data/` - Data access layer 