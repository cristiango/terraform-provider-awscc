// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_ipam_scope", iPAMScopeResourceType)
}

// iPAMScopeResourceType returns the Terraform awscc_ec2_ipam_scope resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::IPAMScope resource type.
func iPAMScopeResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IPAM scope.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IPAM scope.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"ipam_arn": {
			// Property: IpamArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ipam_id": {
			// Property: IpamId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Id of the IPAM this scope is a part of.",
			//   "type": "string"
			// }
			Description: "The Id of the IPAM this scope is a part of.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
			// IpamId is a write-only property.
		},
		"ipam_scope_id": {
			// Property: IpamScopeId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the IPAM scope.",
			//   "type": "string"
			// }
			Description: "Id of the IPAM scope.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ipam_scope_type": {
			// Property: IpamScopeType
			// CloudFormation resource type schema:
			// {
			//   "description": "Determines whether this scope contains publicly routable space or space for a private network",
			//   "enum": [
			//     "public",
			//     "private"
			//   ],
			//   "type": "string"
			// }
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"is_default": {
			// Property: IsDefault
			// CloudFormation resource type schema:
			// {
			//   "description": "Is this one of the default scopes created with the IPAM.",
			//   "type": "boolean"
			// }
			Description: "Is this one of the default scopes created with the IPAM.",
			Type:        types.BoolType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"pool_count": {
			// Property: PoolCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of pools that currently exist in this scope.",
			//   "type": "integer"
			// }
			Description: "The number of pools that currently exist in this scope.",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema of AWS::EC2::IPAMScope Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMScope").WithTerraformTypeName("awscc_ec2_ipam_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"description":     "Description",
		"ipam_arn":        "IpamArn",
		"ipam_id":         "IpamId",
		"ipam_scope_id":   "IpamScopeId",
		"ipam_scope_type": "IpamScopeType",
		"is_default":      "IsDefault",
		"key":             "Key",
		"pool_count":      "PoolCount",
		"tags":            "Tags",
		"value":           "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/IpamId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
