---
trigger: always_on
---

# Validation Script Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST include a comprehensive validation script at `scripts/validate.sh` that ensures code quality, correctness, and readiness for deployment in both local development and CI/CD environments.

## Required File
**Location**: `scripts/validate.sh`

## File Requirements

### File Attributes
- Must be executable: `chmod +x scripts/validate.sh`
- Bash script with shebang: `#\!/bin/bash`
- Fail-fast mode: `set -euo pipefail`
- Cross-platform compatible: Linux, macOS, CI environments

## Core Structure

### 1. Header and Configuration
```bash
#\!/bin/bash
# 🚀 Validation Script
# Comprehensive validation pipeline for local development and CI/CD
# Compatible with Linux, macOS, and CI environments

set -euo pipefail  # 💥 Fail fast on any error
```

### 2. Terminal Colors and Formatting
MUST define:
- Color variables: RED, GREEN, YELLOW, BLUE, PURPLE, CYAN, NC
- Emoji support for visual output
- ANSI escape codes for color output

```bash
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color
```

### 3. Global Counters
Track execution metrics:
- `TOTAL_STEPS`: Total validation steps
- `PASSED_STEPS`: Successful validations
- `FAILED_STEPS`: Failed validations
- `SKIPPED_STEPS`: Skipped validations
- `WARNING_COUNT`: Warning count
- `START_TIME`: Execution start timestamp

### 4. Configuration Variables
Required environment variables with defaults:
- `MODE`: `local` or `ci` (default: local)
- `COVERAGE_THRESHOLD`: Minimum coverage percentage (default: 80)
- `TEST_TIMEOUT`: Test execution timeout (default: 10m)
- `INTEGRATION_TAG`: Build tag for integration tests (default: integration)
- `SKIP_INTEGRATION`: Flag to disable integration tests (default: false)

```bash
MODE=${1:-"local"}  # local|ci
COVERAGE_THRESHOLD=${COVERAGE_THRESHOLD:-80}
TEST_TIMEOUT=${TEST_TIMEOUT:-10m}
INTEGRATION_TAG=${INTEGRATION_TAG:-integration}
SKIP_INTEGRATION=${SKIP_INTEGRATION:-false}
```

## Required Helper Functions

### Output Functions
MUST implement:

#### print_header()
- Display validation pipeline header
- Show mode and configuration
- Use purple color scheme

#### print_step()
- Announce step execution
- Parameters: step_name, icon
- Use blue/cyan colors

#### print_success()
- Mark step as passed
- Increment PASSED_STEPS counter
- Green checkmark output

#### print_failure()
- Mark step as failed
- Increment FAILED_STEPS counter
- Red X output with error message

#### print_skipped()
- Mark step as skipped
- Increment SKIPPED_STEPS counter
- Yellow forward icon with reason

#### print_warning()
- Display warning message
- Increment WARNING_COUNT
- Yellow warning icon

#### print_info()
- Display informational message
- Cyan info icon

### Execution Function

#### run_step()
Required parameters:
1. `step_name`: Display name
2. `step_function`: Function to execute
3. `icon`: Emoji icon
4. `skip_reason`: Optional skip reason

Features:
- Increments TOTAL_STEPS
- Handles skip logic
- Executes function
- Reports success/failure
- Returns appropriate exit code

## Required Validation Steps

### 1. Environment Check
**Function**: `check_environment()`

MUST verify:
- Go installation and version (minimum 1.19)
- Git installation
- Git repository status
- Environment compatibility

Output:
- Go version information
- Environment check confirmation

### 2. Comprehensive Linting
**Function**: `run_linting()`

MUST:
- Add GOPATH/bin to PATH
- Check for golangci-lint (install if missing)
- Run golangci-lint with timeout
- Handle v2 installation
- Report linting results

Covers:
- Code formatting
- Security scanning
- TODO detection
- Static analysis
- Style checks

### 3. Build Validation
**Function**: `validate_build()`

MUST:
- Clean build cache
- Build all packages (`go build ./...`)
- Run `go mod tidy`
- Verify no uncommitted dependency changes (CI mode)
- Allow dependency updates (local mode)

