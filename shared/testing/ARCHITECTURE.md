# Architecture: Testing SDK Subsystem

This document details the interface mocking paradigms, data fixture setups, and HTTP executor pipelines governing the `shared/testing` package.

---

## 1. Mocking Paradigm

We employ **Hook-based Mocking**. Clients are represented as structs containing function fields. This eliminates code-generation dependencies and lets developers override behavior dynamically inside individual test cases:

```
  +------------------+
  |   MockS3Client   |
  |                  |
  |  UploadFunc:     | ----> Custom closure defined inside a Test function
  |  DownloadFunc:   |
  |  DeleteFunc:     |
  +------------------+
```

---

## 2. Test Fixtures

- **Consistent Claims**: Simplifies JWT authentication verification.
- **Consistent Payloads**: Asserts JSON structural compatibility.
