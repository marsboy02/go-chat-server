# 💬 Go Chat Server

[![CI](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/ci.yml/badge.svg)](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/ci.yml)
[![Docker](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/docker.yml/badge.svg)](https://github.com/YOUR_USERNAME/go-chat-server/actions/workflows/docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/go-chat-server)](https://goreportcard.com/report/github.com/YOUR_USERNAME/go-chat-server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, real-time chat server built with Go and WebSockets, featuring a responsive web interface and Docker containerization.

## ✨ Features

- 🚀 **Real-time messaging** with WebSocket connections
- 🎨 **Modern responsive UI** with smooth animations
- 👥 **Multi-user support** with join/leave notifications
- 🔄 **Auto-reconnection** and connection status indicators
- 📱 **Mobile-friendly** responsive design
- 🐳 **Docker support** for easy deployment
- 🔒 **Security-focused** with non-root container execution
- 📊 **Health checks** and monitoring ready
- 🏗️ **Clean architecture** following Go best practices

## 🚀 Quick Start

### Option 1: Docker (Recommended)
```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/go-chat-server.git
cd go-chat-server

# Start with Docker Compose
make docker-up

# Open your browser to http://localhost:8080
```

### Option 2: Local Development
```bash
# Prerequisites: Go 1.20+ installed

# Clone and setup
git clone https://github.com/YOUR_USERNAME/go-chat-server.git
cd go-chat-server

# Install dependencies
go mod tidy

# Run the server
make run

# Open your browser to http://localhost:8080
```

## 🏗️ Architecture

```
go-chat-server/
├── cmd/server/           # Application entry point
├── internal/             # Private application code
│   ├── client/          # WebSocket client management
│   ├── handler/         # HTTP and WebSocket handlers
│   ├── hub/             # Connection hub (broadcast system)
│   └── message/         # Message types and serialization
├── pkg/config/          # Configuration management
├── web/                 # Frontend assets
│   ├── static/          # CSS, JavaScript files
│   └── templates/       # HTML templates
└── docker/              # Docker configuration files
```

### Core Components

- **Hub Pattern**: Central WebSocket connection manager with thread-safe broadcasting
- **Client Management**: Individual WebSocket connections with proper cleanup
- **Message System**: Structured message types (chat, join, leave, error) with JSON serialization
- **Graceful Shutdown**: Proper server shutdown with configurable timeouts

## 🐳 Docker Deployment

### Development Environment
```bash
make docker-dev-up    # Start development environment
make docker-dev-logs  # View logs
make docker-dev-down  # Stop environment
```

### Production Environment
```bash
make docker-prod-up    # Start with Traefik & SSL
make docker-prod-logs  # View logs
make docker-prod-down  # Stop environment
```

## 🛠️ Development

### Available Commands
```bash
# Development
make run              # Start the server
make dev              # Start in development mode
make test             # Run tests
make test-coverage    # Run tests with coverage

# Code Quality
make fmt              # Format code
make vet              # Run go vet
make lint             # Run golangci-lint
make check            # Run all quality checks

# Building
make build            # Build for current platform
make build-all        # Build for all platforms

# Docker
make docker-build     # Build Docker image
make docker-up        # Start with docker-compose
make docker-logs      # View container logs
```

### Prerequisites
- Go 1.20 or later
- Docker and Docker Compose (for containerized development)

### Configuration
Environment variables:
- `PORT`: Server port (default: 8080)
- `HOST`: Server host (default: localhost)
- `DEV_MODE`: Development mode flag (default: true)
- `TEMPLATE_DIR`: Template directory (default: web/templates)
- `STATIC_DIR`: Static files directory (default: web/static)

## 🚀 Deployment

### Docker Hub / GitHub Container Registry
```bash
# Build and push image
docker build -t your-registry/go-chat-server:latest .
docker push your-registry/go-chat-server:latest

# Run from registry
docker run -p 8080:8080 your-registry/go-chat-server:latest
```

### Production Deployment
1. Copy `docker/.env.example` to `.env` and configure
2. Set up reverse proxy (Traefik configuration included)
3. Configure SSL certificates
4. Run with `make docker-prod-up`

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests for specific package
go test -v ./internal/hub

# Run with race detection
go test -race ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and quality checks (`make check`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

Please read our [Contributing Guidelines](.github/CONTRIBUTING.md) for more details.

## 📋 Roadmap

- [ ] User authentication and authorization
- [ ] Private messaging
- [ ] Chat rooms/channels
- [ ] Message history persistence
- [ ] File sharing
- [ ] Emoji support
- [ ] User avatars
- [ ] Admin panel
- [ ] Rate limiting
- [ ] Message encryption

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [gorilla/websocket](https://github.com/gorilla/websocket) for WebSocket implementation
- [gorilla/mux](https://github.com/gorilla/mux) for HTTP routing
- The Go community for excellent tools and libraries

## 📞 Support

- 📚 [Documentation](docs/)
- 🐛 [Bug Reports](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=bug_report.md)
- 💡 [Feature Requests](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=feature_request.md)
- ❓ [Questions](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=question.md)

---

<div align="center">
Made with ❤️ and Go

⭐ Star this repository if you find it useful!
</div>