---
trigger: always_on
---

# golangci-lint Standard for Bold Minds Go Projects

## Purpose
Every Bold Minds Go repository MUST include a `.golangci.yml` configuration file that enforces code quality, security, and best practices through automated linting.

## Required File
**Location**: `.golangci.yml` (in repository root)

## Configuration Requirements

### Version
Must specify golangci-lint version 2:
```yaml
version: "2"
```

### Run Configuration
```yaml
run:
  relative-path-mode: wd
```

## Required Linters

### Core Linters (MUST Enable)

#### 1. depguard
- Enforces allowed dependencies
- Prevents unwanted imports
- Maintains dependency hygiene

```yaml
depguard:
  rules:
    main:
      files:
        - $all
      allow:
        - $gostd
        - github.com/oklog/ulid
        - github.com/stretchr/testify
        - github.com/bold-minds/[REPO_NAME]
        # Add project-specific dependencies
```

**IMPORTANT**: Update `bold-minds/[REPO_NAME]` with actual repository name.

#### 2. errcheck
- Ensures error returns are checked
- Critical for reliability
- Includes type assertion checking

```yaml
errcheck:
  check-type-assertions: true
```

#### 3. godox
- Finds TODO/FIXME comments
- Helps track technical debt
- Encourages issue tracking

#### 4. gosec
- Security vulnerability scanner
- Detects common security issues
- Essential for production code

#### 5. govet
- Official Go static analyzer
- Catches suspicious constructs
- Reliable and comprehensive

```yaml
govet:
  disable:
    - fieldalignment  # Disabled for readability over performance
  enable-all: true
  settings:
    shadow:
      strict: true
```

#### 6. ineffassign
- Detects ineffectual assignments
- Catches potential bugs
- Improves code clarity

#### 7. staticcheck
- Comprehensive static analysis
- Industry standard
- Catches many bug classes

#### 8. unused
- Finds unused code
- Reduces code bloat
- Improves maintainability

### Linter Configuration

#### Default: None
```yaml
linters:
  default: none
  enable:
    - depguard
    - errcheck
    - godox
    - gosec
    - govet
    - ineffassign
    - staticcheck
    - unused
```

**Philosophy**: Explicitly enable linters rather than using defaults. Provides clarity and control.

### Linter Settings

#### cyclop (if enabled)
```yaml
cyclop:
  max-complexity: 30
  package-average: 10
```

#### funlen (if enabled)
```yaml
funlen:
  lines: 100
  statements: 50
  ignore-comments: true
```

#### gocognit (if enabled)
```yaml
gocognit:
  min-complexity: 20
```

#### gocritic (if enabled)
```yaml
gocritic:
  settings:
    captLocal:
      paramsOnly: false
    underef:
      skipRecvDeref: false
```

#### mnd (magic number detector, if enabled)
```yaml
mnd:
  ignored-functions:
    - args.Error
    - flag.Arg
    - flag.Duration.*
    - flag.Float.*
    - flag.Int.*
    - flag.Uint.*
    - os.Chmod
    - os.Mkdir.*
    - os.OpenFile
    - os.WriteFile
    - prometheus.ExponentialBuckets.*
    - prometheus.LinearBuckets
```

#### nakedret (if enabled)
```yaml
nakedret:
  max-func-lines: 0  # Disallow naked returns
```

#### nolintlint
```yaml
nolintlint:
  require-explanation: true
  require-specific: true
  allow-no-explanation:
    - funlen
    - gocognit
    - lll
```

#### sloglint (if enabled)
```yaml
sloglint:
  no-global: all
  context: scope
```

#### usetesting (if enabled)
```yaml
usetesting:
  os-temp-dir: true
```

## Exclusions

### Generated Code
```yaml
exclusions:
  generated: lax
```

### Presets
Standard exclusion presets:
```yaml
presets:
  - comments
  - common-false-positives
  - legacy
  - std-error-handling
```

### Rule-Based Exclusions

#### Test Files
Relax certain linters in test files:
```yaml
rules:
  - linters:
      - bodyclose
      - dupl
      - errcheck
      - funlen
      - goconst
      - gosec
      - noctx
      - wrapcheck
    path: _test\.go
```

**Rationale**: Test code has different requirements than production code.

#### TODOs and Inspection Comments
```yaml
- linters:
    - godot
  source: (noinspection|TODO)
- linters:
    - gocritic
  source: //noinspection
```

