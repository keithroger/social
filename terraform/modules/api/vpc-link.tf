# VPC link to connect to private subnet
resource "aws_apigatewayv2_vpc_link" "vpc_link" {
  name               = "${var.name}-vpc-link"
  security_group_ids = [aws_security_group.vpc_link.id]
  subnet_ids         = var.subnets

  tags = {
    name = "${var.name}-vpc-link"
  }
}

# VPC link security group
resource "aws_security_group" "vpc_link" {
  name        = "${var.name}-vpc-link-sg"
  description = "${var.name} vpc link security group"
  vpc_id      = var.vpc
  tags = {
    name = "${var.name}-vpc-link-sg"
  }
}

resource "aws_vpc_security_group_ingress_rule" "vpc_link" {
  security_group_id = aws_security_group.vpc_link.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "vpc_link" {
  security_group_id = aws_security_group.vpc_link.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}
