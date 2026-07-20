# Contributing Guidelines

Thank you for contributing to the PRAHARI Enterprise EHS Platform! Please follow these standards to ensure a clean codebase.

---

## 1. Branch Strategy

We follow a structured branch naming convention:
- Features: `feat/feature-name`
- Bug fixes: `fix/bug-name`
- DevOps / Chore updates: `chore/task-name`

---

## 2. Commit Message Convention

Commits must follow Conventional Commits standard rules:
```text
feat: add dynamic permit workflow checklists
fix: solve memory leaks inside RTSP camera streams
chore: update terraform variable subnets
```

---

## 3. Pull Request Standards

Before submitting a Pull Request, verify:
- Build compiles without syntax warnings or static check failures.
- Unit tests execute successfully (`GOWORK=off go test -v -race ./...`).
- Code has been formatted using `go fmt` and `prettier` formatting drivers.