### 4. Unit Tests
**Function**: `run_unit_tests()`

MUST:
- Run with race detection
- Generate coverage profile
- Use atomic coverage mode
- Apply test timeout
- Run all test packages

```bash
go test -race -timeout=$TEST_TIMEOUT -coverprofile=coverage.out -covermode=atomic ./...
```

### 5. Integration Tests
**Function**: `run_integration_tests()`

MUST:
- Check common test directories: ./test, ./tests, ./integration, ./e2e
- Look for *_test.go files
- Use integration build tag
- Skip if no tests found (normal for libraries)
- Apply test timeout

### 6. Coverage Validation
**Function**: `validate_coverage()`

MUST:
- Check for coverage.out file
- Extract coverage percentage
- Compare against threshold
- Generate HTML report (CI mode)
- Use main package coverage (not total)

Coverage thresholds:
- Green: ≥80%
- Yellow: 60-79%
- Red: <60%

### 7. Documentation Validation
**Function**: `validate_documentation()`

MUST check for:
- Project root README.md (required in CI)
- Optional package READMEs (informational warnings)
- Document validation results

### 8. Final Validation
**Function**: `final_validation()`

MUST verify (CI mode):
- No uncommitted changes in working directory
- No uncommitted changes in staging area
- Clean git status

### 9. Badge Generation
**Function**: `generate_badges()`

MUST generate JSON badge files in `.github/badges/`:

#### Required Badges

##### golangci-lint.json
```json
{"schemaVersion":1,"label":"golangci-lint","message":"0 issues","color":"brightgreen"}
```
- Green: 0 issues
- Red: N issues
- Grey: not available

##### coverage.json
```json
{"schemaVersion":1,"label":"coverage","message":"84.7%","color":"brightgreen"}
```
- Green: ≥80%
- Yellow: 60-79%
- Red: <60%
- Grey: no data

