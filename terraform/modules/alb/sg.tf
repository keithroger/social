# NLB security group
resource "aws_security_group" "this" {
  name        = "${var.name}-sg"
  description = "${var.name} security group"
  vpc_id      = var.vpc
  tags = {
    name = "${var.name}-sg"
  }
}

# resource "aws_vpc_security_group_ingress_rule" "all_traffic" {
#   security_group_id = aws_security_group.this.id
#   ip_protocol       = "-1"
#   cidr_ipv4         = "0.0.0.0/0"
# }

resource "aws_vpc_security_group_ingress_rule" "vpc_link_traffic" {
  security_group_id = aws_security_group.this.id
  ip_protocol       = "-1"
  referenced_security_group_id = var.vpc_link_sg_id
}

resource "aws_vpc_security_group_egress_rule" "all_traffic" {
  security_group_id = aws_security_group.this.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}