### Path Exclusions
```yaml
paths:
  - third_party$
  - builtin$
  - examples$
```

## Issues Configuration

### Max Same Issues
```yaml
issues:
  max-same-issues: 50
```

Prevents overwhelming output from repeated issues.

## Formatters Configuration

### Formatter Exclusions
```yaml
formatters:
  exclusions:
    generated: lax
    paths:
      - third_party$
      - builtin$
      - examples$
```

## Customization Guidelines

### Adding Dependencies (depguard)
When adding new dependencies:
1. Add to `depguard.rules.main.allow` list
2. Use exact module paths
3. Include only direct dependencies
4. Document why if unusual

Example:
```yaml
allow:
  - $gostd
  - github.com/stretchr/testify      # Testing
  - github.com/your/new-dependency   # Brief reason
  - github.com/bold-minds/[REPO_NAME]
```

### Enabling Additional Linters
Optional linters to consider:
- `bodyclose`: HTTP response body closing
- `cyclop`: Cyclomatic complexity
- `dupl`: Duplicate code detection
- `funlen`: Function length limits
- `gocognit`: Cognitive complexity
- `goconst`: Repeated string constants
- `gocyclo`: Cyclomatic complexity
- `lll`: Line length limits
- `noctx`: HTTP requests without context
- `wrapcheck`: Error wrapping

Add to `enable` list:
```yaml
enable:
  - depguard
  - errcheck
  # ... existing linters ...
  - bodyclose  # New linter
```

### Disabling Specific Rules
Use inline comments sparingly:
```go
//nolint:gosec // G404: Weak random generator acceptable for non-crypto use
rand.Intn(100)
```

**Requirements** (via nolintlint):
- Must include explanation
- Must be specific (name the linter)

## Integration with CI/CD

### GitHub Actions
Test workflow must run golangci-lint:
```yaml
- name: Run golangci-lint
  run: golangci-lint run
```

### Badge Generation
CI generates linter badge showing issue count.

### Pre-commit Hook (Optional)
```bash
#\!/bin/bash
golangci-lint run --new-from-rev=HEAD~1
```

## Common Issues and Solutions

### Issue: Too Many false positives
**Solution**: Add to exclusions or adjust linter settings

### Issue: Slow linting
**Solution**: 
- Reduce enabled linters
- Use `--fast` flag locally
- Optimize exclusions

### Issue: Conflicting rules
**Solution**: Disable one linter or adjust settings

## Best Practices

### Local Development
Run before committing:
```bash
golangci-lint run
```

Fast check:
```bash
golangci-lint run --fast
```

New code only:
```bash
golangci-lint run --new-from-rev=HEAD~1
```

### CI/CD
- Run on every PR
- Block merge on failures
- Generate and display badge
- Cache lint results

### Code Review
- Fix lint issues before requesting review
- Don't silence linters without good reason
- Document nolint directives
- Keep configuration updated

## Template Reference
Base configuration on `https://github.com/bold-minds/oss/blob/main/.golangci.yml`

## Quality Standards
- Valid YAML syntax
- Version 2 specified
- All required linters enabled
- depguard rules updated with correct repo name
- Appropriate exclusions configured
- Settings optimized for project
- No overly permissive exclusions

## Validation Checklist
- [ ] File exists at `.golangci.yml`
- [ ] Version 2 specified
- [ ] Required linters enabled
- [ ] depguard configured
- [ ] Repository name updated in depguard
- [ ] Test file exclusions present
- [ ] Generated code excluded
- [ ] Common false positive presets applied
- [ ] nolintlint configured properly
- [ ] YAML is valid
- [ ] Lints pass successfully

## Maintenance
- Update enabled linters as project matures
- Add new dependencies to depguard allow list
- Review and adjust complexity limits
- Keep golangci-lint version updated
- Monitor for new useful linters
- Refine exclusions based on experience

## Benefits
- Consistent code quality
- Early bug detection
- Security issue prevention
- Best practice enforcement
- Reduced code review burden
- Professional codebase
- Easier onboarding for new contributors

## Reference Documentation
- golangci-lint docs: https://golangci-lint.run/
- Linter descriptions: https://golangci-lint.run/usage/linters/
- Configuration reference: https://golangci-lint.run/usage/configuration/