##### go-version.json
```json
{"schemaVersion":1,"label":"Go","message":"go1.24","color":"00ADD8"}
```
- Go blue color (#00ADD8)

##### last-updated.json
```json
{"schemaVersion":1,"label":"last updated","message":"2025-01-07","color":"teal"}
```
- Last commit date from git

##### dependabot.json
```json
{"schemaVersion":1,"label":"security","message":"all clear","color":"brightgreen"}
```
- Red: alerts present
- Blue: pending updates
- Green: all clear
- Yellow: GH CLI required

Additional output:
- lint-results.json: Raw golangci-lint JSON output

## Summary and Reporting

### print_summary()
MUST display:
- Success or failure banner
- Statistics table:
  - Passed steps (green)
  - Failed steps (red)
  - Skipped steps (yellow)
  - Warnings (yellow)
  - Total steps (blue)
  - Execution time (yellow)
- Formatted with box drawing characters
- Color-coded output

Success message:
```
🎉 ALL VALIDATIONS PASSED\! 🎉
✨ Your code is ready to ship\! ✨
```

Failure message:
```
💥 VALIDATION FAILED\! 💥
❌ Please fix the issues above before proceeding
```

## Main Execution Pipeline

### main()
MUST execute in order:
1. print_header
2. Environment Check (required)
3. Comprehensive Linting (required)
4. Build Validation (required)
5. Unit Tests (required)
6. Integration Tests (skippable)
7. Coverage Check (required)
8. Documentation (required)
9. Final Validation (required)
10. Badge Generation (required)
11. print_summary

Exit behavior:
- Exit 0 if FAILED_STEPS == 0
- Exit 1 if FAILED_STEPS > 0

## Script Entry Point
```bash
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
```

Allows script to be:
- Executed directly
- Sourced for testing
- Called from other scripts

## Usage Modes

### Local Development Mode
```bash
./scripts/validate.sh
```

Characteristics:
- More forgiving
- Warnings for missing tools
- Documentation issues non-blocking
- Allows dependency updates
- No git status validation

### CI/CD Mode
```bash
./scripts/validate.sh ci
```

Characteristics:
- Strict validation
- Fails on uncommitted changes
- Requires all tools present
- Documentation required
- Enforces clean git status

## Environment Variables

### Standard Variables
```bash
# Coverage threshold (default: 80)
COVERAGE_THRESHOLD=85 ./scripts/validate.sh

# Test timeout (default: 10m)
TEST_TIMEOUT=15m ./scripts/validate.sh

# Skip integration tests
SKIP_INTEGRATION=true ./scripts/validate.sh

# Integration test build tag (default: integration)
INTEGRATION_TAG=e2e ./scripts/validate.sh
```

## Integration Requirements

### With CI/CD Workflows
GitHub Actions example:
```yaml
- name: Run Validation Pipeline
  run: ./scripts/validate.sh ci
  env:
    COVERAGE_THRESHOLD: 85
    TEST_TIMEOUT: 15m
```

### With Pull Request Template
PR template should reference:
```markdown
- [ ] I have run the validation pipeline (`./scripts/validate.sh`)
```

### With Pre-commit Hooks
Can be used as pre-commit validation:
```bash
#\!/bin/bash
./scripts/validate.sh local
```

## Error Handling

### Fail-Fast Behavior
- Set with `set -euo pipefail`
- Stops on first error in strict mode
- Each validation step returns proper exit codes
- Pipelines fail if any command fails

### Error Messages
- Clear, descriptive error output
- Context about what failed
- Suggestions for fixing (where applicable)
- Color-coded for visibility

## Performance Optimization

### Tool Installation
- Check before installing
- Add GOPATH/bin to PATH
- Cache tool locations
- Only install when needed

### Build Caching
- Clean cache when needed
- Reuse cached data when possible
- Minimize redundant operations

### Parallel Execution
- Steps run sequentially for clarity
- Individual Go test packages may run in parallel
- Race detector may slow tests (necessary trade-off)

## Output Quality

### Visual Design
- Consistent emoji usage
- Color-coded by severity
- Box drawing for headers/footers
- Aligned output columns
- Clear section separators

### Information Density
- Concise step descriptions
- Detailed error messages when needed
- Summary statistics at end
- Timing information
- Progress indicators

## Dependencies

### Required Tools
- Go 1.19+
- Git
- Bash 4.0+

### Optional Tools (Auto-installed)
- golangci-lint v2
- bc (for coverage calculations)
- jq (for JSON parsing, with fallback)
- gh (GitHub CLI, for badge generation)

### Platform Support
- Linux (primary)
- macOS (full support)
- WSL (Windows Subsystem for Linux)
- CI environments (GitHub Actions, GitLab CI, etc.)

## Companion Documentation

### scripts/README.md
MUST include:
- Script purpose and overview
- Usage examples
- Environment variables table
- Validation steps list
- Features description
- Exit codes
- Prerequisites
- CI/CD integration examples
- Troubleshooting section

## Template Reference
Base implementation on `https://github.com/bold-minds/oss/blob/main/scripts`

## Quality Standards

### Code Quality
- Proper error handling
- Consistent naming conventions
- Clear function purposes
- Commented complex logic
- Shellcheck compliant

### User Experience
- Beautiful, colorful output
- Clear progress indicators
- Helpful error messages
- Reasonable defaults
- Flexible configuration

### Reliability
- Cross-platform compatible
- Handles missing tools gracefully
- Fails fast on errors
- Consistent exit codes
- Reproducible results

## Validation Checklist
- [ ] File exists at `scripts/validate.sh`
- [ ] File is executable (755 permissions)
- [ ] Shebang and fail-fast mode set
- [ ] All required helper functions present
- [ ] All required validation steps implemented
- [ ] Badge generation functional
- [ ] Summary reporting complete
- [ ] Both local and CI modes work
- [ ] Color output functional
- [ ] Error handling robust
- [ ] Documentation (scripts/README.md) complete

## Customization Guidelines

### Adding New Validation Steps
1. Create function: `validate_newstep()` or `run_newstep()`
2. Add to main() pipeline
3. Use run_step() wrapper
4. Include proper emoji icon
5. Return 0 for success, 1 for failure
6. Add to scripts/README.md documentation

### Project-Specific Adjustments
- Adjust coverage threshold for project needs
- Add/remove integration test directories
- Customize badge generation logic
- Add project-specific validatio