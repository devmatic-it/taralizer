# Quality Constraints

This document defines the quality standards and constraints that all code changes, documentation, and processes within this project must follow. These constraints are mandatory and apply to all team members, agents, and contributors.

## 1. Code Quality Standards

### 1.1 Testing Requirements
- All new code must include unit tests with 85%+ coverage
- Existing tests must continue to pass after any changes
- Test names must describe behavior, not implementation
- No tests should be skipped or disabled unless absolutely necessary

### 1.2 Code Style and Structure
- All code must be compilable at all times 
- No blank identifiers `_ = ...` for error handling
- No suppressing errors with `//nolint`
- Code must follow Go best practices and idioms
- Functions should be small, focused, and single-purpose

### 1.3 Documentation Requirements
- All public APIs must have documentation comments (godoc)
- Every exported function, type, and variable must be documented
- Code changes must include updated documentation where relevant
- README files should be kept current with implementation

## 2. Documentation Standards

### 2.1 Specification Documents
- All specifications must be stored in the `docs/` directory
- Specifications must reference all related files and components
- Technical specifications should include mermaid diagrams where appropriate
- PRD document `docs/prd.md` must contain user stories, success metrics, and technical requirements
- The main technical specification document is `docs/specs.md`

### 2.2 Version Control
- Commit messages must follow the format: `<type>: <short description>` with body explaining *why*
- Atomic commits that do one logical thing
- All documentation changes must be version controlled alongside code

## 3. Security Constraints

### 3.1 Input Handling
- All user inputs must be validated before processing
- No secrets should ever be stored in code or version control
- Parameterized queries to prevent SQL injection
- Secure headers must be set (CSP, HSTS)

### 3.2 Code Security
- No hardcoded credentials or sensitive information in source code
- All external integrations must use secure communication protocols
- Dependency audits should be performed regularly for critical/high vulnerabilities

## 4. Performance Requirements

### 4.1 Execution Time
- CLI commands should complete within 5 seconds for typical use cases
- Parsing large Terraform directories should be optimized for performance

### 4.2 Resource Usage
- Memory usage must remain reasonable and not grow unbounded 
- No N+1 query patterns in database operations
- Efficient algorithms should be used for processing large datasets

## 5. Compatibility and Stability

### 5.1 Backward Compatibility
- Breaking changes to APIs must be clearly documented and versioned
- Existing command structure should remain stable
- No silent fallbacks that hide unclear invariants

### 5.2 Dependencies
- All dependencies must be audited for vulnerabilities regularly
- No unverified framework-specific patterns should be used
- Dependencies must be pinned to specific versions when possible

## 6. Agent Governance Compliance

### 6.1 AGENTS.md Adherence
- All work must follow the defined phases: DEFINE → PLAN → BUILD → VERIFY → REVIEW → SHIP
- Skills must be invoked appropriately when work enters each phase
- All changes to AGENTS.md require explicit approval

### 6.2 Documentation Consistency
- `docs/prd.md`, `docs/spec.md`, and `CONSTRAINTS.md` must be kept in sync
- Changes to one document should be reflected in related documents where appropriate
- All documentation must reference the correct files and code locations

## 7. Review and Approval Requirements

### 7.1 Code Reviews
- All code changes must be reviewed by at least one team member
- Reviewers must check for correctness, simplicity, structure, and tests
- Code reviews should ensure security considerations are addressed

### 7.2 Quality Checks
- All tests must pass before merging to main branch
- Build process should succeed with no warnings
- Lint and type checking must pass without errors

## 8. Exception Handling

### 8.1 Violations
- Any violation of these constraints must be documented with an Exception ID in this file
- Exceptions should only be granted for well-defined, temporary circumstances
- All exceptions must include a justification and expiration date

## 9. Quality Metrics

### 9.1 Coverage Targets
- Unit test coverage: 85%+ for new code
- Integration test coverage: 70%+ for critical paths
- No decrease in overall code coverage

### 9.2 Performance Benchmarks  
- Parsing performance: <5 seconds for directories with 100+ files
- Memory usage: <50MB peak for typical operations

This document serves as the absolute source of truth for quality thresholds and must be consulted before any development work begins.