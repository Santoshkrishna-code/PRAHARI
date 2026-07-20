variable "aws_region" {
  type        = string
  description = "AWS deployment region target"
  default     = "us-east-1"
}

variable "environment" {
  type        = string
  description = "Deployment workspace target environment"
  default     = "production"
}

variable "cluster_name" {
  type        = string
  description = "EKS target cluster label name"
  default     = "prahari-production-cluster"
}

variable "vpc_cidr" {
  type        = string
  description = "VPC network IP scope CIDR block"
  default     = "10.0.0.0/16"
}
