terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      configuration_aliases = [ aws, aws.virginia ]
    }
  }
}

data "aws_caller_identity" "main" {}
data "aws_region" "main" {}
