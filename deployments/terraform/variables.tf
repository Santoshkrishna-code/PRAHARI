variable "aws_region" {
  description = "Target AWS deployment region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Target deployment environment"
  type        = string
  default     = "hackathon"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "container_port_ai" {
  description = "Port exposed by Python AI Platform container"
  type        = number
  default     = 8000
}

variable "container_port_gateway" {
  description = "Port exposed by Go Gateway container"
  type        = number
  default     = 8080
}

variable "db_password" {
  description = "Master password for RDS PostgreSQL instance"
  type        = string
  default     = "PrahariHackathonPass2026!"
  sensitive   = true
}
