// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_resiliencehub_app", appDataSourceType)
}

// appDataSourceType returns the Terraform awscc_resiliencehub_app data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ResilienceHub::App resource type.
func appDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_arn": {
			// Property: AppArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the App.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the App.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_assessment_schedule": {
			// Property: AppAssessmentSchedule
			// CloudFormation resource type schema:
			// {
			//   "description": "Assessment execution schedule.",
			//   "enum": [
			//     "Disabled",
			//     "Daily"
			//   ],
			//   "type": "string"
			// }
			Description: "Assessment execution schedule.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_template_body": {
			// Property: AppTemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full ResilienceHub app template body.",
			//   "maxLength": 5000,
			//   "minLength": 0,
			//   "pattern": "^[\\w\\s:,-\\.'{}\\[\\]:\"]+$",
			//   "type": "string"
			// }
			Description: "A string containing full ResilienceHub app template body.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "App description.",
			//   "maxLength": 500,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "App description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the app.",
			//   "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//   "type": "string"
			// }
			Description: "Name of the app.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resiliency_policy_arn": {
			// Property: ResiliencyPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_mappings": {
			// Property: ResourceMappings
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of ResourceMapping objects.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Resource mapping is used to map logical resources from template to physical resource",
			//     "properties": {
			//       "LogicalStackName": {
			//         "type": "string"
			//       },
			//       "MappingType": {
			//         "pattern": "CfnStack|Resource|Terraform",
			//         "type": "string"
			//       },
			//       "PhysicalResourceId": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AwsAccountId": {
			//             "pattern": "^[0-9]{12}$",
			//             "type": "string"
			//           },
			//           "AwsRegion": {
			//             "pattern": "^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$",
			//             "type": "string"
			//           },
			//           "Identifier": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Type": {
			//             "pattern": "Arn|Native",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Identifier",
			//           "Type"
			//         ],
			//         "type": "object"
			//       },
			//       "ResourceName": {
			//         "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//         "type": "string"
			//       },
			//       "TerraformSourceName": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MappingType",
			//       "PhysicalResourceId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of ResourceMapping objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"logical_stack_name": {
						// Property: LogicalStackName
						Type:     types.StringType,
						Computed: true,
					},
					"mapping_type": {
						// Property: MappingType
						Type:     types.StringType,
						Computed: true,
					},
					"physical_resource_id": {
						// Property: PhysicalResourceId
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"aws_account_id": {
									// Property: AwsAccountId
									Type:     types.StringType,
									Computed: true,
								},
								"aws_region": {
									// Property: AwsRegion
									Type:     types.StringType,
									Computed: true,
								},
								"identifier": {
									// Property: Identifier
									Type:     types.StringType,
									Computed: true,
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"resource_name": {
						// Property: ResourceName
						Type:     types.StringType,
						Computed: true,
					},
					"terraform_source_name": {
						// Property: TerraformSourceName
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ResilienceHub::App",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::App").WithTerraformTypeName("awscc_resiliencehub_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":                 "AppArn",
		"app_assessment_schedule": "AppAssessmentSchedule",
		"app_template_body":       "AppTemplateBody",
		"aws_account_id":          "AwsAccountId",
		"aws_region":              "AwsRegion",
		"description":             "Description",
		"identifier":              "Identifier",
		"logical_stack_name":      "LogicalStackName",
		"mapping_type":            "MappingType",
		"name":                    "Name",
		"physical_resource_id":    "PhysicalResourceId",
		"resiliency_policy_arn":   "ResiliencyPolicyArn",
		"resource_mappings":       "ResourceMappings",
		"resource_name":           "ResourceName",
		"tags":                    "Tags",
		"terraform_source_name":   "TerraformSourceName",
		"type":                    "Type",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
