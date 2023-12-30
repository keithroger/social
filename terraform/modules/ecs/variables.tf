variable "name" {
  description = "Name for creating resources"
  type        = string
}

variable "region" {
  description = "AWS region"
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

variable "repository_url" {
  description = "Url to pull container image needed to run tasks"
  type        = string
}

variable "target_group_arn" {
  description = "Target group for sending requests to ECS"
  type        = string
}

variable "load_balancer_sg_id" {
  description = "Load balancer security group id to allow ingress traffic"
  type        = string
}

variable "postgres_secret_arn" {
  description = "Secret arn for accessing postgress password"
  type = string
}