# Contributing to [PROJECT_NAME]

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## 🚀 Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

### Pull Requests

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## 🧪 Testing

We maintain high test coverage and all contributions should include appropriate tests.

```bash
# Run all tests
go test -v ./...

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Check test coverage
go test -cover ./...
```

## 📝 Code Style

We follow standard Go conventions:

- Use `gofmt` to format your code
- Use `golint` and `go vet` to catch common issues
- Follow effective Go guidelines
- Write clear, self-documenting code
- Add comments for exported functions and types

### Code Organization

- Keep functions focused and small
- Use meaningful variable and function names
- Group related functionality together
- Maintain consistent error handling patterns

## 🐛 Bug Reports

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/bold-minds/[REPO_NAME]/issues/new).

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

### Bug Report Template

```markdown
**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Code Sample**
```go
// Minimal code sample that reproduces the issue
```

**Environment:**
- Go version: [e.g. 1.21]
- OS: [e.g. macOS, Linux, Windows]
- Library version: [e.g. v1.0.0]

**Additional context**
Add any other context about the problem here.
```

## 💡 Feature Requests

We welcome feature requests! Please provide:

- **Use case**: Describe the problem you're trying to solve
- **Proposed solution**: How you envision the feature working
- **Alternatives considered**: Other approaches you've thought about
- **Additional context**: Any other relevant information

## 🏗️ Development Setup

1. **Prerequisites**
   ```bash
   # Go 1.21 or later
   go version
   ```

2. **Clone and setup**
   ```bash
   git clone https://github.com/bold-minds/id.git
   cd id
   go mod download
   ```

3. **Run tests**
   ```bash
   go test -v ./...
   ```

4. **Run benchmarks**
   ```bash
   go test -bench=. ./...
   ```

## 📋 Commit Guidelines

We follow conventional commits for clear history:

- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `test:` adding or updating tests
- `refactor:` code refactoring
- `perf:` performance improvements
- `chore:` maintenance tasks

### Examples

```
feat: add batch generation with custom time ranges
fix: resolve race condition in entropy generation
docs: update README with new API examples
test: add comprehensive validation tests
perf: optimize string conversion allocations
```

## 🔄 Release Process

Releases follow semantic versioning (SemVer):

- **MAJOR**: incompatible API changes
- **MINOR**: backwards-compatible functionality additions
- **PATCH**: backwards-compatible bug fixes

## 🤝 Code of Conduct

### Our Pledge

We pledge to make participation in our project a harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Our Standards

**Positive behavior includes:**
- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

**Unacceptable behavior includes:**
- Trolling, insulting/derogatory comments, and personal attacks
- Public or private harassment
- Publishing others' private information without explicit permission
- Other conduct which could reasonably be considered inappropriate

## 📞 Getting Help

- **Documentation**: Check the README and code comments
- **Issues**: Search existing issues before creating new ones
- **Discussions**: Use GitHub Discussions for questions and ideas

## 🏆 Recognition

Contributors will be recognized in:
- The project's README
- Release notes for significant contributions
- Special thanks in documentation

Thank you for contributing! 🎉
