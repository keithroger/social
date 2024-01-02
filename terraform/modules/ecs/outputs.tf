output "task_revision" {
  value = aws_ecs_task_definition.ecs_task_definition.revision
}

output "ecs_sg_id" {
  description = "Security group id to send traffic from ECS to RDS"
  value = aws_security_group.this.id
}