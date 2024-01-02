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

variable "rds_sg_id" {
  description = "RDS security group id to allow egress traffic"
  type = string
}

variable "postgres_read_write_endpoint" {
  description = "Host used to connect the postgres reader/writer client"
  type = string
}

variable "postgres_read_endpoint" {
  description = "Host used to connect the postgres reader client"
  type = string
}

variable "postgres_db" {
  description = "Name of the postgres database to connect to"
  type = string
}

variable "postgres_user" {
  description = "Name of the user used to access postgres"
  type = string
}

variable "postgres_secret_arn" {
  description = "Secret arn for accessing postgress password"
  type = string
}
