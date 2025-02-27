// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_connect_contact_flow", contactFlowDataSourceType)
}

// contactFlowDataSourceType returns the Terraform awscc_connect_contact_flow data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Connect::ContactFlow resource type.
func contactFlowDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"contact_flow_arn": {
			// Property: ContactFlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the contact flow (ARN).",
			//   "maxLength": 500,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the contact flow (ARN).",
			Type:        types.StringType,
			Computed:    true,
		},
		"content": {
			// Property: Content
			// CloudFormation resource type schema:
			// {
			//   "description": "The content of the contact flow in JSON format.",
			//   "maxLength": 256000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The content of the contact flow in JSON format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the contact flow.",
			//   "maxLength": 500,
			//   "type": "string"
			// }
			Description: "The description of the contact flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance (ARN).",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance (ARN).",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the contact flow.",
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the contact flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the contact flow.",
			//   "enum": [
			//     "ACTIVE",
			//     "ARCHIVED"
			//   ],
			//   "type": "string"
			// }
			Description: "The state of the contact flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the contact flow.",
			//   "enum": [
			//     "CONTACT_FLOW",
			//     "CUSTOMER_QUEUE",
			//     "CUSTOMER_HOLD",
			//     "CUSTOMER_WHISPER",
			//     "AGENT_HOLD",
			//     "AGENT_WHISPER",
			//     "OUTBOUND_WHISPER",
			//     "AGENT_TRANSFER",
			//     "QUEUE_TRANSFER"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of the contact flow.",
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
		Description: "Data Source schema for AWS::Connect::ContactFlow",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::ContactFlow").WithTerraformTypeName("awscc_connect_contact_flow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_flow_arn": "ContactFlowArn",
		"content":          "Content",
		"description":      "Description",
		"instance_arn":     "InstanceArn",
		"key":              "Key",
		"name":             "Name",
		"state":            "State",
		"tags":             "Tags",
		"type":             "Type",
		"value":            "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
