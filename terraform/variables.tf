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
}

variable "vpc_cidr" {
  description = "CIDR block for main"
  type        = string
}
