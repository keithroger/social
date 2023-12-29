output "postgres_secret" {
    value = aws_rds_cluster.postgres.master_user_secret
}
