variable "name" {
  description = "Name for creating resources"
  type        = string
}

variable "vpc" {
  description = "VPC ID"
  type        = string
}

variable "subnets" {
  description = "List of subnets to deploy ECS"
  type        = list(string)
}

variable "listener_arn" {
  description = "ARN of ALB used to send request to ECS"
  type        = string
}

variable "domain" {
  description = "custom domain for API Gateway"
  type        = string
}

variable "certificate_arn" {
  description = "Arn of domain certificate"
  type        = string
}

variable "route_53_hosted_zone_id" {
  description = "route53 hosted zone id for custom domain"
  type        = string
}
