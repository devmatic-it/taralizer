# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| Latest release | ✅ Yes |
| Previous major version | ⚠️ Limited support |
| Older versions | ❌ No |

We recommend always using the latest stable release.

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security
vulnerability, please report it responsibly.

### Reporting Process

1. **DO NOT** open a public GitHub issue
2. Email details to: [bator@devmatic-it.com](mailto:bator@devmatic-it.com)
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact assessment
   - Suggested fix (if any)
4. We will acknowledge receipt within 48 hours
5. We will provide a fix or mitigation plan within 7 days
6. Vulnerability will be disclosed after a fix is available

### What to Expect

- **Acknowledgment**: Within 48 hours
- **Status Updates**: Every week until resolved
- **Fix Timeline**: Within 7 days for critical, 30 days for low-severity
- **Disclosure**: After fix is available, with credit to reporter (if desired)

### Safe Harbor

We consider security research and vulnerability corrections to be
legitimate and safe. We will not take legal action against those who
follow this responsible disclosure process. We ask that you:

- Give us reasonable time to fix the issue before disclosure
- Avoid exploiting the vulnerability
- Provide reasonable privacy regarding other users

### Recognition

We gladly credit vulnerability reporters in release notes and security
advisories. If you'd like recognition, please indicate this in your report.

## Security Best Practices

When using Taralizer:

- Keep dependencies updated
- Use TLS for all network communications
- Validate input models before processing
- Run reports in isolated environments for untrusted inputs
- Review generated reports for false positives

## Security-Related Configuration

- All reports are generated locally — no data leaves your system
- PDF generation uses headless Chrome (chromedp)
- HTML reports load mermaid.js from CDN (configurable to local)
- No telemetry or external API calls are made

## Previous Security Advisories

_No previous security advisories._
