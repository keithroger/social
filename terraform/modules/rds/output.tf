output "postgres_read_write_endpoint" {
    value = aws_rds_cluster.postgres.endpoint
}

output "postgres_read_endpoint" {
    value = aws_rds_cluster.postgres.reader_endpoint
}

output "postgres_db" {
    value = aws_rds_cluster.postgres.database_name
}

output "postgres_user" {
    value = aws_rds_cluster.postgres.master_username
}

output "postgres_secret_arn" {
    # Work around since master_user_secret.secret_arn does not work
    # https://github.com/hashicorp/terraform-provider-aws/issues/31519
    value = lookup(aws_rds_cluster.postgres.master_user_secret[0], "secret_arn")
}

output "sg_id" {
    description = "Security group id to allow traffic traffic from ECS"
    value = aws_security_group.this.id
}