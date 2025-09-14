# Contributing to Go Chat Server

First off, thank you for considering contributing to Go Chat Server! It's people like you that make this project better.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible using our [bug report template](.github/ISSUE_TEMPLATE/bug_report.md).

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please use our [feature request template](.github/ISSUE_TEMPLATE/feature_request.md) and include:

- A clear and descriptive title
- A step-by-step description of the suggested enhancement
- Specific examples to demonstrate the steps
- A description of the current behavior and what behavior you expected to see instead
- An explanation of why this enhancement would be useful

### Pull Requests

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes
4. Add tests for any new functionality
5. Ensure all tests pass
6. Update documentation if necessary
7. Submit a pull request

Please follow our [pull request template](.github/pull_request_template.md).

## Development Setup

### Prerequisites

- Go 1.20 or later
- Docker and Docker Compose (optional, for containerized development)
- Make (optional, but recommended)

### Getting Started

1. **Fork and clone the repository:**
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-chat-server.git
   cd go-chat-server
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the development server:**
   ```bash
   make run
   # or
   go run ./cmd/server
   ```

4. **Run tests:**
   ```bash
   make test
   # or
   go test ./...
   ```

### Development Workflow

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes and test them:**
   ```bash
   # Run tests
   make test

   # Run linting
   make lint

   # Format code
   make fmt

   # Run all checks
   make check
   ```

3. **Commit your changes:**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

4. **Push to your fork:**
   ```bash
   git push origin feature/your-feature-name
   ```

5. **Create a Pull Request**

### Docker Development

For containerized development:

```bash
# Start development environment
make docker-dev-up

# View logs
make docker-dev-logs

# Stop environment
make docker-dev-down
```

## Style Guidelines

### Go Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Write comments for exported functions and packages
- Keep functions small and focused
- Handle errors appropriately
- Write tests for new functionality

### Commit Message Guidelines

We follow the [Conventional Commits](https://conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

Examples:
```
feat: add user authentication
fix: resolve WebSocket connection timeout
docs: update API documentation
refactor: improve hub message broadcasting
```

### Testing

- Write unit tests for new functions
- Write integration tests for new features
- Ensure all tests pass before submitting PR
- Maintain or improve code coverage

### Documentation

- Update README.md if needed
- Update CLAUDE.md for development-related changes
- Comment complex code sections
- Update API documentation if applicable

## Project Structure

```
go-chat-server/
├── cmd/server/           # Application entry point
├── internal/             # Private application code
│   ├── client/          # WebSocket client management
│   ├── handler/         # HTTP and WebSocket handlers
│   ├── hub/             # Connection hub
│   └── message/         # Message types
├── pkg/config/          # Configuration management
├── web/                 # Frontend assets
└── docker/              # Docker configurations
```

## Additional Notes

### Issue and PR Labels

We use labels to categorize issues and PRs:

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Improvements or additions to documentation
- `good first issue`: Good for newcomers
- `help wanted`: Extra attention is needed
- `invalid`: This doesn't seem right
- `question`: Further information is requested
- `wontfix`: This will not be worked on

### Getting Help

If you need help, you can:

1. Check existing issues and documentation
2. Create a [question issue](https://github.com/YOUR_USERNAME/go-chat-server/issues/new?template=question.md)
3. Join our discussions (if available)

### Recognition

Contributors will be recognized in our README.md file and release notes.

Thank you for contributing to Go Chat Server! 🎉