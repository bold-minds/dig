---
trigger: always_on
---

# CI/CD Workflows Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST have comprehensive CI/CD workflows that ensure code quality, run tests, and maintain project health through automated validation.

## Required Workflows

### 1. Test Workflow
**Location**: `.github/workflows/test.yaml` or `.github/workflows/test.yml`

**Name**: `tests`

#### Required Triggers
```yaml
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
```

#### Required Permissions
```yaml
permissions:
  contents: write        # For badge commits
  pull-requests: read    # For PR metadata
  security-events: read  # For security alerts
```

#### Testing Matrix Requirements

##### Cross-Platform Testing
MUST test on:
- **ubuntu-latest**: Primary platform, full test suite
- **windows-latest**: Windows compatibility
- **macos-latest**: macOS compatibility

##### Version Testing (Go Projects)
- **Latest stable** (e.g., 1.24): Full testing on all platforms
- **Minimum supported** (e.g., 1.22): Linux only
- **Intermediate versions** (e.g., 1.23): Linux only

Example matrix:
```yaml
strategy:
  matrix:
    include:
      # Primary testing: Latest Go on all platforms
      - go-version: '1.24'
        os: ubuntu-latest
      - go-version: '1.24'
        os: windows-latest
      - go-version: '1.24'
        os: macos-latest
      # Compatibility testing
      - go-version: '1.22'
        os: ubuntu-latest
      - go-version: '1.23'
        os: ubuntu-latest
```

#### Required Test Steps

##### 1. Checkout Code
```yaml
- uses: actions/checkout@v4
```

##### 2. Setup Language Environment
For Go:
```yaml
- name: Setup Go
  uses: actions/setup-go@v5
  with:
    go-version: ${{ matrix.go-version }}
```

##### 3. Install Dependencies
```yaml
- name: Install dependencies
  run: go mod download
```

##### 4. Build Validation
MUST verify:
- Code builds successfully
- Dependencies are tidy
- No uncommitted changes

```yaml
- name: Build validation
  run: |
    go build ./...
    go mod tidy
    git diff --exit-code go.mod go.sum
```

##### 5. Unit Tests
MUST include:
- Full test suite
- Race detection (Unix)
- Coverage reporting

```yaml
- name: Run tests (Unix)
  if: runner.os \!= 'Windows'
  run: go test -v -race -coverprofile=coverage.out ./...

- name: Run tests (Windows)
  if: runner.os == 'Windows'
  run: go test -v -coverprofile="coverage.out" ./...
```

##### 6. Benchmarks
```yaml
- name: Run benchmarks
  run: go test -bench=. -benchmem ./...
```

#### Badge Generation
REQUIRED badges (generated on main branch, primary platform only):

##### Coverage Badge
- Green (≥80%)
- Yellow (60-80%)
- Red (<60%)

##### Go Version Badge
- Current Go version

##### Last Updated Badge
- Last commit date

##### Linter Badge
- golangci-lint status
- Issue count

##### Security Badge
- Dependabot alerts
- Code scanning alerts
- Open security PRs

Example badge generation:
```yaml
- name: Generate badge data
  if: github.ref == 'refs/heads/main' && matrix.os == 'ubuntu-latest' && matrix.go-version == '1.24'
  env:
    GH_TOKEN: ${{ github.token }}
  run: |
    mkdir -p .github/badges
    # Badge generation scripts...
```

##### Badge Commit
MUST commit badges back to main:
```yaml
- name: Commit badges to main branch
  if: github.ref == 'refs/heads/main' && matrix.os == 'ubuntu-latest' && matrix.go-version == '1.24'
  run: |
    git config --global user.name "github-actions[bot]"
    git config --global user.email "github-actions[bot]@users.noreply.github.com"
    git add .github/badges/
    if git diff --staged --quiet; then
      echo "No badge changes to commit"
    else
      git commit -m "Update badges from CI run ${{ github.run_number }} [skip ci]"
      git push origin main
    fi
```

### 2. CodeQL Security Analysis
**Location**: `.github/workflows/codeql.yml`

**Name**: `CodeQL Security Analysis`

#### Required Configuration

##### Triggers
```yaml
on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
  schedule:
    - cron: '0 6 * * 1'  # Weekly on Mondays
  workflow_dispatch:      # Allow manual trigger
```

##### Permissions
```yaml
permissions:
  actions: read
  contents: read
  security-events: write
```

