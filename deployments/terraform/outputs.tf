output "cloudfront_domain_name" {
  description = "CloudFront distribution domain name for React web frontend"
  value       = aws_cloudfront_distribution.frontend.domain_name
}

output "alb_dns_name" {
  description = "Application Load Balancer DNS Endpoint"
  value       = aws_lb.main.dns_name
}

output "ecr_ai_platform_repository" {
  description = "ECR Repository URL for AI Runtime Platform"
  value       = aws_ecr_repository.ai_platform.repository_url
}

output "ecr_gateway_repository" {
  description = "ECR Repository URL for Go API Gateway"
  value       = aws_ecr_repository.gateway_service.repository_url
}

output "rds_endpoint" {
  description = "Amazon RDS PostgreSQL Host Address"
  value       = aws_db_instance.postgres.endpoint
}

output "elasticache_endpoint" {
  description = "Amazon ElastiCache Redis Primary Endpoint"
  value       = aws_elasticache_cluster.redis.cache_nodes[0].address
}

output "s3_frontend_bucket" {
  description = "Amazon S3 Static Hosting Bucket Name"
  value       = aws_s3_bucket.frontend.id
}
