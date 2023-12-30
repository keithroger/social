# Security group for endpoints
resource "aws_security_group" "vpc_endpoint_sg" {
  name        = "${var.name}-vpc-endpoint-sg"
  description = "${var.name} vpc endpoint security group"
  vpc_id      = aws_vpc.vpc.id
  tags = {
    name = "${var.name}-vpce-sg"
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

# S3 gateway endpoint
resource "aws_vpc_endpoint" "s3" {
  vpc_id       = aws_vpc.vpc.id
  service_name = "com.amazonaws.${var.region}.s3"
  # route_table_ids   = 
  # security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  # subnet_ids = aws_subnet.private.*.id
  route_table_ids   = [aws_route_table.private_route_table.id]
  auto_accept       = true
  vpc_endpoint_type = "Gateway"
  tags = {
    Name = "${var.name}-vpce"
  }
}

# Interface endpoints
resource "aws_vpc_endpoint" "cloud-map-discovery-2" {
  for_each = toset([
    # "com.amazonaws.${var.region}.execute-api",
    "com.amazonaws.${var.region}.logs",
    "com.amazonaws.${var.region}.ecr.api",
    "com.amazonaws.${var.region}.ecr.dkr",
    "com.amazonaws.${var.region}.ecs",
    "com.amazonaws.${var.region}.ecs-agent",
    "com.amazonaws.${var.region}.ecs-telemetry",
    "com.amazonaws.${var.region}.secretsmanager"
    # "com.amazonaws.${var.region}.servicediscovery",
    # "com.amazonaws.${var.region}.servicediscovery-fips",
    # "com.amazonaws.${var.region}.data-servicediscovery",
    # "com.amazonaws.${var.region}.data-servicediscovery-fips",
  ])
  vpc_id              = aws_vpc.vpc.id
  service_name        = each.value
  security_group_ids  = [aws_security_group.vpc_endpoint_sg.id]
  subnet_ids          = aws_subnet.private.*.id
  private_dns_enabled = true
  auto_accept         = true
  vpc_endpoint_type   = "Interface"
  tags = {
    Name = "${var.name}-vpce"
  }
}
