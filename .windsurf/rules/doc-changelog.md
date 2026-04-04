---
trigger: always_on
---

# CHANGELOG Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST maintain a CHANGELOG.md file that documents all notable changes in a structured, human-readable format.

## Required Standard
- Follow **Keep a Changelog** format (https://keepachangelog.com/en/1.0.0/)
- Adhere to **Semantic Versioning** (https://semver.org/spec/v2.0.0.html)

## File Requirements

### File Name
- Must be named exactly `CHANGELOG.md`
- Placed in the root directory of the repository

### File Structure

#### 1. Title
- "Changelog" as H1 heading

#### 2. Preamble
Must include these two lines:
```markdown
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
```

#### 3. Unreleased Section
- Always maintain an `## [Unreleased]` section at the top
- Collects changes for next release
- Helps contributors know what's coming

#### 4. Version Sections
Each released version must have:
- Header format: `## [VERSION] - YYYY-MM-DD`
- Optional emoji for major releases (🎉)
- Semantic version number (e.g., 1.0.0, 0.2.1)
- ISO 8601 date format (YYYY-MM-DD)

## Change Categories
Each version should organize changes into these categories (as applicable):

### ✨ Added
- New features
- New functionality
- New capabilities

### 🔄 Changed
- Changes to existing functionality
- Updates to behavior
- Modifications

### 🗑️ Deprecated
- Features marked for removal
- Soon-to-be-removed functionality
- Migration warnings

### 🚫 Removed
- Removed features
- Deleted functionality
- Breaking removals

### 🐛 Fixed
- Bug fixes
- Issue resolutions
- Corrections

### 🔒 Security
- Security fixes
- Vulnerability patches
- Security improvements

## Style Guidelines

### Version Entries
- Use H2 (##) for version headers
- Include date in YYYY-MM-DD format
- Versions in descending order (newest first)
- Keep [Unreleased] at top

### Change Entries
- Use H3 (###) for category headers
- Use emoji prefixes for visual clarity
- Bullet points for individual changes
- Start with emoji, then brief description

### Writing Style
- Use present tense or past tense consistently
- Be concise but informative
- Include relevant emoji for context
- Group related changes together

### Emoji Usage
Standard emoji for categories:
- ✨ Added
- 🔄 Changed
- 🗑️ Deprecated
- 🚫 Removed
- 🐛 Fixed
- 🔒 Security

Feature-specific emoji (within entries):
- ��️ Architecture
- 📦 Dependencies
- 🧪 Testing
- �� Validation
- 📊 Metrics/Stats
- 🛡️ Security
- �� Documentation
- ⚡ Performance
- 🎯 Accuracy
- 🧩 Integration
- 📈 Improvements
- 🎨 UI/UX
- 🔧 Configuration

### Links and References
- Link version numbers to release tags
- Link issue numbers (#123)
- Link pull requests where relevant
- Link to documentation for major changes

## Semantic Versioning Rules

### MAJOR version (X.0.0)
- Incompatible API changes
- Breaking changes
- Major rewrites

### MINOR version (0.X.0)
- Backwards-compatible new functionality
- New features
- Enhancements

### PATCH version (0.0.X)
- Backwards-compatible bug fixes
- Security patches
- Minor corrections

## Content Guidelines

### What to Include
- All notable changes
- Breaking changes (emphasized)
- New features
- Bug fixes
- Security updates
- Deprecations
- Performance improvements

### What to Exclude
- Minor refactoring (unless significant)
- Code style changes
- Documentation typo fixes
- Internal changes not affecting users
- Build system changes (unless user-facing)

## First Release (1.0.0)
First production release should:
- Use version 1.0.0
- Include 🎉 emoji
- List all initial features under "Added"
- Document all major components
- Set baseline for future changes

## Pre-release Versions
- Use semantic versioning pre-release tags (e.g., 1.0.0-beta.1)
- Document clearly as pre-release
- Note stability expectations

## Integration with Releases

### GitHub Releases
- Each version in CHANGELOG should correspond to a GitHub release
- Copy relevant CHANGELOG section into release notes
- Link release to version tag

### Version Tagging
- Git tags must match CHANGELOG versions
- Format: v1.0.0 (with 'v' prefix)
- Tag at release commit

## Maintenance

### When to Update
- Update [Unreleased] section as changes are made
- Move to versioned section on release
- Update dates when releasing

### Review Process
- Review CHANGELOG in PR reviews
- Ensure all significant changes are documented
- Verify correct category placement
- Check for accuracy and clarity

## Template Reference
Base structure on `https://github.com/bold-minds/oss/blob/main/CHANGELOG.md`

Example structure:
```markdown
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-08-09 🎉

### ✨ Added
- 🏗️ Complete feature description
- �� Integration with package

### 🐛 Fixed
- Fixed specific bug

### 🔄 Changed
- Updated behavior
```

## Quality Standards
- Accurate dates
- Correct version numbers
- Clear descriptions
- Proper categorization
- Consistent formatting
- Working links
- Complete information

## Validation
- File exists in repository root
- Named exactly `CHANGELOG.md`
- Includes preamble with links
- [Unreleased] section present
- Versions in descending order
- Dates in ISO 8601 format
- Categories properly used
- Semantic versioning followed
- Changes are accurate and complete