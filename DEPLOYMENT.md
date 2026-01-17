# Deployment Guide

This guide covers how to deploy the VotePeace Backend to a production environment.

## Prerequisites

- **Server**: Linux (Ubuntu/Debian recommended) or Windows Server.
- **Runtime**: Go 1.21+ installed or Docker.
- **Process Manager**: Systemd or PM2 (recommended for keeping the app alive).

## Method 1: Build from Source

1.  **Clone Source Code**:
    ```bash
    git clone https://github.com/MDF05/VotePeace-Backend.git
    cd VotePeace-Backend
    ```

2.  **Build Binary**:
    ```bash
    go build -o server main.go
    ```

3.  **Run**:
    ```bash
    ./server
    ```
    *(Note: Ensure `votepeace.db` is writable in the directory)*

## Method 2: Docker (Recommended for Future)

*Currently, a Dockerfile is not included, but here is a standard template for Go/Fiber:*

```dockerfile
# Start from the official Golang base image
FROM golang:1.21-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
```

## Production Considerations

-   **Environment Variables**: Ensure all sensitive keys (JWT secrets) are set via environment variables, not hardcoded.
-   **Database**: For high-concurrency production, switch from SQLite to PostgreSQL. Update `database/connect.go` to use `gorm.io/driver/postgres`.
-   **Reverse Proxy**: Use Nginx or Apache in front of the application to handle SSL/TLS (HTTPS) and static file caching.
