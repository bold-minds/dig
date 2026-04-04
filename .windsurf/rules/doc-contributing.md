---
trigger: always_on
---

# CONTRIBUTING Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST have a CONTRIBUTING.md file that provides clear guidelines for community contributions and development workflows.

## Required Sections

### 1. Title and Introduction
- Title: "Contributing to [PROJECT_NAME]"
- Welcoming introduction stating what contributions are accepted:
  - Reporting bugs
  - Discussing code
  - Submitting fixes
  - Proposing features
  - Becoming a maintainer

### 2. Development Process (🚀 Development Process)
- State use of GitHub for hosting and collaboration
- Clear pull request guidelines

### 3. Pull Request Guidelines
Must include these steps:
1. Fork the repo and create branch from `main`
2. Add tests for new code
3. Update documentation for API changes
4. Ensure test suite passes
5. Ensure code lints
6. Submit pull request

### 4. Testing Section (🧪 Testing)
- Statement about maintaining high test coverage
- Commands for:
  - Running all tests (`go test -v ./...`)
  - Race detection (`go test -race ./...`)
  - Benchmarks (`go test -bench=. ./...`)
  - Coverage checking (`go test -cover ./...`)

### 5. Code Style Section (📝 Code Style)
Must specify:
- Follow standard Go conventions
- Use `gofmt` for formatting
- Use `golint` and `go vet`
- Follow Effective Go guidelines
- Write clear, self-documenting code
- Add comments for exported functions and types

### 6. Code Organization Guidelines
- Keep functions focused and small
- Use meaningful names
- Group related functionality
- Maintain consistent error handling

### 7. Bug Reports Section (🐛 Bug Reports)
- Link to GitHub issues for bug tracking
- List what makes a great bug report:
  - Quick summary
  - Steps to reproduce (specific, with sample code)
  - Expected vs actual behavior
  - Notes and investigation

### 8. Bug Report Template
Provide a complete markdown template including:
- Description
- Reproduction steps
- Expected behavior
- Code sample section
- Environment details (Go version, OS, library version)
- Additional context

### 9. Feature Requests Section (💡 Feature Requests)
Required elements:
- Use case description
- Proposed solution
- Alternatives considered
- Additional context

### 10. Development Setup Section (🏗️ Development Setup)
Step-by-step guide:
1. Prerequisites (Go version requirements)
2. Clone and setup commands
3. Run tests
4. Run benchmarks

### 11. Commit Guidelines Section (📋 Commit Guidelines)
Must follow conventional commits:
- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `test:` adding or updating tests
- `refactor:` code refactoring
- `perf:` performance improvements
- `chore:` maintenance tasks

Include examples of good commit messages.

### 12. Release Process Section (🔄 Release Process)
- State adherence to Semantic Versioning (SemVer)
- Define version bumping rules:
  - MAJOR: incompatible API changes
  - MINOR: backwards-compatible functionality additions
  - PATCH: backwards-compatible bug fixes

### 13. Code of Conduct Section (🤝 Code of Conduct)
- Reference or embed key points:
  - Pledge for harassment-free experience
  - Standards for positive behavior
  - Examples of unacceptable behavior
- Can reference full CODE_OF_CONDUCT.md or include abbreviated version

### 14. Getting Help Section (📞 Getting Help)
- Documentation resources
- Issues search guidance
- GitHub Discussions for questions

### 15. Recognition Section (🏆 Recognition)
State where contributors will be recognized:
- Project README
- Release notes
- Documentation

End with thank you message.

## Style Guidelines

### Emoji Usage
Consistent emojis for sections:
- 🚀 Development Process
- 🧪 Testing
- 📝 Code Style
- 🐛 Bug Reports
- 💡 Feature Requests
- 🏗️ Development Setup
- 📋 Commit Guidelines
- 🔄 Release Process
- �� Code of Conduct
- 📞 Getting Help
- 🏆 Recognition

### Tone
- Welcoming and inclusive
- Clear and direct
- Professional yet friendly
- Encouraging to new contributors

### Code Examples
- Provide actual commands that can be copy-pasted
- Use bash syntax highlighting for shell commands
- Use markdown for template examples

## Placeholders
Use these placeholders for repository-specific content:
- `[PROJECT_NAME]` - Full project name
- `[REPO_NAME]` - Repository name in GitHub URL

## Repository-Specific Links
Must update these links from template:
- Bug report issue creation link
- Repository clone URL
- Any other GitHub-specific URLs

## Quality Standards
- All commands must be accurate and tested
- Links must be valid
- Template examples must be properly formatted
- Conventional commit examples must be realistic

## Template Reference
Base structure on `https://github.com/bold-minds/oss/blob/main/CONTRIBUTING.md`

## Integration with Other Files
- Reference CODE_OF_CONDUCT.md if it contains more details
- Link from README.md contributing section
- Align with PR template requirements
- Coordinate with issue templates

## Validation
- File must exist in repository root
- All required sections present
- Commands are valid and accurate
- Links point to correct repository
- Placeholders replaced with actual values
- Conventional commit format is accurate
