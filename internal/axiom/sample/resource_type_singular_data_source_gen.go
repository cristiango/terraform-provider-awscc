// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sample

import (
	"context"

	. "github.com/eVisionSoftware/axiom/terraform-provider-axiom/internal/generic"
	"github.com/eVisionSoftware/axiom/terraform-provider-axiom/internal/registry"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sample_resource_type", resourceTypeDataSourceType)
}

// resourceTypeDataSourceType returns the Terraform awscc_sample_resource_type data source type.
// This Terraform data source type corresponds to the CloudFormation Axiom::Sample::ResourceType resource type.
func resourceTypeDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A meaninfull description",
			//   "type": "string"
			// }
			Description: "A meaninfull description",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The Identifier of the Resource Type",
			//   "type": "string"
			// }
			Description: "The Identifier of the Resource Type",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Resource Type",
			//   "type": "string"
			// }
			Description: "The name of the Resource Type",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for Axiom::Sample::ResourceType",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("Axiom::Sample::ResourceType").WithTerraformTypeName("awscc_sample_resource_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description": "Description",
		"id":          "Id",
		"name":        "Name",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
