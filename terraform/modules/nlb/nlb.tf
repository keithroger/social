resource "aws_lb" "this" {
  name               = var.name
  internal           = false
  load_balancer_type = "network"
  security_groups = [aws_security_group.this.id]
  subnets = var.subnets
  # TODO add logging
  tags = {
    name = "${var.name}"
  }
}



resource "aws_lb_listener" "this" {
  load_balancer_arn = aws_lb.this.arn
  port              = "80"
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ecs_tg.arn
  }
}



resource "aws_lb_target_group" "ecs_tg" {
  name        = "${var.name}-tg"
  port        = 8080
  protocol    = "TCP"
  target_type = "ip"
  vpc_id      = var.vpc

  health_check {
    path = "/"
  }
}
