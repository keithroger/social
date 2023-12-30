output "postgres_secret_arn" {
    # Work around since master_user_secret.secret_arn does not work
    # https://github.com/hashicorp/terraform-provider-aws/issues/31519
    value = lookup(aws_rds_cluster.postgres.master_user_secret[0], "secret_arn")
}
