// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_macie_allow_list", allowListResourceType)
}

// allowListResourceType returns the Terraform awscc_macie_allow_list resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Macie::AllowList resource type.
func allowListResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList ARN.",
			//   "type": "string"
			// }
			Description: "AllowList ARN.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"criteria": {
			// Property: Criteria
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList criteria.",
			//   "properties": {
			//     "Regex": {
			//       "description": "The S3 object key for the AllowList.",
			//       "type": "string"
			//     },
			//     "S3WordsList": {
			//       "additionalProperties": false,
			//       "description": "The S3 location for the AllowList.",
			//       "properties": {
			//         "BucketName": {
			//           "type": "string"
			//         },
			//         "ObjectKey": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "BucketName",
			//         "ObjectKey"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "AllowList criteria.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"regex": {
						// Property: Regex
						Description: "The S3 object key for the AllowList.",
						Type:        types.StringType,
						Optional:    true,
					},
					"s3_words_list": {
						// Property: S3WordsList
						Description: "The S3 location for the AllowList.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_name": {
									// Property: BucketName
									Type:     types.StringType,
									Required: true,
								},
								"object_key": {
									// Property: ObjectKey
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of AllowList.",
			//   "type": "string"
			// }
			Description: "Description of AllowList.",
			Type:        types.StringType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList ID.",
			//   "type": "string"
			// }
			Description: "AllowList ID.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of AllowList.",
			//   "type": "string"
			// }
			Description: "Name of AllowList.",
			Type:        types.StringType,
			Required:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList status.",
			//   "enum": [
			//     "OK",
			//     "S3_OBJECT_NOT_FOUND",
			//     "S3_USER_ACCESS_DENIED",
			//     "S3_OBJECT_ACCESS_DENIED",
			//     "S3_THROTTLED",
			//     "S3_OBJECT_OVERSIZE",
			//     "S3_OBJECT_EMPTY",
			//     "UNKNOWN_ERROR"
			//   ],
			//   "type": "string"
			// }
			Description: "AllowList status.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Macie AllowList resource schema",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::AllowList").WithTerraformTypeName("awscc_macie_allow_list")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"bucket_name":   "BucketName",
		"criteria":      "Criteria",
		"description":   "Description",
		"id":            "Id",
		"key":           "Key",
		"name":          "Name",
		"object_key":    "ObjectKey",
		"regex":         "Regex",
		"s3_words_list": "S3WordsList",
		"status":        "Status",
		"tags":          "Tags",
		"value":         "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
