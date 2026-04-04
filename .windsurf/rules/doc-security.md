---
trigger: always_on
---

# SECURITY Standard for Bold Minds

## Purpose
Every Bold Minds repository MUST have a SECURITY.md file that defines how security vulnerabilities should be reported and handled.

## Required Sections

### 1. Title
- "Security Policy" as H1 heading

### 2. Supported Versions Section
- Table showing which versions receive security updates
- Minimum requirement: show current major version support
- Example format:
  ```markdown
  | Version | Supported          |
  | ------- | ------------------ |
  | 1.x.x   | :white_check_mark: |
  ```

### 3. Reporting a Vulnerability Section
Must include clear reporting process with these subsections:

#### 3.1 Do Not Create Public Issue
- Explicit warning against public disclosure
- State not to use:
  - Public GitHub issues
  - Discussions
  - Pull requests

#### 3.2 Report Privately
Required contact information:
- **Email**: security@boldminds.tech
- **Subject format**: "Security Vulnerability in bold-minds/[REPO_NAME]"

Required information to request:
- Description of vulnerability
- Steps to reproduce
- Impact and severity assessment
- Suggested fix (optional)

#### 3.3 Response Timeline
Must specify:
- Initial response time (within 48 hours)
- Status update time (within 7 days)
- Resolution timeframe (varies by complexity, typically within 30 days)

#### 3.4 Disclosure Process
Must outline these steps:
1. Acknowledge receipt
2. Investigate and validate
3. Develop and test fix
4. Coordinate disclosure timing
5. Release security update
6. Public acknowledgment of reporter (if desired)

### 4. Security Considerations Section
- Project-specific security guidance
- Use placeholder: `[PROJECT_SPECIFIC_SECURITY_SECTION]`
- Include relevant security topics such as:
  - Input validation
  - Authentication
  - Authorization
  - Data protection
  - Rate limiting
  - Error handling

### 5. Best Practices Subsection
- List of security best practices for users
- Minimum 3-5 practices
- Mix of generic and project-specific guidance
- Use placeholders for customization:
  - `[PRACTICE_1]`
  - `[PRACTICE_2]`
  - `[PRACTICE_3]`

### 6. Known Limitations Section
- Document any known security limitations
- Use placeholders:
  - `[LIMITATION_1]`
  - `[LIMITATION_2]`
- Be transparent about constraints

### 7. Security Updates Section
Specify how security updates are communicated:
- Released as patch versions
- Documented in CHANGELOG.md
- Announced through GitHub releases
- Tagged with security labels

### 8. Acknowledgments Section
- Thank security researchers
- Encourage responsible disclosure
- Closing statement thanking contributors

## Key Requirements

### Contact Information
- MUST use: security@boldminds.tech
- No personal emails
- Consistent across all Bold Minds repositories

### Subject Line Format
- Must include repository identifier
- Format: "Security Vulnerability in bold-minds/[REPO_NAME]"
- Replace `[REPO_NAME]` with actual repository name

### Response Commitments
- Be realistic with timeframes
- Standard timeframes:
  - 48 hours for initial response
  - 7 days for status update
  - 30 days typical resolution (complexity-dependent)

### Transparency
- Be honest about supported versions
- Document known limitations
- Don't overcommit on response times

## Style Guidelines

### Tone
- Professional and serious
- Clear and direct
- Reassuring to reporters
- Grateful to security researchers

### Formatting
- Use tables for version support
- Use bold for emphasis on key points
- Use clear hierarchical headings
- Number steps in processes

### Sections Organization
- Information flows from "what to report" to "how we handle it"
- Critical "do not" warnings come first
- Process details follow
- General security guidance at end

## Placeholders
Required placeholders for customization:
- `[REPO_NAME]` - Repository name in URLs and references
- `[PROJECT_SPECIFIC_SECURITY_SECTION]` - Title for security section
- `[PRACTICE_1-3]` - Security practice descriptions
- `[LIMITATION_1-2]` - Known limitation descriptions

## Template Reference
Base structure on `https://github.com/bold-minds/oss/blob/main/SECURITY.md`

## Integration with Other Files
- Link from README.md if security is critical
- Coordinate with GitHub Security Advisories
- Reference in CONTRIBUTING.md where appropriate

## Legal and Compliance
- No legal commitments beyond stated response times
- Clear process protects both reporters and organization
- Responsible disclosure process encourages private reporting

## Quality Standards
- Email address must be valid and monitored
- Response times must be realistic and achievable
- Process must be clear and unambiguous
- Known limitations must be accurate
- Security practices must be current and relevant

## Validation
- File must exist in repository root
- All required sections present
- Contact email is security@boldminds.tech
- Response timeframes are specified
- Repository name is updated from template
- Disclosure process is complete

## Security Best Practices for This File
- Review and update annually
- Keep response times realistic
- Update supported versions table regularly
- Add project-specific security considerations
- Remove placeholder text
