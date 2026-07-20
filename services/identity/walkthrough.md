# Walkthrough: Identity & Access Management (IAM) Service

This document summarizes the complete implementation of the central **PRAHARI Identity & Access Management (IAM) Service** under `services/identity/`.

---

## 1. Directory Structure

- **`cmd/`**: Entry point invoking bootstrap app runners.
- **`migrations/`**: SQL migration scripts defining tables for users, organizations, sessions, and audit logs.
- **`internal/domain/`**: DDD Aggregate entities for Users, Organizations, Roles, Permissions, API Keys, Service Accounts, and Audit records.
- **`internal/application/`**: Orchestration logic executing logins (`auth_service`) and signups (`user_service`).
- **`internal/infrastructure/`**: Postgres stores, Redis blacklists cache, and Kafka publishers.
- **`internal/interfaces/`**: REST handlers and gRPC ValidateToken servers.

---

## 2. API Endpoints Catalog

- **`POST /register`**: Instantiates new user profiles, encrypting passwords.
- **`POST /login`**: Validates user credentials, returning Bearer JWT access tokens.
- **`ValidateToken (gRPC)`**: Invoked by gateway/microservices to inspect token scopes.
