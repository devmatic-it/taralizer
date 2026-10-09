# Agent Governance

YOU MUST follow the phase order DEFINE -> PLAN ->  BUILD -> VERIFY -> REVIEW -> SHIP
Each phase has a dedicated skill that enforces its discipline. Invoke the relevant skill by name when the work enters that phase.

---

## Phase 1: DEFINE — `/spec`

**Goal:** Turn vague ideas into sharp, testable specifications before writing code.

### Skills

| Skill | When to Use |
|-------|-------------|
| **idea-refine** | Raw ideas, vague requirements, stress-testing plans before committing |
| **interview-me** | Underspecified asks — "build me X" without "for whom" or "why now" |
| **spec-driven-development** | New projects, features, or significant changes with no specification yet |

### Discipline

1. **Idea Refine** — Restate as a "How Might We" problem. Ask 3-5 sharpening questions. Generate 5-8 variations. Produce a markdown one-pager with Problem Statement, Recommended Direction, Key Assumptions, MVP Scope, and **Not Doing** list.

2. **Interview** — One question at a time. Build ~95% confidence about intent before proceeding.

3. **Spec** — Write a PRD with objectives, scope, success criteria, boundaries (Always/Ask First/Never), and a capability map if the requirement spans multiple testable modules.

**Do NOT proceed to planning until the spec is approved and saved to a file.**

---

## Phase 2: PLAN — `/plan`

**Goal:** Break the approved spec into ordered, implementable tasks.

### Skills

| Skill | When to Use |
|-------|-------------|
| **planning-and-task-breakdown** | Spec or clear requirements exist; break work into implementable tasks |
| **constraint-driven-development** | Establish quality bar as a written contract before planning |

### Discipline

1. **Constraints First** — Define quality thresholds (coverage %, performance, security) as a written contract in `CONSTRAINTS.md`. Use `constraint-driven-development` to interview on dimensions that matter and record defaults when the user has no number in mind.

2. **Task Breakdown** — Each task should touch ~5 files max. Create checkpoints between major phases. Ensure the human reviews and approves the plan before any implementation begins.

3. **Parallel Opportunities** — Identify tasks that can be worked on in parallel (different modules, different files).

**Do NOT proceed to building until the plan is approved.**

---

## Phase 3: BUILD — `/build`

**Goal:** Implement tasks incrementally in thin, verifiable slices.

### Skills

| Skill | When to Use |
|-------|-------------|
| **incremental-implementation** | Any multi-file change, feature from a task breakdown, or refactoring |
| **test-driven-development** | Implementing logic, fixing bugs, or changing behavior (red-green-refactor) |
| **source-driven-development** | Framework-specific code — verify against official docs before implementing |
| **security-and-hardening** | Handling untrusted data, auth, storage, external integrations |

### Discipline

**Incremental Implementation Rules:**

1. **One Thing at a Time** — Each increment changes one logical thing. Don't mix concerns.
2. **Keep It Compilable** — After each increment, the project must build and existing tests must pass.
3. **Feature Flags** — Ship incomplete features behind flags. Deploy code, disable the flag.
4. **Scope Discipline** — Touch only what the task requires. Note improvements for later — don't fix them.
5. **Simplicity First** — Build the naive, obviously-correct version first. Optimize only after correctness is proven.

**Test-Driven Development (for logic):**

1. **Red** — Write a failing test that expresses the desired behavior.
2. **Green** — Write the minimum code to make the test pass.
3. **Refactor** — Improve the code without changing behavior.

**Source-Driven Development (for framework-specific code):**

1. **Detect** — Read `go.mod` / `package.json` / `requirements.txt` for exact versions.
2. **Fetch** — Get the relevant official documentation page (not the homepage).
3. **Implement** — Follow documented patterns, not memory.
4. **Cite** — Every framework-specific pattern gets a citation with full URL.

**Security (when handling untrusted data):**

- Input validation on all endpoints
- Authentication and authorization checks
- No secrets in code or version control
- Parameterized queries (no SQL injection)
- Secure headers (CSP, HSTS)

