// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apigateway_documentation_version", documentationVersionDataSourceType)
}

// documentationVersionDataSourceType returns the Terraform awscc_apigateway_documentation_version data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ApiGateway::DocumentationVersion resource type.
func documentationVersionDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the API documentation snapshot.",
			//   "type": "string"
			// }
			Description: "The description of the API documentation snapshot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"documentation_version": {
			// Property: DocumentationVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version identifier of the API documentation snapshot.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The version identifier of the API documentation snapshot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the API.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The identifier of the API.",
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
		Description: "Data Source schema for AWS::ApiGateway::DocumentationVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationVersion").WithTerraformTypeName("awscc_apigateway_documentation_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":           "Description",
		"documentation_version": "DocumentationVersion",
		"rest_api_id":           "RestApiId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
