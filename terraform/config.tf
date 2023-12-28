terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region  = var.region
  profile = var.credentials_profile
  # default_tags {
  #   tags = {
  #     app = var.app_name
  #   }
  # }
}

