resource "aws_rds_cluster" "postgres" {
  cluster_identifier          = "aurora-cluster"
  availability_zones          = var.availability_zones
  engine = "aurora-postgresql"
  engine_mode = "provisioned"
  database_name               = "postgres"
  master_username             = "postgres"
  manage_master_user_password = true
  backup_retention_period     = 7
  preferred_backup_window = "09:00-11:00"

  serverlessv2_scaling_configuration {
    max_capacity = 2.0
    min_capacity = 0.5
  }
}

resource "aws_rds_cluster_instance" "this" {
  cluster_identifier = aws_rds_cluster.postgres.id
  instance_class     = "db.serverless"
  engine             = aws_rds_cluster.postgres.engine
  engine_version     = aws_rds_cluster.postgres.engine_version
}