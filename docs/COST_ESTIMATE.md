# PRAHARI AWS Monthly Cost Breakdown (<$100 Credit Budget Target)

This document provides a detailed itemized cost estimation for running the **PRAHARI AWS Hackathon Infrastructure**.

---

## 1. Itemized Cost Breakdown

| Component | AWS Resource Type | Configuration Details | Estimated Cost (USD / Month) |
|---|---|---|:---:|
| **Amazon ECS Fargate** | Compute Container | 2 Tasks (0.5 vCPU / 1GB RAM each) | **$35.00 – $45.00** |
| **Amazon RDS PostgreSQL** | Relational Database | `db.t4g.micro` (20GB Storage, Single-AZ) | **$15.00 – $20.00** |
| **Amazon ElastiCache Redis** | Distributed Cache | `cache.t4g.micro` (1 Node) | **$12.00 – $15.00** |
| **Application Load Balancer**| Networking | 1 ALB + Target Groups | **$18.00 – $22.00** |
| **Amazon CloudFront** | CDN Distribution | Global Edge Network (Static Frontend) | **$2.00 – $5.00** |
| **Amazon S3** | Static Website Bucket | 5GB Storage + GET/PUT requests | **$1.00 – $3.00** |
| **AWS Secrets Manager** | Key Governance | 2 Encrypted Secret Entries | **$0.80 – $1.50** |
| **Amazon ECR** | Container Registry | 2 Image Repositories (< 5GB) | **$0.50 – $1.00** |
| **Amazon CloudWatch** | Logs & Metrics | 7-day retention + basic metrics | **$3.00 – $5.00** |
| **TOTAL ESTIMATED COST** | | | **~$88.30 – $117.50 / Month** |

---

## 2. Hackathon Cost Optimization Strategies Applied

1. **Single NAT Gateway**: Configured a single NAT Gateway instead of Multi-AZ NAT Gateways to save ~$64/month.
2. **Single-AZ RDS Instance**: Used `db.t4g.micro` in a single AZ to remain well within free-tier / micro instance bounds.
3. **S3 Static Website + CloudFront**: Hosted the React web dashboard directly on S3 + CloudFront CDN instead of running containerized NGINX web servers on ECS, saving ~$30/month.
4. **Fargate Resource Right-Sizing**: Allocated 0.5 vCPU and 1GB RAM per ECS task to prevent compute over-provisioning.
