# Go Web Server Playground

A learning playground project for understanding the basic mechanisms of web servers in Go.

## Project Structure

```
go-web-server-playground/
├── .github/
│   └── copilot-instructions.md  # AI development guidelines
├── handler/
│   └── userHandler.go           # User request handlers
├── resthandler/
│   └── resthandler.go           # REST HTTP method delegation
├── router/
│   └── router.go                # Custom routing logic
├── go.mod                       # Go module definition
├── main.go                      # Application entry point
└── README.md                   # This file
```

## Quick Start

### 1. Start Server

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

### 2. Test Endpoints

```bash
# GET request (User Index)
curl http://localhost:8080/users

# POST request (User Create)
curl -X POST http://localhost:8080/users
```

## Architecture

```
HTTP Request
     ↓
router.Router ──→ resthandler.RESTHandler ──→ handler.UserIndexHandler
     ↓                      ↓
404 Response        handler.UserCreateHandler
```

### Package Structure
- **router**: Custom routing logic, implements `http.Handler`
- **resthandler**: HTTP method-based request delegation (GET, POST, etc.)
- **handler**: Individual request handlers by resource/action
