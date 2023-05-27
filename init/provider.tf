terraform {
  required_version = ">= 1.3"
  backend "s3" {
    bucket = "terraform-example-tkoide"
    key    = "lambda-dynamodb/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

