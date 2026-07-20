output "eks_cluster_endpoint" {
  value       = aws_eks_cluster.eks.endpoint
  description = "EKS API server master endpoint URL address"
}

output "eks_cluster_security_group_id" {
  value       = aws_eks_cluster.eks.vpc_config[0].cluster_security_group_id
  description = "Security Group ID applied to EKS Master nodes"
}

output "rds_endpoint" {
  value       = aws_db_instance.postgres.endpoint
  description = "PostgreSQL RDS connection string endpoint"
}
