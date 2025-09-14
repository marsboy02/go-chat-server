# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

## Reporting a Vulnerability

The Go Chat Server team takes security bugs seriously. We appreciate your efforts to responsibly disclose your findings.

### How to Report Security Issues

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities to us directly:

1. **Email**: Send an email to [security@yourproject.com] with the subject line "SECURITY: [Brief Description]"
2. **GitHub Security Advisory**: Use GitHub's private security advisory feature

### What to Include

Please include the following information in your report:

1. **Description**: A clear description of the vulnerability
2. **Steps to Reproduce**: Step-by-step instructions to reproduce the issue
3. **Impact**: Description of the potential impact
4. **Affected Versions**: Which versions are affected
5. **Suggested Fix**: If you have suggestions for how to fix the issue

### Response Timeline

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Initial Response**: We will send a more detailed response within 7 days
- **Resolution**: We aim to resolve critical vulnerabilities within 30 days

### Security Update Process

1. **Confirmation**: We confirm the vulnerability and determine its severity
2. **Fix Development**: We develop and test a fix
3. **Advisory**: We prepare a security advisory
4. **Release**: We release the fix and publish the advisory
5. **Disclosure**: We coordinate disclosure with the reporter

## Security Best Practices

### For Users

When deploying Go Chat Server:

1. **Always use HTTPS** in production
2. **Keep dependencies updated** regularly
3. **Use environment variables** for sensitive configuration
4. **Run with minimal privileges** (non-root user)
5. **Enable security headers** if using a reverse proxy
6. **Monitor logs** for suspicious activity
7. **Use Docker security best practices** if containerizing

### For Developers

When contributing to Go Chat Server:

1. **Input validation**: Always validate and sanitize user input
2. **Error handling**: Don't expose sensitive information in error messages
3. **Dependencies**: Keep dependencies updated and scan for vulnerabilities
4. **Authentication**: Implement proper authentication and authorization
5. **Logging**: Don't log sensitive information
6. **Testing**: Include security tests in your test suite

## Known Security Considerations

### Current Implementation

- **WebSocket Security**: The current implementation allows connections from any origin. In production, configure proper CORS policies.
- **Rate Limiting**: No built-in rate limiting is implemented. Consider adding rate limiting in production.
- **Authentication**: The current version doesn't include user authentication. Plan to add authentication for production use.
- **Input Sanitization**: User messages should be properly sanitized to prevent XSS attacks.

### Docker Security

- **Non-root User**: The Docker image runs as a non-root user
- **Minimal Image**: Uses a minimal base image (scratch) to reduce attack surface
- **Security Scanning**: Docker images are scanned for vulnerabilities in CI/CD

## Security Resources

- [Go Security Best Practices](https://go.dev/security/)
- [WebSocket Security](https://owasp.org/www-community/attacks/WebSocket_security)
- [Docker Security](https://docs.docker.com/engine/security/)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)

## Disclosure Policy

We follow responsible disclosure practices:

- We will acknowledge security researchers who report vulnerabilities
- We will provide credit in our security advisories (unless you prefer to remain anonymous)
- We will coordinate with you on the disclosure timeline
- We will not take legal action against researchers who follow responsible disclosure practices

Thank you for helping keep Go Chat Server and our users safe!