# Contributing to PRAHARI Shared Platform

Thank you for contributing to the shared platform library. Since every microservice in the PRAHARI ecosystem imports these packages, changes must be treated with high caution.

---

## Egotist Rules of Contribution

1. **Maintain Backward Compatibility**: Avoid breaking API exports in existing packages. If a model changes, deprecate the old one and introduce a versioned interface wrapper.
2. **Zero Placeholders**: Do not commit TODO blocks or empty stubs. All written functions must be fully realized.
3. **High Unit Test Mandate**: PRs updating shared libraries will not be merged without achieving **>=95% unit test coverage** and passing fuzzing/benchmark regressions.
4. **Lint Compliances**: Run `make lint` prior to checkin. Ensure no errors are returned from `golangci-lint`.
