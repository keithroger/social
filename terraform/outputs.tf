output "ecr_repository" {
  value = module.ecr.repository_url
}

output "api_gateway_uri" {
  value = module.api.api_gateway_uri
}

output "ecs_task_revision" {
  value = module.ecs.task_revision
}

output "postgres_secret" {
  value = module.rds.postgres_secret
}