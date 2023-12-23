resource "aws_ecs_capacity_provider" "ecs_capacity_provider" {
  name = "test1"

  auto_scaling_group_provider {
    auto_scaling_group_arn = aws_autoscaling_group.ecs_asg.arn

    managed_scaling {
      maximum_scaling_step_size = 1000
      minimum_scaling_step_size = 1
      status                    = "ENABLED"
      target_capacity           = 3
    }
  }
}

resource "aws_ecs_cluster_capacity_providers" "example" {
  cluster_name = aws_ecs_cluster.ecs_cluster.name

  capacity_providers = [aws_ecs_capacity_provider.ecs_capacity_provider.name]

  default_capacity_provider_strategy {
    base              = 1
    weight            = 100
    capacity_provider = aws_ecs_capacity_provider.ecs_capacity_provider.name
  }
}

resource "aws_cloudwatch_log_group" "cluster_log_group" {
  name = "${var.app_name}-cluster-log-group"
}

# define the ecs task definition for the service
resource "aws_ecs_task_definition" "ecs_task_definition" {
  family       = "my-ecs-task"
  network_mode = "awsvpc"
  execution_role_arn = aws_iam_role.ecs-execution-role.arn
  cpu                = 256
  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }
  container_definitions = jsonencode([
    {
      name      = "${var.app_name}-server"
      image     = "${aws_ecr_repository.ecr.repository_url}:latest"
      cpu       = 256
      memory    = 512
      essential = true
      portmappings = [
        {
          containerport = 8080
          hostport      = 8080
          protocol      = "tcp"
        }
      ]
    }
  ])
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "ecs-execution-role" {
  name               = "${var.app_name}-ecs-execution-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "policy" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "policy" {
  name        = "${var.app_name}-ecs-policy"
  description = "Allow application to pull from ECR and write to Cloud Watch"
  policy      = data.aws_iam_policy_document.policy.json
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.ecs-execution-role.name
  policy_arn = aws_iam_policy.policy.arn
}

# Cloud Map namesspace
resource "aws_service_discovery_private_dns_namespace" "private_dns" {
  name        = "${var.app_name}-dns"
  description = "service discovery endpoint"
  vpc         = aws_vpc.main.id
  tags = {
    name = "${var.app_name}-cloud-map-dns"
  }
}

resource "aws_service_discovery_service" "service_discovery" {
  name = "${var.app_name}-discovery"
  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.private_dns.id
    dns_records {
      ttl  = 10
      type = "A"
    }
    dns_records {
      ttl  = 10
      type = "SRV"
    }
    routing_policy = "MULTIVALUE"
  }
  health_check_custom_config {
    failure_threshold = 3
  }
  tags = {
    name = "${var.app_name}-cloud-map-service"
  }
}


resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.app_name}-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
  configuration {
    execute_command_configuration {
      logging    = "OVERRIDE"

      log_configuration {
        cloud_watch_log_group_name     = aws_cloudwatch_log_group.cluster_log_group.name
      }
    }
  }
}

resource "aws_ecs_service" "ecs_service" {

  name            = "${var.app_name}-service"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.ecs_task_definition.arn
  desired_count   = 3

  network_configuration {
    subnets          = aws_subnet.private.*.id
    security_groups  = [aws_security_group.security_group.id]
    # assign_public_ip = false
  }

  force_new_deployment = true

  ordered_placement_strategy {
    type  = "binpack"
    field = "cpu"
  }

  # placement_constraints {
  #   type = "distinctInstance"
  # }

  triggers = {
    redeployment = timestamp()
  }

  capacity_provider_strategy {
    capacity_provider = aws_ecs_capacity_provider.ecs_capacity_provider.name
    weight            = 100
  }
  service_registries {
    registry_arn = aws_service_discovery_service.service_discovery.arn
    port = 8080
  }

  lifecycle {
    ignore_changes = [desired_count]
  }

  depends_on = [aws_autoscaling_group.ecs_asg]
}
