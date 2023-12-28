module "vpc" {
  source = "./modules/vpc"

  name               = "${var.app_name}-vpc"
  region             = var.region
  availability_zones = var.availability_zones
  vpc_cidr           = var.vpc_cidr
}

module "nlb" {
  source = "./modules/nlb"

  name            = "${var.app_name}-nlb"
  vpc             = module.vpc.vpc_id
  subnets         = module.vpc.private_subnets
  certificate_arn = var.certificate_arn
}

module "ecr" {
  source = "./modules/ecr"

  name = "${var.app_name}-ecr"
}

module "ecs" {
  source = "./modules/ecs"

  name    = "${var.app_name}-ecs"
  vpc     = module.vpc.vpc_id
  subnets = module.vpc.private_subnets
  repository_url = "${module.ecr.repository_url}:latest"
  target_group_arn    = module.nlb.target_group_arn
  load_balancer_sg_id = module.nlb.load_balancer_sg_id
}

module "api" {
  source = "./modules/api"

  name                    = "${var.app_name}-api"
  vpc                     = module.vpc.vpc_id
  subnets                 = module.vpc.private_subnets
  listener_arn            = module.nlb.listener_arn
  domain                  = var.domain
  certificate_arn         = var.certificate_arn
  route_53_hosted_zone_id = var.route_53_hosted_zone_id
}
