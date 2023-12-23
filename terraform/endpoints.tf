# Security group for endpoints
resource "aws_security_group" "vpc_endpoint_sg" {
  name        = "${var.app_name}-vpc-endpoint-sg"
  description = "${var.app_name} vpc endpoint security group"
  vpc_id      = aws_vpc.main.id
  tags = {
    name = "${var.app_name}-vpc-endpoint-sg"
  }
}

resource "aws_vpc_security_group_ingress_rule" "vpc_endpoint_sg_ingress" {
  security_group_id = aws_security_group.vpc_endpoint_sg.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "vpc_endpoint_sg_egress" {
  security_group_id = aws_security_group.vpc_endpoint_sg.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

# S3
resource "aws_vpc_endpoint" "s3" {
  vpc_id       = aws_vpc.main.id
  service_name = "com.amazonaws.${var.region}.s3"
  # route_table_ids   = 
  # security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  # subnet_ids = aws_subnet.private.*.id
  route_table_ids   = [aws_route_table.private_route_table.id]
  auto_accept       = true
  vpc_endpoint_type = "Gateway"
  tags = {
    Name = "${var.app_name}-s3-endpoint"
  }
}

# API Gateway
resource "aws_vpc_endpoint" "api_endpoint" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.${var.region}.execute-api"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-api-endpoint"
  }
}

# CloudWatch
resource "aws_vpc_endpoint" "logs_endpoint" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.${var.region}.logs"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-cloudwatch-endpoint"
  }
}

# ECR
resource "aws_vpc_endpoint" "ecr_endpoint" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.${var.region}.ecr.api"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-ecr-endpoint"
  }
}

# ECR Docker
resource "aws_vpc_endpoint" "ecr_dkr_endpoint" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.${var.region}.ecr.dkr"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-ecr-dkr-endpoint"
  }
}


# ECS Agent
resource "aws_vpc_endpoint" "ecs-agent" {
  vpc_id              = aws_vpc.main.id
  service_name = "com.amazonaws.${var.region}.ecs-agent"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-ecs-agent-endpoint"
  }
}

# ECS Telemetry
resource "aws_vpc_endpoint" "ecs-telemetry" {
  vpc_id              = aws_vpc.main.id
  service_name = "com.amazonaws.${var.region}.ecs-telemetry"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-ecs-telemetry-endpoint"
  }
}

# ECS 
resource "aws_vpc_endpoint" "ecs" {
  vpc_id              = aws_vpc.main.id
  service_name = "com.amazonaws.${var.region}.ecs"
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.app_name}-ecs-endpoint"
  }
}
