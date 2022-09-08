locals {
  partition_labels = toset([
    "aws-cn",
    "aws-iso",
    "aws-iso-b",
    "aws-us-gov",
  ])
}

resource "github_issue_label" "partition" {
  for_each = local.partition_labels

  repository = "terraform-provider-axiom-awscc"
  name       = "partition/${each.value}"
  color      = "844fba" # color:terraform (main)
}
