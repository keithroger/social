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
