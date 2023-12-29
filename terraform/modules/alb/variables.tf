variable "name" {
  description = "Name for creating resources"
  type        = string
}

variable "vpc" {
  description = "VPC ID"
  type        = string
}

variable "subnets" {
  description = "List of subnets to deploy load balancer"
  type        = list(string)
}

variable "certificate_arn" {
  description = "Arn of domain certificate"
  type        = string
}

variable "vpc_link_sg_id" {
  description = "Security group id for recieving API Gateway VPC Link requests"
  type = string
}
