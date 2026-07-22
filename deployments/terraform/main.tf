# PRAHARI AWS Hackathon Infrastructure as Code (Terraform)
# Budget Target: < $100 AWS Credits / month

terraform {
  required_version = ">= 1.5.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# ------------------------------------------------------------------------------
# 1. Networking (VPC, Subnets, NAT Gateway, Routes)
# ------------------------------------------------------------------------------
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "prahari-${var.environment}-vpc"
    Environment = var.environment
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "prahari-${var.environment}-igw"
  }
}

resource "aws_subnet" "public_a" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${var.aws_region}a"
  map_public_ip_on_launch = true

  tags = {
    Name = "prahari-public-subnet-a"
  }
}

resource "aws_subnet" "public_b" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "${var.aws_region}b"
  map_public_ip_on_launch = true

  tags = {
    Name = "prahari-public-subnet-b"
  }
}

resource "aws_subnet" "private_a" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.10.0/24"
  availability_zone = "${var.aws_region}a"

  tags = {
    Name = "prahari-private-subnet-a"
  }
}

resource "aws_subnet" "private_b" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.20.0/24"
  availability_zone = "${var.aws_region}b"

  tags = {
    Name = "prahari-private-subnet-b"
  }
}

# Single NAT Gateway for private subnets to optimize hackathon budget
resource "aws_eip" "nat_eip" {
  domain = "vpc"
}

resource "aws_nat_gateway" "nat" {
  allocation_id = aws_eip.nat_eip.id
  subnet_id     = aws_subnet.public_a.id

  tags = {
    Name = "prahari-nat-gateway"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "prahari-public-rt"
  }
}

resource "aws_route_table_association" "public_a" {
  subnet_id      = aws_subnet.public_a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_b" {
  subnet_id      = aws_subnet.public_b.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat.id
  }

  tags = {
    Name = "prahari-private-rt"
  }
}

resource "aws_route_table_association" "private_a" {
  subnet_id      = aws_subnet.private_a.id
  route_table_id = aws_route_table.private.id
}

resource "aws_route_table_association" "private_b" {
  subnet_id      = aws_subnet.private_b.id
  route_table_id = aws_route_table.private.id
}

