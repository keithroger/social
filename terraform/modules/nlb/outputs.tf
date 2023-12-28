output "target_group_arn" {
  description = "target group used to send requests to ECS"
  value       = aws_lb_target_group.ecs_tg.arn
}

output "listener_arn" {
  description = "ARN of listener for application load balancer"
  value       = aws_lb_listener.this.arn
}

output "load_balancer_sg_id" {
  description = "Load balancer security group id to allow ingress traffic to ECS"
  value = aws_security_group.this.id
}