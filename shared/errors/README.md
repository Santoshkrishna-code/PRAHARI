# Core Error Framework

This package standardizes all domain, application, database, and infrastructure errors within the PRAHARI ecosystem.

---

## 1. Features

- **AppError struct**: Custom error schema holding ErrorCodes, human messages, underlying trace frames, and retry variables.
- **CaptureStackTrace**: Collects caller frames and serializes them in logging (via `%+v` formatting).
- **RFC 7807 problem details**: Unifies HTTP REST JSON error returns.
- **Status Mapping**: Automated HTTP and gRPC status code mapping.
- **Database/Network Translators**: Parses driver exceptions (e.g. Postgres duplicates) and maps them to AppErrors.

---

## 2. API Reference & Code Examples

### A. Constructing and wrapping errors
```go
import "prahari/shared/errors"

// Create new domain error
err := errors.New(errors.CodeInvalidArgument, "invalid format parameter")

// Wrap raw infra error
dbErr := queryDatabase()
if dbErr != nil {
    return errors.Wrap(dbErr, errors.CodeInternal, "failed to load details")
}
```

### B. Accessing details and causes
```go
// Check cause
rootCause := errors.Cause(err)

// Check code
code := errors.GetCode(err) // returns CodeUnknown if not AppError
```

### C. Translating SQL errors
```go
err := db.QueryRow(...).Scan(...)
if err != nil {
    return errors.TranslateDatabaseError(err, "User")
}
```
