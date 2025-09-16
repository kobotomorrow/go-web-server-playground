# Copilot Instructions - Go Web Server Playground

## Project Context
- **Purpose**: Learning playground for Go web server fundamentals
- **Language**: Go 1.25.0
- **Dependencies**: Standard library only (`net/http`)
- **Architecture**: Custom Router + Handler pattern
- **Learning Level**: Beginner-focused, educational code

## Code Structure and Patterns

### Core Components
1. **Router**: Custom routing logic, implements `http.Handler`
2. **RESTHandler**: HTTP method-based request delegation
3. **Individual Handlers**: Single-responsibility handlers per action

### Data Flow
```
http.Request → Router.ServeHTTP() → RESTHandler.ServeHTTP() → {UserIndexHandler|UserCreateHandler}.ServeHTTP()
```

### File Organization
- `main.go`: Application entry point and dependency wiring
- `router/`: Custom routing logic and HTTP request handling
- `resthandler/`: REST HTTP method delegation (GET, POST, etc.)
- `handler/`: Individual request handlers by resource/action
- `go.mod`: Module definition only
- No external dependencies or frameworks

## Development Rules

### When adding new handlers:
1. **MUST** follow naming: `{Resource}{Action}Handler` (e.g., `ProductIndexHandler`)
2. **MUST** implement `http.Handler` interface with `ServeHTTP(http.ResponseWriter, *http.Request)`
3. **MUST** be struct type (not function)
4. **MUST** use `http.NotFound(w, r)` for unsupported operations

### When adding new routes:
1. **MUST** register in `main()` function via `router.Handle(path, handler)`
2. **SHOULD** use RESTHandler pattern for CRUD resources
3. **MUST** handle 404 cases explicitly

### Response patterns:
- Current: Plain text with `fmt.Fprintln(w, "message")`
- Future: JSON responses (not yet implemented)
- Error responses: Use `http.NotFound` or appropriate status codes

### Code style requirements:
- Single file architecture (`main.go`)
- No external dependencies beyond standard library
- Educational clarity over production optimization
- Explicit error handling over implicit

### Git Commit Rules:
Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification:

**Format**: `<type>[optional scope]: <description>`

**Common types**:
- `feat`: New feature addition
- `fix`: Bug fix
- `refactor`: Code restructuring without feature/bug changes
- `docs`: Documentation only changes
- `style`: Code formatting, whitespace changes
- `test`: Adding or modifying tests
- `chore`: Build tools, dependencies, or maintenance

**Examples**:
```
feat: add product CRUD endpoints
fix: handle empty request body in user creation
refactor: separate handlers into dedicated packages
docs: update API usage examples in README
```

## Implementation Templates

### New Resource Handler Template:
```go
type {Resource}IndexHandler struct{}
func (h *{Resource}IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "{Resource} Index")
}

type {Resource}CreateHandler struct{}
func (h *{Resource}CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "{Resource} Create")
}

// In main():
rest{Resource}Handler := &RESTHandler{
    Index:  &{Resource}IndexHandler{},
    Create: &{Resource}CreateHandler{},
}
router.Handle("/{resources}", rest{Resource}Handler)
```

### Testing Commands:
```bash
go run main.go                           # Start server
curl http://localhost:8080/{resource}    # GET test
curl -X POST http://localhost:8080/{resource}  # POST test
```
