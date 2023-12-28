variable "app_name" {
  description = "Name of application"
  type        = string
}
variable "credentials_profile" {
  description = "AWS CLI credentials profile name"
  type        = string
}
variable "region" {
  description = "AWS region"
  type        = string
}

variable "availability_zones" {
  description = "List of availability zones"
  type        = list(string)
}

variable "vpc_cidr" {
  description = "CIDR block for main"
  type        = string
}

variable "domain" {
  description = "domain name to use as API gateway custom domain"
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
