// Code generated by generators/resource/main.go; DO NOT EDIT.

package resourcegroups

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
	registry.AddResourceTypeFactory("awscc_resourcegroups_group", groupResourceType)
}

// groupResourceType returns the Terraform awscc_resourcegroups_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ResourceGroups::Group resource type.
func groupResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Resource Group ARN.",
			//   "type": "string"
			// }
			Description: "The Resource Group ARN.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Parameters": {
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Name": {
			//               "type": "string"
			//             },
			//             "Values": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "Type": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameters": {
						// Property: Parameters
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"values": {
									// Property: Values
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the resource group",
			//   "maxLength": 512,
			//   "type": "string"
			// }
			Description: "The description of the resource group",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(512),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the resource group",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "The name of the resource group",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resource_query": {
			// Property: ResourceQuery
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Query": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ResourceTypeFilters": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "StackIdentifier": {
			//           "type": "string"
			//         },
			//         "TagFilters": {
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Key": {
			//                 "type": "string"
			//               },
			//               "Values": {
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Type": {
			//       "enum": [
			//         "TAG_FILTERS_1_0",
			//         "CLOUDFORMATION_STACK_1_0"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"query": {
						// Property: Query
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_type_filters": {
									// Property: ResourceTypeFilters
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"stack_identifier": {
									// Property: StackIdentifier
									Type:     types.StringType,
									Optional: true,
								},
								"tag_filters": {
									// Property: TagFilters
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
											},
											"values": {
												// Property: Values
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"TAG_FILTERS_1_0",
								"CLOUDFORMATION_STACK_1_0",
							}),
						},
					},
				},
			),
			Optional: true,
		},
		"resources": {
			// Property: Resources
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
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
		Description: "Schema for ResourceGroups::Group",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResourceGroups::Group").WithTerraformTypeName("awscc_resourcegroups_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"configuration":         "Configuration",
		"description":           "Description",
		"key":                   "Key",
		"name":                  "Name",
		"parameters":            "Parameters",
		"query":                 "Query",
		"resource_query":        "ResourceQuery",
		"resource_type_filters": "ResourceTypeFilters",
		"resources":             "Resources",
		"stack_identifier":      "StackIdentifier",
		"tag_filters":           "TagFilters",
		"tags":                  "Tags",
		"type":                  "Type",
		"value":                 "Value",
		"values":                "Values",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
