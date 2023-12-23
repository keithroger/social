terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~>2.20.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "docker" {}

provider "aws" {
  region  = var.region
  profile = var.credentials_profile
  default_tags {
    tags = {
      app = var.app_name
    }
  }
}

