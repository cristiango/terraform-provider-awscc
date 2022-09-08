terraform {
  required_providers {
    axiom = {
      version = "~> 1.0.0"
      source  = "enablon/axiom"
    }
  }
}

resource "axiom_sample_resource_type" "test" {
  name = "testing"
}