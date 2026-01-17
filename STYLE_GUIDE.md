# Go Style Guide

To ensure consistency and readability, we adhere to standard Go best practices.

## Tools

We use the standard `go fmt` tool for formatting.
**Always run `go fmt ./...` before committing.**

## Conventions

### Naming
-   **Files**: Snake_case (e.g., `user_controller.go`) or simple lowercase (e.g., `main.go`).
-   **Functions**: CamelCase. Exported functions start with Uppercase (e.g., `GetUsers`), unexported with lowercase (e.g., `hashPassword`).
-   **Variables**: CamelCase. Short names (e.g., `i`, `c`) are acceptable in short scopes/loops.

### Error Handling
-   Handle errors explicitly. Do not ignore them using `_` unless absolutely necessary.
-   Use proper error wrapping if needed, or return clean JSON error responses in controllers.

```go
// Bad
f, _ := os.Open("filename.ext")

// Good
f, err := os.Open("filename.ext")
if err != nil {
    return err
}
```

### Structs & JSON
-   Use struct tags for JSON serialization.
-   Keep models simple and separate from controller logic.

```go
type User struct {
    Name string `json:"name"` // Lowercase JSON key
}
```
