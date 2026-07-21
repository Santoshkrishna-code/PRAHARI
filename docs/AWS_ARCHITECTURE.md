# PRAHARI AWS Hackathon Architectural Specification

This document details the cost-optimized AWS production architecture for the **PRAHARI Monorepo & AI Platform**.

---

## 1. System Topology Graph

```
                                              Internet Users / Judges
                                                         │
                                                         ▼
                                            Amazon CloudFront (CDN)
                                                         │
                        ┌────────────────────────────────┴────────────────────────────────┐
                        ▼                                                                 ▼
               Amazon S3 Static Bucket                                    Application Load Balancer (ALB)
                 (React Web Frontend)                                                     │
                                                                                          ▼
                                                                               Amazon ECS Fargate Cluster
                                                                             - Go API Gateway (8080)
                                                                             - Python AI Platform (8000)
                                                                                          │
               ┌─────────────────────────────────────────┼─────────────────────────────────────────┐
               ▼                                         ▼                                         ▼
   Amazon RDS PostgreSQL                    Amazon ElastiCache Redis                  AWS Secrets Manager & Logs
 (Transactional DB db.t4g.micro)               (Session Cache cache.t4g.micro)           (CloudWatch Container Insights)
```

---

## 2. Infrastructure Component Matrix

| Subsystem | AWS Service Used | Instance / Configuration Type | Hackathon Cost Strategy |
|---|---|---|---|
| **Web Frontend** | Amazon S3 + CloudFront | S3 Static Hosting + CDN | $2–5 / month |
| **API Load Balancing** | Application Load Balancer (ALB) | HTTP/HTTPS Routing to ECS Tasks | $18–22 / month |
| **Containers Orchestration** | Amazon ECS Fargate | 1 CPU / 1GB RAM per Task | $35–50 / month |
| **Container Registry** | Amazon ECR | `prahari-ai-platform` & `prahari-gateway-service` | $1–3 / month |
| **Relational Database** | Amazon RDS PostgreSQL 15 | `db.t4g.micro` (20GB Storage) | $15–20 / month |
| **In-Memory Cache** | Amazon ElastiCache Redis | `cache.t4g.micro` (1 node) | $12–15 / month |
| **Secrets Governance** | AWS Secrets Manager | Encrypted API keys & JWT secrets | $1–2 / month |
| **Observability & Logs** | Amazon CloudWatch | Log Groups + Container Insights | $5–10 / month |

---

## 3. High Availability & Multi-AZ Network Design

- **VPC Range**: `10.0.0.0/16`
- **Availability Zones**: `us-east-1a`, `us-east-1b`
- **Public Subnets**: `10.0.1.0/24`, `10.0.2.0/24` (Hosts ALB & Single NAT Gateway)
- **Private Subnets**: `10.0.10.0/24`, `10.0.20.0/24` (Hosts ECS Tasks, RDS, and ElastiCache)
