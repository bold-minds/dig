# Security Policy

## Supported Versions

We actively support the following versions with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 0.x.x   | :white_check_mark: |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability, please follow these steps:

### 1. **Do Not** Create a Public Issue

Please do not report security vulnerabilities through public GitHub issues, discussions, or pull requests.

### 2. Report Privately

Send an email to **security@boldminds.tech** with the following information:

- **Subject**: Security Vulnerability in bold-minds/dig
- **Description**: Detailed description of the vulnerability
- **Steps to Reproduce**: Clear steps to reproduce the issue
- **Impact**: Potential impact and severity assessment
- **Suggested Fix**: If you have ideas for a fix (optional)

### 3. Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days
- **Resolution**: Varies based on complexity, typically within 30 days

### 4. Disclosure Process

1. We will acknowledge receipt of your vulnerability report
2. We will investigate and validate the vulnerability
3. We will develop and test a fix
4. We will coordinate disclosure timing with you
5. We will release a security update
6. We will publicly acknowledge your responsible disclosure (if desired)

## Security Considerations

`dig` is a read-only library with a very small attack surface:

- **No network I/O.** `dig` does not make network calls.
- **No file I/O.** `dig` does not read or write files.
- **No reflection.** `dig` uses concrete type switches only.
- **No external dependencies.** `dig` is pure Go stdlib.
- **Immutable.** `dig` never modifies input data.
- **Nil-safe.** All functions handle `nil` inputs gracefully without panicking.
- **No unchecked bounds.** Slice indexing validates non-negative and in-range indices before access.

### Known Limitations

- `dig` does not parse or validate the structure of input data. It trusts that `any` trees have been produced by a safe source (`json.Unmarshal`, `yaml.Unmarshal`, hand-constructed maps, etc.). If you pass a maliciously-constructed data structure with deeply nested self-references, `dig` will traverse it without looping detection.
- `dig` does not limit path depth or slice size. Callers passing untrusted path arrays of unbounded length should validate the path length before calling.

## Security Updates

Security updates will be:

- Released as patch versions (e.g., 0.1.1)
- Documented in the CHANGELOG.md
- Announced through GitHub releases
- Tagged with security labels

## Acknowledgments

We appreciate responsible disclosure and will acknowledge security researchers who help improve the security of this project.

Thank you for helping keep our project and users safe!
