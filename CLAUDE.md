# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a modern Go-based real-time chat server that provides a WebSocket-powered chat application with a responsive web interface. The server follows Go project layout standards and implements a hub-and-spoke pattern for managing WebSocket connections.

## Development Commands

### Quick Start
```bash
# Install dependencies
go mod tidy

# Run the development server
make run
# or with development mode
make dev

# Build the application
make build

# Run all checks (format, vet, lint, test)
make check
```

### Common Operations
```bash
# Format code
make fmt

# Run tests
make test

# Run tests with coverage
make test-coverage

# Lint code (requires golangci-lint)
make lint

# Clean build artifacts
make clean

# Build for production
make build

# Cross-platform builds
make build-all
```

## Architecture

### Directory Structure
```
go-chat-server/
├── cmd/server/           # Application entry point
├── internal/             # Private application code
│   ├── client/          # WebSocket client management
│   ├── handler/         # HTTP and WebSocket handlers
│   ├── hub/             # Connection hub (broadcast system)
│   └── message/         # Message types and serialization
├── pkg/config/          # Configuration management
└── web/                 # Frontend assets
    ├── static/          # CSS, JavaScript files
    └── templates/       # HTML templates
```

### Core Components

**Hub Pattern**: The `hub` package manages all WebSocket connections using channels for thread-safe communication. It handles client registration/deregistration and message broadcasting.

**Client Management**: Each WebSocket connection is wrapped in a `Client` that handles read/write operations in separate goroutines with proper cleanup and error handling.

**Message System**: Structured message types (chat, join, leave, error) with JSON serialization for client-server communication.

**Graceful Shutdown**: The server implements proper graceful shutdown with configurable timeouts.

### Frontend Architecture

**Responsive Design**: Modern CSS with flexbox layout, animations, and mobile-first responsive design.

**WebSocket Client**: Class-based JavaScript client with automatic reconnection, message queuing, and connection status indicators.

**Real-time Features**: Live user count, join/leave notifications, message timestamps, and auto-scroll.

## Configuration

Environment variables:
- `PORT`: Server port (default: 8080)
- `HOST`: Server host (default: localhost)
- `DEV_MODE`: Development mode flag (default: true)
- `TEMPLATE_DIR`: Template directory (default: web/templates)
- `STATIC_DIR`: Static files directory (default: web/static)

## Running the Application

1. Start the server: `make run`
2. Open browser to `http://localhost:8080`
3. Enter a username and start chatting
4. Open multiple tabs/browsers to test multi-user functionality

## Testing

The project structure supports unit testing for each package. Run `make test` to execute all tests.

## Dependencies

- `github.com/gorilla/websocket`: WebSocket implementation
- `github.com/gorilla/mux`: HTTP router (for potential future expansion)

## Development Notes

- Follow Go naming conventions and package organization
- WebSocket connections are managed with proper cleanup and error handling
- The frontend uses modern JavaScript (ES6+) without external frameworks
- CSS follows BEM methodology and uses CSS custom properties
- All user input is properly escaped and validated