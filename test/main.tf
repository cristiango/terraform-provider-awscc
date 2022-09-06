terraform {
  required_providers {

    aws = {
      source = "hashicorp/aws"
    }

    awscc = {
      source = "hashicorp/awscc"
    }

    axiom = {
      source = "hashicorp/enablon-axiom"
    }
  }
}

provider "aws" {}
provider "awscc" {}
provider "axiom" {}

resource "awscc_sample_resource_type" "test" {
}