# ------------------------------------------------------------------------------
# 2. Security Groups
# ------------------------------------------------------------------------------
resource "aws_security_group" "alb_sg" {
  name        = "prahari-alb-sg-${var.environment}"
  description = "Allow inbound HTTP/HTTPS traffic to ALB"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "ecs_sg" {
  name        = "prahari-ecs-sg-${var.environment}"
  description = "Allow inbound traffic from ALB to ECS tasks"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 8000
    to_port         = 8000
    protocol        = "tcp"
    security_groups = [aws_security_group.alb_sg.id]
  }

  ingress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.alb_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "db_sg" {
  name        = "prahari-db-sg-${var.environment}"
  description = "Allow PostgreSQL inbound traffic from ECS"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.ecs_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "redis_sg" {
  name        = "prahari-redis-sg-${var.environment}"
  description = "Allow Redis inbound traffic from ECS"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.ecs_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# ------------------------------------------------------------------------------
# 3. Amazon ECR Repositories
# ------------------------------------------------------------------------------
resource "aws_ecr_repository" "ai_platform" {
  name                 = "prahari-ai-platform"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "gateway_service" {
  name                 = "prahari-gateway-service"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

# ------------------------------------------------------------------------------
# 4. Amazon RDS PostgreSQL Database (Single-AZ, db.t4g.micro for < $20/mo)
# ------------------------------------------------------------------------------
resource "aws_db_subnet_group" "main" {
  name       = "prahari-db-subnets-${var.environment}"
  subnet_ids = [aws_subnet.private_a.id, aws_subnet.private_b.id]
}

resource "aws_db_instance" "postgres" {
  identifier             = "prahari-postgres-${var.environment}"
  allocated_storage      = 20
  storage_type           = "gp3"
  engine                 = "postgres"
  engine_version         = "15.7"
  instance_class         = "db.t4g.micro"
  db_name                = "prahari_db"
  username               = "postgres"
  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.db_sg.id]
  skip_final_snapshot    = true
}

# ------------------------------------------------------------------------------
# 5. Amazon ElastiCache Redis (Single Node, cache.t4g.micro for < $15/mo)
# ------------------------------------------------------------------------------
resource "aws_elasticache_subnet_group" "main" {
  name       = "prahari-redis-subnets-${var.environment}"
  subnet_ids = [aws_subnet.private_a.id, aws_subnet.private_b.id]
}

resource "aws_elasticache_cluster" "redis" {
  cluster_id           = "prahari-redis-${var.environment}"
  engine               = "redis"
  node_type            = "cache.t4g.micro"
  num_cache_nodes      = 1
  parameter_group_name = "default.redis7"
  port                 = 6379
  subnet_group_name    = aws_elasticache_subnet_group.main.name
  security_group_ids   = [aws_security_group.redis_sg.id]
}

# ------------------------------------------------------------------------------
# 6. AWS Secrets Manager
# ------------------------------------------------------------------------------
resource "aws_secretsmanager_secret" "app_secrets" {
  name                    = "prahari-${var.environment}-secrets"
  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "app_secrets_val" {
  secret_id     = aws_secretsmanager_secret.app_secrets.id
  secret_string = jsonencode({
    JWT_SECRET_KEY        = "prahari-super-secret-jwt-key-2026"
    POSTGRES_PASSWORD     = var.db_password
    OPENAI_API_KEY        = "sk-placeholder-key"
    GEMINI_API_KEY        = "gemini-placeholder-key"
  })
}

# ------------------------------------------------------------------------------
# 7. Frontend Static Web Hosting (S3 Website)
# ------------------------------------------------------------------------------
resource "aws_s3_bucket" "frontend" {
  bucket        = "prahari-hackathon-frontend-${var.environment}"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "frontend" {
  bucket                  = aws_s3_bucket.frontend.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_website_configuration" "frontend" {
  bucket = aws_s3_bucket.frontend.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

resource "aws_s3_bucket_policy" "frontend_policy" {
  bucket     = aws_s3_bucket.frontend.id
  depends_on = [aws_s3_bucket_public_access_block.frontend]

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "PublicReadGetObject"
        Effect    = "Allow"
        Principal = "*"
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.frontend.arn}/*"
      }
    ]
  })
}

# resource "aws_cloudfront_distribution" "frontend" {
#   origin {
#     domain_name = aws_s3_bucket_website_configuration.frontend.website_endpoint
#     origin_id   = "S3-Frontend"
# 
#     custom_origin_config {
#       http_port              = 80
#       https_port             = 443
#       origin_protocol_policy = "http-only"
#       origin_ssl_protocols   = ["TLSv1.2"]
#     }
#   }
# 
#   enabled             = true
#   is_ipv6_enabled     = true
#   default_root_object = "index.html"
# 
#   default_cache_behavior {
#     allowed_methods  = ["GET", "HEAD"]
#     cached_methods   = ["GET", "HEAD"]
#     target_origin_id = "S3-Frontend"
# 
#     forwarded_values {
#       query_string = false
#       cookies {
#         forward = "none"
#       }
#     }
# 
#     viewer_protocol_policy = "redirect-to-https"
#     min_ttl                = 0
#     default_ttl            = 3600
#     max_ttl                = 86400
#   }
# 
#   restrictions {
#     geo_restriction {
#       restriction_type = "none"
#     }
#   }
# 
#   viewer_certificate {
#     cloudfront_default_certificate = true
#   }
# }

# ------------------------------------------------------------------------------
# 8. IAM Execution & Task Roles
# ------------------------------------------------------------------------------
resource "aws_iam_role" "ecs_execution_role" {
  name = "prahari-ecs-execution-role-${var.environment}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_execution_policy" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role" "ecs_task_role" {
  name = "prahari-ecs-task-role-${var.environment}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      }
    }]
  })
}

# ------------------------------------------------------------------------------
# 9. CloudWatch Log Group
# ------------------------------------------------------------------------------
resource "aws_cloudwatch_log_group" "ecs" {
  name              = "/ecs/prahari-${var.environment}"
  retention_in_days = 7
}

# ------------------------------------------------------------------------------
# 10. Application Load Balancer & Target Groups
# ------------------------------------------------------------------------------
resource "aws_lb" "main" {
  name               = "prahari-alb-${var.environment}"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = [aws_subnet.public_a.id, aws_subnet.public_b.id]
}

resource "aws_lb_target_group" "ai_platform" {
  name        = "prahari-ai-tg-${var.environment}"
  port        = 8000
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"

  health_check {
    path                = "/health"
    interval            = 15
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 3
  }
}

resource "aws_lb_target_group" "gateway" {
  name        = "prahari-gw-tg-${var.environment}"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"

  health_check {
    path                = "/health"
    interval            = 15
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 3
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.main.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ai_platform.arn
  }
}

# ------------------------------------------------------------------------------
# 11. Amazon ECS Fargate Cluster & Services
# ------------------------------------------------------------------------------
resource "aws_ecs_cluster" "main" {
  name = "prahari-${var.environment}-cluster"
}

resource "aws_ecs_task_definition" "ai_platform" {
  family                   = "prahari-ai-platform"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "512"
  memory                   = "1024"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([{
    name      = "ai-platform"
    image     = "${aws_ecr_repository.ai_platform.repository_url}:latest"
    essential = true
    portMappings = [{
      containerPort = 8000
      hostPort      = 8000
    }]
    environment = [
      { name = "APP_ENV", value = "production" },
      { name = "DATABASE_URL", value = "postgresql+asyncpg://postgres:${var.db_password}@${aws_db_instance.postgres.address}:${aws_db_instance.postgres.port}/prahari_db" },
      { name = "REDIS_HOST", value = aws_elasticache_cluster.redis.cache_nodes[0].address },
      { name = "REDIS_PORT", value = "6379" }
    ]
    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = aws_cloudwatch_log_group.ecs.name
        "awslogs-region"        = var.aws_region
        "awslogs-stream-prefix" = "ai-platform"
      }
    }
  }])
}

resource "aws_ecs_service" "ai_platform" {
  name            = "prahari-ai-platform-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.ai_platform.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public_a.id, aws_subnet.public_b.id]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.ai_platform.arn
    container_name   = "ai-platform"
    container_port   = 8000
  }
}
