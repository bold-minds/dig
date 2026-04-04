---
trigger: always_on
---

# GitHub Templates Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST include standardized GitHub templates for issues, pull requests, and code ownership to ensure consistent contributions and reviews.

## Required Templates

### 1. Pull Request Template
**Location**: `.github/pull_request_template.md`

#### Required Sections

##### Description Section
- Brief description placeholder
- Open-ended for contributor input

##### Type of Change Checklist
Must include these options:
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Code refactoring

##### Testing Section
Required checklist items:
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] I have run the validation pipeline (`./scripts/validate.sh`)

##### Checklist Section
Required items:
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] Any dependent changes have been merged and published in downstream modules

##### Related Issues Section
- Placeholder for: "Fixes #(issue number)"
- Placeholder for: "Closes #(issue number)"

##### Additional Notes Section
- Open field for any additional context

#### Template Reference
Base on `https://github.com/bold-minds/oss/blob/main/.github/pull_request_template.md`

### 2. Bug Report Template
**Location**: `.github/ISSUE_TEMPLATE/bug_report.md`

#### Required Elements
Must include:
- Title/name for the template
- Description of when to use it
- Labels to auto-apply

#### Required Sections
- **Describe the bug**: Clear description field
- **To Reproduce**: Numbered steps
- **Expected behavior**: What should happen
- **Screenshots/Code**: If applicable
- **Environment**:
  - Go version
  - OS (Linux/macOS/Windows)
  - Library/tool version
- **Additional context**: Any other relevant information

#### YAML Front Matter
Must include:
```yaml
---
name: Bug Report
about: Create a report to help us improve
title: '[BUG] '
labels: bug
assignees: ''
---
```

### 3. Feature Request Template
**Location**: `.github/ISSUE_TEMPLATE/feature_request.md`

#### Required Sections
- **Is your feature request related to a problem?**: Problem description
- **Describe the solution you'd like**: Desired solution
- **Describe alternatives you've considered**: Alternative approaches
- **Additional context**: Any other relevant information

#### YAML Front Matter
Must include:
```yaml
---
name: Feature Request
about: Suggest an idea for this project
title: '[FEATURE] '
labels: enhancement
assignees: ''
---
```

### 4. CODEOWNERS File
**Location**: `.github/CODEOWNERS`

#### Required Content
Minimum requirement:
```
# Global code owners - all files require review from @bold-minds
* @bold-minds
```

#### Purpose
- Automatically request reviews from specified owners
- Ensure all PRs reviewed by Bold Minds team
- Can be customized per-directory if needed

#### Advanced Usage (Optional)
```
# Global owners
* @bold-minds

# Specific subdirectories or files
/docs/ @bold-minds/docs-team
*.md @bold-minds/docs-team
/.github/ @bold-minds/devops-team
```

## Directory Structure
Required GitHub directory structure:
```
.github/
├── CODEOWNERS
├── pull_request_template.md
└── ISSUE_TEMPLATE/
    ├── bug_report.md
    └── feature_request.md
```

## Style Guidelines

### Markdown Formatting
- Use H2 (##) for major sections
- Use checkboxes for required actions
- Use clear, concise language
- Include examples where helpful

### Placeholders
- Use parentheses for placeholders: `(description here)`
- Use example formatting: `#(issue number)`
- Provide clear guidance on what to fill in

### Labels
Issue templates should auto-apply appropriate labels:
- Bug reports → `bug`
- Feature requests → `enhancement`
- Can include additional labels as needed

## Integration Requirements

### With CI/CD
- PR template should reference validation scripts
- Mention required CI checks
- Note any special testing requirements

### With Documentation
- Templates should guide contributors to update docs
- Reference CONTRIBUTING.md where relevant
- Link to relevant documentation sections

### With Code Review
- CODEOWNERS ensures proper review
- PR template guides reviewer attention
- Checklist ensures quality standards

## Validation Checklist
- [ ] `.github/` directory exists
- [ ] `pull_request_template.md` present with all required sections
- [ ] `ISSUE_TEMPLATE/` directory exists
- [ ] `bug_report.md` present with YAML front matter
- [ ] `feature_request.md` present with YAML front matter
- [ ] `CODEOWNERS` file present
- [ ] CODEOWNERS includes `@bold-minds`
- [ ] All templates use clear, welcoming language
- [ ] Templates reference correct project resources
- [ ] Validation scripts referenced correctly