# PRAHARI AWS Hackathon Deployment Verification Checklist

This document provides a pre-flight deployment verification checklist for the **PRAHARI AWS Hackathon Infrastructure**.

---

## 1. Pre-Flight Infrastructure Checklist

- [x] **AWS CLI Configuration**: Authenticated with valid IAM permissions.
- [x] **Terraform Syntax & Validation**: `terraform validate` passed with zero errors.
- [x] **Docker Engine Daemon**: Active and container images compiled (`prahari-ai-platform:test`).
- [x] **Hard Budget Cap**: Total estimated monthly AWS spend is **~$88.30 / month** (strictly under the $100 AWS Credit limit).
- [x] **Single NAT Gateway Policy**: Configured single NAT Gateway to eliminate Multi-AZ NAT charges.
- [x] **Single-AZ Database Instance**: Configured `db.t4g.micro` RDS PostgreSQL instance.
- [x] **S3 + CloudFront Static Hosting**: React frontend configured for static web hosting on S3 with CloudFront CDN distribution.

---

## 2. Post-Deployment Smoke Test Checklist

- [x] `GET /health` returns `200 OK`.
- [x] `POST /analytics/dashboard` returns Executive CEO Dashboard KPIs.
- [x] `POST /incident/investigate` processes 5 Whys Root Cause Analysis.
- [x] `POST /permit/create` issues 5x5 Risk Matrix Work Permit.
- [x] `POST /supervisor/chat` orchestrates Multi-Agent AI OS prompt dispatching.
- [x] `python3 scripts/demo.py` executes all 12 operational user journey steps cleanly.
