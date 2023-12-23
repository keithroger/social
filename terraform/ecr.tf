resource "aws_ecr_repository" "ecr" {
  name                 = "${var.app_name}-ecs"
  image_tag_mutability = "MUTABLE"
  tags = {
    name = "${var.app_name}-ecr"
  }
  image_scanning_configuration {
    scan_on_push = true
  }
}