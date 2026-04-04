---
trigger: always_on
---

# Dependabot Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST include Dependabot configuration to automatically maintain dependencies and receive security updates.

## Required File
**Location**: `.github/dependabot.yml`

## Configuration Requirements

### File Structure
- YAML format
- Version 2 syntax
- Multiple package ecosystems supported

### Required Package Ecosystems

#### 1. Primary Language Ecosystem
For Go projects:
```yaml
- package-ecosystem: "gomod"
  directory: "/"
  schedule:
    interval: "weekly"
    day: "monday"
    time: "09:00"
    timezone: "America/Los_Angeles"
```

For other languages, adjust accordingly:
- `npm` for Node.js
- `pip` for Python
- `cargo` for Rust
- `bundler` for Ruby
- `maven` or `gradle` for Java

#### 2. GitHub Actions Ecosystem
REQUIRED for all repositories:
```yaml
- package-ecosystem: "github-actions"
  directory: "/"
  schedule:
    interval: "weekly"
    day: "monday"
    time: "09:00"
    timezone: "America/Los_Angeles"
```

### Required Configuration Options

#### Schedule
- **interval**: `"weekly"` (standard for all Bold Minds repos)
- **day**: `"monday"` (consistent across organization)
- **time**: `"09:00"` (9 AM Pacific Time)
- **timezone**: `"America/Los_Angeles"` (Pacific Time)

#### Pull Request Limits
- **gomod**: `open-pull-requests-limit: 5`
- **github-actions**: `open-pull-requests-limit: 3`
- Prevents PR spam while maintaining updates

#### Reviewers and Assignees
- **reviewers**: `["bold-minds"]`
- **assignees**: `["bold-minds"]`
- Ensures Bold Minds team is notified

#### Commit Messages
For Go modules:
```yaml
commit-message:
  prefix: "deps"
  include: "scope"
```

For GitHub Actions:
```yaml
commit-message:
  prefix: "ci"
  include: "scope"
```

Follows conventional commit format.

#### Labels
For Go modules:
```yaml
labels:
  - "dependencies"
  - "go"
```

For GitHub Actions:
```yaml
labels:
  - "dependencies"
  - "github-actions"
```

#### Grouping (Go Projects)
Group minor and patch updates to reduce PR noise:
```yaml
groups:
  go-dependencies:
    patterns:
      - "*"
    update-types:
      - "minor"
      - "patch"
```

Major updates remain separate for careful review.

## Complete Template

### For Go Projects
```yaml
version: 2
updates:
  # Enable version updates for Go modules
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
      day: "monday"
      time: "09:00"
      timezone: "America/Los_Angeles"
    open-pull-requests-limit: 5
    reviewers:
      - "bold-minds"
    assignees:
      - "bold-minds"
    commit-message:
      prefix: "deps"
      include: "scope"
    labels:
      - "dependencies"
      - "go"
    groups:
      go-dependencies:
        patterns:
          - "*"
        update-types:
          - "minor"
          - "patch"

  # Enable security updates for GitHub Actions
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
      day: "monday"
      time: "09:00"
      timezone: "America/Los_Angeles"
    open-pull-requests-limit: 3
    reviewers:
      - "bold-minds"
    assignees:
      - "bold-minds"
    commit-message:
      prefix: "ci"
      include: "scope"
    labels:
      - "dependencies"
      - "github-actions"
```

## Schedule Philosophy

### Weekly Updates
- Balance between staying current and avoiding PR fatigue
- Monday schedule allows team to review at start of week
- Pacific Time aligns with Bold Minds team timezone

### Alternative Schedules
For high-security projects, consider:
```yaml
schedule:
  interval: "daily"
```

For stable projects with fewer dependencies:
```yaml
schedule:
  interval: "monthly"
```

## Security Considerations

### Automatic Security Updates
- Dependabot automatically creates PRs for security vulnerabilities
- Security updates take priority over version updates
- Review security PRs promptly

### GitHub Security Advisories
- Dependabot integrates with GitHub Security Advisories
- Receives notifications about known vulnerabilities
- Creates PRs even outside scheduled updates

## Best Practices

### PR Management
1. **Review promptly**: Don't let Dependabot PRs pile up
2. **Test thoroughly**: Run full test suite
3. **Group review**: Review multiple minor updates together
4. **Major versions**: Review major version updates carefully

### Customization
- Adjust `open-pull-requests-limit` based on project activity
- Add `ignore` rules for specific dependencies if needed
- Customize labels to match project label scheme

### Ignore Rules (Optional)
To ignore specific updates:
```yaml
ignore:
  - dependency-name: "dependency-name"
    update-types: ["version-update:semver-major"]
```

## Integration with CI/CD

### Automatic Testing
- All Dependabot PRs trigger CI/CD workflows
- Tests must pass before merge
- Use branch protection to enforce

### Badge Integration
- Security badge in README reflects Dependabot status
- Show open dependency update PRs
- Show security alert count

## Multi-Language Projects
For projects with multiple package ecosystems:
```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    # ... config ...
  
  - package-ecosystem: "npm"
    directory: "/web"
    # ... config ...
  
  - package-ecosystem: "github-actions"
    directory: "/"
    # ... config ...
```

## Template Reference
Base configuration on `https://github.com/bold-minds/oss/blob/main/.github/dependabot.yml`

## Quality Standards
- Valid YAML syntax
- All required fields present
- Correct package ecosystem for project language
- Appropriate schedule interval
- Proper reviewer/assignee configuration
- Conventional commit prefixes
- Appropriate labels

## Validation
- [ ] File exists at `.github/dependabot.yml`
- [ ] Version 2 syntax used
- [ ] Primary language ecosystem configured
- [ ] GitHub Actions ecosystem configured
- [ ] Schedule is weekly on Monday
- [ ] Timezone is America/Los_Angeles
- [ ] Reviewers include `bold-minds`
- [ ] Assignees include `bold-minds`
- [ ] Commit message prefixes configured
- [ ] Labels configured appropriately
- [ ] Open PR limits set reasonably
- [ ] Grouping configured for minor/patch (Go projects)
- [ ] YAML is valid and parseable

## Maintenance
- Review configuration quarterly
- Adjust PR limits based on workload
- Update ignore rules as needed
- Monitor Dependabot PR quality
- Ensure team is notified of updates