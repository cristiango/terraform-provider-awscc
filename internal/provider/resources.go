// Code generated by generators/schema/main.go; DO NOT EDIT.
//go:generate go run generators/resource/main.go -resource awscc_sample_resource_type -cfschema ../service/cloudformation/schemas/Axiom_Sample_ResourceType.json -package sample -- ../axiom/sample/resource_type_resource_gen.go ../axiom/sample/resource_type_resource_gen_test.go

package provider

import (
	_ "github.com/hashicorp/terraform-provider-awscc/internal/axiom/sample"
)
