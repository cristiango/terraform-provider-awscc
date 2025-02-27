// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

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
	registry.AddResourceTypeFactory("awscc_eks_addon", addonResourceType)
}

// addonResourceType returns the Terraform awscc_eks_addon resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EKS::Addon resource type.
func addonResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"addon_name": {
			// Property: AddonName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Addon",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of Addon",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"addon_version": {
			// Property: AddonVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Version of Addon",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Version of Addon",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the add-on",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the add-on",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Cluster",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of Cluster",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resolve_conflicts": {
			// Property: ResolveConflicts
			// CloudFormation resource type schema:
			// {
			//   "description": "Resolve parameter value conflicts",
			//   "enum": [
			//     "NONE",
			//     "OVERWRITE"
			//   ],
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Resolve parameter value conflicts",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
				validate.StringInSlice([]string{
					"NONE",
					"OVERWRITE",
				}),
			},
			// ResolveConflicts is a write-only property.
		},
		"service_account_role_arn": {
			// Property: ServiceAccountRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "IAM role to bind to the add-on's service account",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "IAM role to bind to the add-on's service account",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
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
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
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
		Description: "Resource Schema for AWS::EKS::Addon",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Addon").WithTerraformTypeName("awscc_eks_addon")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addon_name":               "AddonName",
		"addon_version":            "AddonVersion",
		"arn":                      "Arn",
		"cluster_name":             "ClusterName",
		"key":                      "Key",
		"resolve_conflicts":        "ResolveConflicts",
		"service_account_role_arn": "ServiceAccountRoleArn",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ResolveConflicts",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
