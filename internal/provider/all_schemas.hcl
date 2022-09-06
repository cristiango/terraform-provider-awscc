defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

# 1 CloudFormation resource types schemas are available for use with the Cloud Control API.

resource_schema "axiom_sample_resource_type" {
  cloudformation_type_name               = "Axiom::Sample::ResourceType"
  suppress_plural_data_source_generation = true
}