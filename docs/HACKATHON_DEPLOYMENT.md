# PRAHARI AWS Hackathon Deployment Guide

This guide provides step-by-step instructions for deploying the **PRAHARI Monorepo & AI Platform** to AWS under a **<$100 AWS Credit Budget**.

---

## 1. Prerequisites & Required CLI Tools

Ensure the following CLI tools are available:
- **AWS CLI** (`aws --version`)
- **Terraform** (`terraform --version` >= 1.5.0)
- **Docker** (`docker --version`)

---

## 2. Step-by-Step AWS Hackathon Deployment

### Step 1: Authenticate AWS CLI
```bash
aws configure
```

### Step 2: Provision Infrastructure with Terraform
```bash
cd deployments/terraform
terraform init
terraform plan
terraform apply -auto-approve
```

### Step 3: Build & Push Docker Container Images
```bash
# Log in to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com

# Build & Push Python AI Platform Image
docker build -t <ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/prahari-ai-platform:latest -f ai/platform/Dockerfile ai/platform
docker push <ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/prahari-ai-platform:latest

# Build & Push Go Gateway Image
docker build -t <ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/prahari-gateway-service:latest -f services/gateway-service/Dockerfile .
docker push <ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/prahari-gateway-service:latest
```

### Step 4: Deploy Frontend Assets to S3 & Invalidate CloudFront
```bash
# Build React Frontend
npm run build --workspaces

# Upload static assets to S3
aws s3 sync frontend/build s3://prahari-hackathon-frontend-hackathon --delete

# Invalidate CloudFront CDN Cache
aws cloudfront create-invalidation --distribution-id <DISTRIBUTION_ID> --paths "/*"
```

### Step 5: Verify Live Demo Execution
```bash
source ai/platform/.venv/bin/activate && python3 scripts/demo.py
```
