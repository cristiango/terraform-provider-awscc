// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package refactorspaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_refactorspaces_environment", environmentDataSourceType)
}

// environmentDataSourceType returns the Terraform awscc_refactorspaces_environment data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::RefactorSpaces::Environment resource type.
func environmentDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:(aws[a-zA-Z-]*)?:refactor-spaces:[a-zA-Z0-9\\-]+:\\w{12}:[a-zA-Z_0-9+=,.@\\-_/]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-_\\s\\.\\!\\*\\#\\@\\']+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"environment_identifier": {
			// Property: EnvironmentIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^env-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"network_fabric_type": {
			// Property: NetworkFabricType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "TRANSIT_GATEWAY"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Environment resource",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 21,
			//   "minLength": 21,
			//   "pattern": "^tgw-[-a-f0-9]{17}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::RefactorSpaces::Environment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Environment").WithTerraformTypeName("awscc_refactorspaces_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"description":            "Description",
		"environment_identifier": "EnvironmentIdentifier",
		"key":                    "Key",
		"name":                   "Name",
		"network_fabric_type":    "NetworkFabricType",
		"tags":                   "Tags",
		"transit_gateway_id":     "TransitGatewayId",
		"value":                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
