# ==============================================================================
# PRAHARI Terraform Infrastructure Module Template
# ==============================================================================

terraform {
  required_version = ">= 1.5.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Example encrypted S3 Bucket showing platform standards:
# 1. KMS Encryption required
# 2. Public access blocked
# 3. Tagging rules applied
resource "aws_s3_bucket" "this" {
  bucket        = "prahari-${var.environment}-${var.region_short}-${var.bucket_name}"
  force_destroy = var.force_destroy

  tags = merge(
    var.tags,
    {
      Name        = "prahari-${var.environment}-${var.region_short}-${var.bucket_name}"
      Environment = var.environment
      ManagedBy   = "Terraform"
    }
  )
}

# Enforce encryption at rest using KMS
resource "aws_s3_bucket_server_side_encryption_configuration" "this" {
  bucket = aws_s3_bucket.this.id

  rule {
    apply_server_side_encryption_by_default {
      kms_master_key_id = var.kms_key_arn
      sse_algorithm     = "aws:kms"
    }
  }
}

# Block all public ingress/egress to bucket (Zero Trust compliance)
resource "aws_s3_bucket_public_access_block" "this" {
  bucket = aws_s3_bucket.this.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