##### Language Matrix
```yaml
strategy:
  fail-fast: false
  matrix:
    language: [ 'go' ]  # Adjust for project language
```

##### Required Steps
```yaml
steps:
- name: Checkout repository
  uses: actions/checkout@v4

- name: Initialize CodeQL
  uses: github/codeql-action/init@v3
  with:
    languages: ${{ matrix.language }}

- name: Autobuild
  uses: github/codeql-action/autobuild@v3

- name: Perform CodeQL Analysis
  uses: github/codeql-action/analyze@v3
  with:
    category: "/language:${{matrix.language}}"
```

## Language-Specific Requirements

### Go Projects
MUST include:
- golangci-lint validation
- Race detection on Unix platforms
- Module tidiness check
- Benchmark tests
- Coverage reporting (>80% target)

### JavaScript/TypeScript Projects
MUST include:
- ESLint validation
- Prettier formatting check
- Unit tests with Jest/Vitest
- Type checking (TypeScript)
- Build verification

## Best Practices

### Performance
- Use caching for dependencies
- Run expensive operations only when necessary
- Parallelize independent jobs
- Use matrix strategies efficiently

### Security
- Use specific action versions (@v4, @v5)
- Limit permissions to minimum required
- Protect sensitive data with secrets
- Regular security scanning

### Reliability
- Add timeouts to prevent hung jobs
- Use fail-fast: false for matrix builds
- Handle platform differences (Windows vs Unix)
- Include retry logic for flaky operations

### Maintainability
- Clear job and step names
- Comments explaining complex logic
- Modular workflow design
- Consistent formatting

## Badge Requirements

### Required Badges in .github/badges/
All as JSON files for shields.io endpoint badges:

1. **coverage.json**: Test coverage percentage
2. **go-version.json**: Current Go version
3. **last-updated.json**: Last commit date
4. **golangci-lint.json**: Linter status
5. **dependabot.json**: Security/dependency status

### Badge Format
```json
{
  "schemaVersion": 1,
  "label": "label text",
  "message": "badge message",
  "color": "color name or hex"
}
```

### Color Standards
- **brightgreen**: ≥80% coverage, no issues, all clear
- **green**: Good status
- **yellow**: Warning status, 60-80% coverage
- **red**: Error status, <60% coverage, issues present
- **blue**: Informational
- **00ADD8**: Go blue

## Workflow File Best Practices

### Naming
- Use lowercase with hyphens: `test.yaml`, `codeql.yml`
- Descriptive names: `test.yaml` not `ci.yaml`
- Consistent extension (prefer `.yaml`)

### Comments
- Explain non-obvious logic
- Document platform-specific behavior
- Note any workarounds or hacks

### Conditional Execution
- Use `if:` conditions appropriately
- Check for main branch before expensive operations
- Handle platform differences

### Error Handling
- Don't fail silently
- Provide useful error messages
- Use `|| echo "default"` for optional commands

## Integration Requirements

### With Pull Request Template
- PR template should reference validation pipeline
- Contributors should verify CI passes

### With Branch Protection
- Require status checks to pass
- Require tests workflow
- Require CodeQL (for security-sensitive projects)

### With Dependabot
- CI must run on Dependabot PRs
- Auto-merge safe PRs (optional)
- Badge updates reflect Dependabot status

## Template Reference 
- Main test workflow: `https://github.com/bold-minds/oss/blob/main/.github/workflows/test.yaml`
- CodeQL workflow: `https://github.com/bold-minds/oss/blob/main/.github/workflows/codeql.yml`

## Quality Standards
- All workflows must be valid YAML
- Required jobs must be present
- Platform matrix appropriate for project
- Badge generation functional
- Security scanning enabled
- Permissions appropriately scoped
- Comments explain complex logic

## Validation Checklist
- [ ] `.github/workflows/` directory exists
- [ ] Test workflow present and named appropriately
- [ ] CodeQL workflow present
- [ ] Test workflow runs on push and PR to main
- [ ] Cross-platform testing configured
- [ ] Version matrix appropriate
- [ ] All required test steps present
- [ ] Race detection enabled (Go)
- [ ] Coverage reporting enabled
- [ ] Badge generation configured
- [ ] Badge commit logic present
- [ ] Security analysis scheduled weekly
- [ ] Permissions appropriately scoped
- [ ] All YAML is valid
- [ ] Workflows trigger correctly

## Maintenance
- Review and optimize slow jobs
- Add tests for new project areas
- Update language versions
- Update security scanning as needed