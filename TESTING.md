# Testing Guide

Currently, the project does not contain a comprehensive automated test suite. However, we encourage Test-Driven Development (TDD) and adding unit tests for new features.

## Running Tests

To run all Go tests in the repository:

```bash
go test ./...
```

TO run tests with verbose output:

```bash
go test -v ./...
```

## Writing Tests

We use Go's standard `testing` package.

### Example: Controller Test

Create a file `controllers/auth_test.go`:

```go
package controllers

import (
    "net/http/httptest"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert" // Requires: go get github.com/stretchr/testify
)

func TestHealthCheck(t *testing.T) {
    app := fiber.New()
    app.Get("/check", Check)

    req := httptest.NewRequest("GET", "/check", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, 200, resp.StatusCode)
}
```

## Manual Testing (Postman/cURL)

For now, manual testing via Postman is recommended. Import the API routes defined in `API_DOCUMENTATION.md`.
