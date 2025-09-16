# Go Web Server Playground

A learning playground project for understanding the basic mechanisms of web servers in Go.

## Project Structure

```
go-web-server-playground/
├── .github/
│   └── copilot-instructions.md  # AI development guidelines
├── go.mod                       # Go module definition
├── main.go                      # Main application
└── README.md                   # This file
```

## Quick Start

### 1. Start Server

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

## Architecture

```
HTTP Request
     ↓
   Router ──→ RESTHandler ──→ UserIndexHandler
     ↓              ↓
404 Response    UserCreateHandler
```