**Do NOT proceed to verification until the build is clean and tests pass.**

---

## Phase 4: VERIFY — `/test`

**Goal:** Prove the change works as specified before it enters the codebase.

### Skills

| Skill | When to Use |
|-------|-------------|
| **test-driven-development** | All logic must be proven by tests (red-green-refactor loop) |
| **debugging-and-error-recovery** | Tests fail, builds break, or behavior doesn't match expectations |
| **performance-optimization** | Performance requirements exist or regressions are suspected |

### Discipline

**Test Requirements:**

- [ ] All existing tests still pass
- [ ] New tests cover the change (unit, integration, e2e as appropriate)
- [ ] Bug fixes include a reproduction test that failed before the fix
- [ ] Test names describe behavior, not implementation
- [ ] No tests were skipped or disabled
- [ ] Coverage hasn't decreased (if tracked)

**Debugging Discipline:**

1. **Systematic** — Find the root cause, not symptoms. Use logs, debuggers, and reproducible test cases.
2. **Reproduce First** — Write a failing test that demonstrates the bug before fixing it.
3. **Minimal Fix** — Fix only what's broken. Don't refactor adjacent code.

**Performance (when relevant):**

- Profile before optimizing — identify the bottleneck
- Benchmark after — prove the optimization helped
- Document the trade-offs

**Do NOT proceed to review until verification passes.**

---

## Phase 5: REVIEW — `/review`

**Goal:** Multi-axis quality assessment before merging.

### Skills

| Skill | When to Use |
|-------|-------------|
| **code-review-and-quality** | Before merging any change — review code written by yourself, another agent, or a human |
| **code-simplification** | Code works but is harder to read, maintain, or extend than necessary |
| **doubt-driven-development** | High-stakes changes (production auth, security-sensitive logic, migrations) |

### Discipline

**Code Review Axes:**

1. **Correctness** — Does the code do what the spec says?
2. **Simplicity** — Could this be simpler? Are abstractions earning their complexity?
3. **Structure** — Does this change make the architecture better or worse?
4. **Tests** — Are they comprehensive, meaningful, and passing?
5. **Documentation** — Is the "why" documented, not just the "what"?
6. **Security** — No secrets, input validation, proper auth, no injection vectors.
7. **Performance** — No N+1 queries, appropriate caching, within budgets.
8. **Accessibility** — WCAG compliance for user-facing changes.

**Presumptive Blockers (propose simpler design; escalate to Required only when actively worse):**

- Refactor that relocates complexity instead of reducing it
- File past size boundary with no decomposition
- Feature logic added to a shared module
- Near-duplicate of an existing canonical helper
- Silent fallback that hides an unclear invariant

**Doubt-Driven Development (for high-stakes changes):**

- Subject every assumption to fresh-context adversarial review
- Stress-test before committing
- When correctness matters more than speed

**Do NOT proceed to shipping until review passes.**

---

## Phase 6: SHIP — `/ship`

**Goal:** Deploy safely with monitoring, rollback, and confidence.

### Skills

| Skill | When to Use |
|-------|-------------|
| **git-workflow-and-versioning** | Every code change — commits, branches, PRs, releases |
| **ci-cd-and-automation** | Setting up build and deployment pipelines |
| **shipping-and-launch** | Deploying to production, staged rollouts, rollback plans |
| **deprecation-and-migration** | Removing old systems, migrating users, expand/contract patterns |
| **documentation-and-adrs** | Documenting the why, not just the what |
| **observability-and-instrumentation** | Structured logs, metrics, traces, alerts |

### Discipline

**Git Workflow:**

- **Trunk-Based Development** — Keep `main` always deployable. Short-lived feature branches (1-3 days).
- **Atomic Commits** — Each commit does one logical thing.
- **Descriptive Messages** — `<type>: <short description>` with body explaining the *why*.
- **Commit Early, Commit Often** — Commits are save points.

**Shipping Checklist:**

