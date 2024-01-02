# ECS security group
resource "aws_security_group" "this" {
  name        = "${var.name}-sg"
  description = "${var.name} ecs security group"
  vpc_id      = var.vpc
  tags = {
    name = "${var.name}-sg"
  }
}

resource "aws_vpc_security_group_ingress_rule" "all_ingress_trafic" {
  security_group_id = aws_security_group.this.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

# resource "aws_vpc_security_group_ingress_rule" "load_balancer" {
#   security_group_id = aws_security_group.this.id
#   ip_protocol       = "-1"
#   referenced_security_group_id = var.load_balancer_sg_id
# }

resource "aws_vpc_security_group_egress_rule" "all_egress_traffic" {
  security_group_id = aws_security_group.this.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

# resource "aws_vpc_security_group_egress_rule" "rds" {
#   security_group_id = aws_security_group.this.id
#   ip_protocol       = "-1"
#   referenced_security_group_id = var.rds_sg_id
# }
