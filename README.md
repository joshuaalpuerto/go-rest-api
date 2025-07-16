# Go REST API Learning Project

Welcome to your Go learning journey! This project is designed to help you learn Go by building a REST API from scratch.

## 🎯 Learning Objectives

By the end of this project, you'll understand:
- Go syntax and language fundamentals
- HTTP server creation with Go's standard library
- REST API design principles
- JSON handling in Go
- Error handling patterns
- Go modules and dependency management
- Testing in Go
- Database integration (optional)

## 📋 Prerequisites

- Basic programming knowledge (any language)
- Understanding of HTTP and REST concepts
- Enthusiasm to learn Go! 🚀

## 🛠️ Setup

### 1. Install Go
Visit [golang.org](https://golang.org/downloads/) and download Go for your system.

Verify installation:
```bash
go version
```

### 2. Initialize Go Module
```bash
go mod init go-rest-api
```

### 3. Create your first Go file
```bash
touch main.go
```

## 📁 Project Structure

As you build your API, your project might look like this:
```
go-rest-api/
├── main.go              # Entry point
├── go.mod              # Go module file
├── go.sum              # Dependency checksums
├── handlers/           # HTTP handlers
├── models/             # Data models
├── middleware/         # HTTP middleware
└── tests/              # Test files
```

## 🚀 Getting Started

### Step 1: Basic HTTP Server
Create a simple "Hello World" server in `main.go`:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go World!")
    })
    
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

Run it:
```bash
go run main.go
```

Visit `http://localhost:8080` to see your first Go web server!

## 🎓 Learning Path Suggestions

1. **Basic HTTP Server** - Start with a simple "Hello World"
2. **REST Endpoints** - Add GET, POST, PUT, DELETE endpoints
3. **JSON Handling** - Learn to parse and return JSON
4. **Data Models** - Create structs for your data
5. **Error Handling** - Implement proper error responses
6. **Middleware** - Add logging, CORS, authentication
7. **Testing** - Write unit and integration tests
8. **Database** - Connect to a database (PostgreSQL, MongoDB, etc.)

## 📚 Helpful Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Web Examples](https://gowebexamples.com/)
- [Go REST API Tutorial](https://tutorialedge.net/golang/creating-restful-api-with-golang/)

## 🔧 Common Go Commands

```bash
# Run your program
go run main.go

# Build an executable
go build

# Run tests
go test

# Get dependencies
go get package-name

# Format your code
go fmt

# Check for issues
go vet
```

## 📝 Notes

Keep track of what you learn here:
- [ ] Created basic HTTP server
- [ ] Implemented REST endpoints
- [ ] Added JSON handling
- [ ] Created data models
- [ ] Implemented error handling
- [ ] Added middleware
- [ ] Wrote tests
- [ ] Connected to database

## 🤝 Contributing to Your Learning

- Experiment with different approaches
- Break things and fix them
- Read Go source code
- Join the Go community ([Gophers Slack](https://gophers.slack.com/))

Happy coding! 🎉