- [ ] All tests pass (unit, integration, e2e)
- [ ] Build succeeds with no warnings
- [ ] Lint and type checking pass
- [ ] Code reviewed and approved
- [ ] No TODO comments that should be resolved before launch
- [ ] No debugging statements in production code
- [ ] Error handling covers expected failure modes
- [ ] No secrets in code or version control
- [ ] Dependency audit shows no critical/high vulnerabilities
- [ ] Feature flag configured (if applicable)
- [ ] Rollback plan documented
- [ ] Monitoring dashboards set up
- [ ] Health check endpoint exists and responds

**Staged Rollout:**

```
1. DEPLOY to staging → Full test suite in staging environment
2. DEPLOY to production (flag OFF) → Verify deployment succeeded
3. ENABLE for team (flag ON internally) → 24-hour monitoring window
4. CANARY rollout (5% of users) → Monitor error rates, latency, user behavior
5. GRADUAL increase (25% → 50% → 100%) → Same monitoring at each step
6. FULL rollout (flag ON for all) → Monitor for 1 week, clean up flag
```

**Rollback Plan (must exist before deploying):**

```markdown
## Rollback Plan for [Feature/Release]

### Trigger Conditions
- Error rate > 2x baseline
- P95 latency > [X]ms
- User reports of [specific issue]

### Rollback Steps
1. Disable feature flag (if applicable)
   OR
1. Deploy previous version: `git revert <commit> && git push`
2. Verify rollback: health check, error monitoring
3. Communicate: notify team of rollback

### Time to Rollback
- Feature flag: < 1 minute
- Redeploy previous version: < 5 minutes
```

**Do NOT ship until the shipping checklist is complete.**

---

## Quality Control

Bypassing checks is strictly forbidden unless an Exception ID is documented in `CONSTRAINTS.md`.

- **No silencing linters** (`//nolint`).
- **No suppressing errors** (no blank identifiers `_ = ...` for error handling).
- **No skipping tests** (`t.Skip()`) to force a green build.
- **No unverified framework-specific patterns** — always check official docs.

---

## Audit Requirement

Every PR/CR must include a **Quality Compliance Report**:

- Build Status (`make build`)
- Test Status (`make test`)
- Lint Status (`make lint`)
- Race Safety (`go test -race ./...`)
- Coverage Delta (e.g., +2.5%)
- Constraint Check (No floor violations)
- Review Status (Approved by code-review-and-quality)
- Shipping Checklist (All items green)

---

## Quick Reference: Skill Invocation

When the work enters a phase, invoke the relevant skill by name:

```
# Phase 1: Define
/skill idea-refine "Help me refine this feature idea"
/skill spec-driven-development "Create a spec for the auth module"
/skill interview-me "Build me a notification system"

# Phase 2: Plan
/skill planning-and-task-breakdown "Break this spec into tasks"
/skill constraint-driven-development "Define our quality constraints"

# Phase 3: Build
/skill incremental-implementation "Implement Task 3 from the plan"
/skill test-driven-development "Write tests for the validation logic"
/skill source-driven-development "Verify this approach against the docs"
/skill security-and-hardening "Review this auth flow for vulnerabilities"

# Phase 4: Verify
/skill debugging-and-error-recovery "This test is failing — help me find the root cause"
/skill performance-optimization "Profile the API endpoint for bottlenecks"

# Phase 5: Review
/skill code-review-and-quality "Review this PR"
/skill code-simplification "Simplify this function for clarity"
/skill doubt-driven-development "Stress-test this production auth change"

# Phase 6: Ship
/skill git-workflow-and-versioning "Help me cut a release"
/skill shipping-and-launch "I need to deploy to production — what's the checklist?"
/skill ci-cd-and-automation "Set up CI for this pipeline"
```

---

## See Also

- `CONSTRAINTS.md` — The absolute source of truth for quality thresholds
- `references/definition-of-done.md` — The standing bar every task clears before it counts as done
- `references/security-checklist.md` — Security pre-launch checks
- `references/performance-checklist.md` — Performance verification before launch
- `references/accessibility-checklist.md` — Accessibility verification before launch
