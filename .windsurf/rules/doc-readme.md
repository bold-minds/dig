---
trigger: always_on
---

# README Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST have a comprehensive README.md that serves as the primary documentation and entry point for the project.

## Required Sections

### 1. Title and Badges
- **Project name** as H1 heading using `[PROJECT_NAME]` placeholder
- **Badges** in the following order:
  - License badge
  - Go version badge
  - Latest release badge
  - Last updated badge
  - golangci-lint status badge
  - Coverage badge
  - Dependabot badge

### 2. Brief Description
- One-paragraph summary immediately after badges
- Use placeholder: `[BRIEF_DESCRIPTION_OF_YOUR_PROJECT]`

### 3. Features Section (🚀 Features)
- Bullet list with emoji prefixes
- Highlight key features with bold labels
- Must include "Production Ready" feature mentioning test coverage

### 4. Installation Section (📦 Installation)
- Clear installation command
- For Go projects: `go get github.com/bold-minds/[REPO_NAME]`

### 5. Quick Start Section (🎯 Quick Start)
- Minimal working code example
- Must be copy-paste runnable
- Show basic usage patterns

### 6. Core Features Section (🔧 Core Features)
- Multiple subsections for major features
- Each with code examples
- Use placeholders: `[FEATURE_SECTION_1]`, `[FEATURE_SECTION_2]`, etc.

### 7. Performance Section (🏎️ Performance)
- Benchmark information
- Commands to run benchmarks
- Example benchmark output

### 8. Testing Section (🧪 Testing)
- Commands for running tests
- Commands for race detection
- Commands for benchmarks
- Commands for coverage reports

### 9. API Reference Section (📚 API Reference)
- Document main types and interfaces
- List key functions with signatures
- Use code blocks with proper syntax highlighting

### 10. Contributing Section (🤝 Contributing)
- Link to CONTRIBUTING.md
- Brief welcoming statement

### 11. License Section (📄 License)
- State "MIT License"
- Link to LICENSE file

### 12. Acknowledgments Section (🙏 Acknowledgments)
- Thank dependencies and libraries
- Credit the Go community
- Use placeholders for project-specific acknowledgments

### 13. Related Projects Section (🔗 Related Projects)
- List related projects in the domain
- Include links and brief descriptions
- Use placeholders for project-specific links

## Style Guidelines

### Emoji Usage
- Use emojis sparingly and consistently for section headers
- Standard emojis:
  - 🚀 Features
  - 📦 Installation
  - 🎯 Quick Start
  - 🔧 Core Features
  - 🏎️ Performance
  - 🧪 Testing
  - 📚 API Reference
  - 🤝 Contributing
  - 📄 License
  - 🙏 Acknowledgments
  - 🔗 Related Projects

### Code Examples
- All code examples must be valid and runnable
- Use proper syntax highlighting
- Include package declarations and imports
- Keep examples concise but complete

### Placeholders
- Use ALL_CAPS with underscores for placeholders
- Common placeholders:
  - `[PROJECT_NAME]`
  - `[REPO_NAME]`
  - `[PACKAGE_NAME]`
  - `[BRIEF_DESCRIPTION]`
  - `[FEATURE_SECTION_N]`
  - `[YOUR_CODE_EXAMPLE_HERE]`

### Formatting
- Use H2 (##) for major sections
- Use H3 (###) for subsections
- Use fenced code blocks with language specifiers
- Keep lines readable (no excessive length)

## Template Compliance
- README must follow the structure in `https://github.com/bold-minds/oss/blob/main/README.md`
- All required sections must be present
- Placeholders should be replaced with project-specific content
- Maintain consistent formatting and style

## Quality Standards
- Accurate and up-to-date information
- Clear, concise writing
- Proper grammar and spelling
- Links must be valid
- Badges must point to correct repositories
- Code examples must be tested and working

## Badge Requirements
All badges must use the shields.io format and point to:
- Repository-specific URLs (update from template)
- Correct badge JSON files in `.github/badges/`
- Proper GitHub Actions workflows

## Validation
- All links must be functional
- Table of contents (if present) must be accurate