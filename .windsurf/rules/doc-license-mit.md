---
trigger: always_on
---

# LICENSE Standard for Bold Minds OSS

## Purpose
Every Bold Minds OSS repository MUST include a LICENSE file that clearly defines the terms under which the software can be used, modified, and distributed.

## Required License
- **MIT License** is the default for all Bold Minds OSS projects
- Ensures maximum compatibility and adoption
- Provides clear permissions and limitations

## File Requirements

### File Name
- Must be named exactly `LICENSE` (no file extension)
- Placed in the root directory of the repository

### Content Structure
The LICENSE file must contain:

1. **License Name**
   - First line: "MIT License"

2. **Copyright Notice**
   - Format: `Copyright (c) [YEAR] Bold Minds`
   - Year should be the current year or year of first publication
   - Copyright holder is always "Bold Minds"

3. **Full License Text**
   - Complete MIT License text
   - Must include all standard clauses:
     - Permission grant
     - Conditions
     - Warranty disclaimer
     - Liability limitation

```
MIT License

Copyright (c) 2025 Bold Minds

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Key Points

### Copyright Year
- Update the year to reflect the current year for new projects
- For existing projects, can keep original year or use year range (e.g., "2024-2025")

### Copyright Holder
- MUST always be "Bold Minds" (not individual developers)
- Maintains organizational consistency
- Simplifies legal matters

### No Modifications
- Do not modify the MIT License text itself
- Only update the copyright year as needed
- Do not add additional terms or conditions

## README Integration
- README.md must include a License section
- Must link to the LICENSE file
- Should include MIT License badge

## Badge Requirement
README should include the MIT License badge:
```markdown
[\![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
```

## Legal Compliance
- Ensures code can be freely used, modified, and distributed
- Protects Bold Minds from liability
- Grants necessary permissions to users
- Maintains compatibility with most open source projects

## Validation
- File must exist in repository root
- Must be named exactly "LICENSE"
- Copyright year must be present
- Copyright holder must be "Bold Minds"
- Complete MIT License text must be present
- No unauthorized modifications to license text