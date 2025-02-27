// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

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
	registry.AddResourceTypeFactory("awscc_imagebuilder_distribution_configuration", distributionConfigurationResourceType)
}

// distributionConfigurationResourceType returns the Terraform awscc_imagebuilder_distribution_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::DistributionConfiguration resource type.
func distributionConfigurationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the distribution configuration.",
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
			//   "description": "The description of the distribution configuration.",
			//   "type": "string"
			// }
			Description: "The description of the distribution configuration.",
			Type:        types.StringType,
			Optional:    true,
		},
		"distributions": {
			// Property: Distributions
			// CloudFormation resource type schema:
			// {
			//   "description": "The distributions of the distribution configuration.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The distributions of the distribution configuration.",
			//     "properties": {
			//       "AmiDistributionConfiguration": {
			//         "additionalProperties": false,
			//         "description": "The specific AMI settings (for example, launch permissions, AMI tags).",
			//         "properties": {
			//           "AmiTags": {
			//             "additionalProperties": false,
			//             "description": "The tags to apply to AMIs distributed to this Region.",
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Description": {
			//             "description": "The description of the AMI distribution configuration.",
			//             "type": "string"
			//           },
			//           "KmsKeyId": {
			//             "description": "The KMS key identifier used to encrypt the distributed image.",
			//             "type": "string"
			//           },
			//           "LaunchPermissionConfiguration": {
			//             "additionalProperties": false,
			//             "description": "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
			//             "properties": {
			//               "OrganizationArns": {
			//                 "description": "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "OrganizationalUnitArns": {
			//                 "description": "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "UserGroups": {
			//                 "description": "The name of the group.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "UserIds": {
			//                 "description": "The AWS account ID.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Name": {
			//             "description": "The name of the AMI distribution configuration.",
			//             "type": "string"
			//           },
			//           "TargetAccountIds": {
			//             "description": "The ID of accounts to which you want to distribute an image.",
			//             "insertionOrder": true,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ContainerDistributionConfiguration": {
			//         "additionalProperties": false,
			//         "description": "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
			//         "properties": {
			//           "ContainerTags": {
			//             "description": "Tags that are attached to the container distribution configuration.",
			//             "insertionOrder": true,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "Description": {
			//             "description": "The description of the container distribution configuration.",
			//             "type": "string"
			//           },
			//           "TargetRepository": {
			//             "additionalProperties": false,
			//             "description": "The destination repository for the container distribution configuration.",
			//             "properties": {
			//               "RepositoryName": {
			//                 "description": "The repository name of target container repository.",
			//                 "type": "string"
			//               },
			//               "Service": {
			//                 "description": "The service of target container repository.",
			//                 "enum": [
			//                   "ECR"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "FastLaunchConfigurations": {
			//         "description": "The Windows faster-launching configurations to use for AMI distribution.",
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "The Windows faster-launching configuration to use for AMI distribution.",
			//           "properties": {
			//             "AccountId": {
			//               "description": "The owner account ID for the fast-launch enabled Windows AMI.",
			//               "type": "string"
			//             },
			//             "Enabled": {
			//               "description": "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
			//               "type": "boolean"
			//             },
			//             "LaunchTemplate": {
			//               "additionalProperties": false,
			//               "description": "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
			//               "properties": {
			//                 "LaunchTemplateId": {
			//                   "description": "The ID of the launch template to use for faster launching for a Windows AMI.",
			//                   "type": "string"
			//                 },
			//                 "LaunchTemplateName": {
			//                   "description": "The name of the launch template to use for faster launching for a Windows AMI.",
			//                   "type": "string"
			//                 },
			//                 "LaunchTemplateVersion": {
			//                   "description": "The version of the launch template to use for faster launching for a Windows AMI.",
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "MaxParallelLaunches": {
			//               "description": "The maximum number of parallel instances that are launched for creating resources.",
			//               "type": "integer"
			//             },
			//             "SnapshotConfiguration": {
			//               "additionalProperties": false,
			//               "description": "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
			//               "properties": {
			//                 "TargetResourceCount": {
			//                   "description": "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "LaunchTemplateConfigurations": {
			//         "description": "A group of launchTemplateConfiguration settings that apply to image distribution.",
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "launchTemplateConfiguration settings that apply to image distribution.",
			//           "properties": {
			//             "AccountId": {
			//               "description": "The account ID that this configuration applies to.",
			//               "type": "string"
			//             },
			//             "LaunchTemplateId": {
			//               "description": "Identifies the EC2 launch template to use.",
			//               "type": "string"
			//             },
			//             "SetDefaultVersion": {
			//               "description": "Set the specified EC2 launch template as the default launch template for the specified account.",
			//               "type": "boolean"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "LicenseConfigurationArns": {
			//         "description": "The License Manager Configuration to associate with the AMI in the specified Region.",
			//         "insertionOrder": true,
			//         "items": {
			//           "description": "The Amazon Resource Name (ARN) of the License Manager configuration.",
			//           "type": "string"
			//         },
			//         "type": "array"
			//       },
			//       "Region": {
			//         "description": "region",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Region"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The distributions of the distribution configuration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"ami_distribution_configuration": {
						// Property: AmiDistributionConfiguration
						Description: "The specific AMI settings (for example, launch permissions, AMI tags).",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"ami_tags": {
									// Property: AmiTags
									Description: "The tags to apply to AMIs distributed to this Region.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"description": {
									// Property: Description
									Description: "The description of the AMI distribution configuration.",
									Type:        types.StringType,
									Optional:    true,
								},
								"kms_key_id": {
									// Property: KmsKeyId
									Description: "The KMS key identifier used to encrypt the distributed image.",
									Type:        types.StringType,
									Optional:    true,
								},
								"launch_permission_configuration": {
									// Property: LaunchPermissionConfiguration
									Description: "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"organization_arns": {
												// Property: OrganizationArns
												Description: "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
												Type:        types.ListType{ElemType: types.StringType},
												Optional:    true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
											"organizational_unit_arns": {
												// Property: OrganizationalUnitArns
												Description: "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
												Type:        types.ListType{ElemType: types.StringType},
												Optional:    true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
											"user_groups": {
												// Property: UserGroups
												Description: "The name of the group.",
												Type:        types.ListType{ElemType: types.StringType},
												Optional:    true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
											"user_ids": {
												// Property: UserIds
												Description: "The AWS account ID.",
												Type:        types.ListType{ElemType: types.StringType},
												Optional:    true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													Multiset(),
												},
											},
										},
									),
									Optional: true,
								},
								"name": {
									// Property: Name
									Description: "The name of the AMI distribution configuration.",
									Type:        types.StringType,
									Optional:    true,
								},
								"target_account_ids": {
									// Property: TargetAccountIds
									Description: "The ID of accounts to which you want to distribute an image.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"container_distribution_configuration": {
						// Property: ContainerDistributionConfiguration
						Description: "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"container_tags": {
									// Property: ContainerTags
									Description: "Tags that are attached to the container distribution configuration.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
								},
								"description": {
									// Property: Description
									Description: "The description of the container distribution configuration.",
									Type:        types.StringType,
									Optional:    true,
								},
								"target_repository": {
									// Property: TargetRepository
									Description: "The destination repository for the container distribution configuration.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"repository_name": {
												// Property: RepositoryName
												Description: "The repository name of target container repository.",
												Type:        types.StringType,
												Optional:    true,
											},
											"service": {
												// Property: Service
												Description: "The service of target container repository.",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"ECR",
													}),
												},
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"fast_launch_configurations": {
						// Property: FastLaunchConfigurations
						Description: "The Windows faster-launching configurations to use for AMI distribution.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"account_id": {
									// Property: AccountId
									Description: "The owner account ID for the fast-launch enabled Windows AMI.",
									Type:        types.StringType,
									Optional:    true,
								},
								"enabled": {
									// Property: Enabled
									Description: "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"launch_template": {
									// Property: LaunchTemplate
									Description: "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"launch_template_id": {
												// Property: LaunchTemplateId
												Description: "The ID of the launch template to use for faster launching for a Windows AMI.",
												Type:        types.StringType,
												Optional:    true,
											},
											"launch_template_name": {
												// Property: LaunchTemplateName
												Description: "The name of the launch template to use for faster launching for a Windows AMI.",
												Type:        types.StringType,
												Optional:    true,
											},
											"launch_template_version": {
												// Property: LaunchTemplateVersion
												Description: "The version of the launch template to use for faster launching for a Windows AMI.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
								"max_parallel_launches": {
									// Property: MaxParallelLaunches
									Description: "The maximum number of parallel instances that are launched for creating resources.",
									Type:        types.Int64Type,
									Optional:    true,
								},
								"snapshot_configuration": {
									// Property: SnapshotConfiguration
									Description: "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"target_resource_count": {
												// Property: TargetResourceCount
												Description: "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
												Type:        types.Int64Type,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"launch_template_configurations": {
						// Property: LaunchTemplateConfigurations
						Description: "A group of launchTemplateConfiguration settings that apply to image distribution.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"account_id": {
									// Property: AccountId
									Description: "The account ID that this configuration applies to.",
									Type:        types.StringType,
									Optional:    true,
								},
								"launch_template_id": {
									// Property: LaunchTemplateId
									Description: "Identifies the EC2 launch template to use.",
									Type:        types.StringType,
									Optional:    true,
								},
								"set_default_version": {
									// Property: SetDefaultVersion
									Description: "Set the specified EC2 launch template as the default launch template for the specified account.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"license_configuration_arns": {
						// Property: LicenseConfigurationArns
						Description: "The License Manager Configuration to associate with the AMI in the specified Region.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"region": {
						// Property: Region
						Description: "region",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the distribution configuration.",
			//   "type": "string"
			// }
			Description: "The name of the distribution configuration.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags associated with the component.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags associated with the component.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
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
		Description: "Resource schema for AWS::ImageBuilder::DistributionConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::DistributionConfiguration").WithTerraformTypeName("awscc_imagebuilder_distribution_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                           "AccountId",
		"ami_distribution_configuration":       "AmiDistributionConfiguration",
		"ami_tags":                             "AmiTags",
		"arn":                                  "Arn",
		"container_distribution_configuration": "ContainerDistributionConfiguration",
		"container_tags":                       "ContainerTags",
		"description":                          "Description",
		"distributions":                        "Distributions",
		"enabled":                              "Enabled",
		"fast_launch_configurations":           "FastLaunchConfigurations",
		"kms_key_id":                           "KmsKeyId",
		"launch_permission_configuration":      "LaunchPermissionConfiguration",
		"launch_template":                      "LaunchTemplate",
		"launch_template_configurations":       "LaunchTemplateConfigurations",
		"launch_template_id":                   "LaunchTemplateId",
		"launch_template_name":                 "LaunchTemplateName",
		"launch_template_version":              "LaunchTemplateVersion",
		"license_configuration_arns":           "LicenseConfigurationArns",
		"max_parallel_launches":                "MaxParallelLaunches",
		"name":                                 "Name",
		"organization_arns":                    "OrganizationArns",
		"organizational_unit_arns":             "OrganizationalUnitArns",
		"region":                               "Region",
		"repository_name":                      "RepositoryName",
		"service":                              "Service",
		"set_default_version":                  "SetDefaultVersion",
		"snapshot_configuration":               "SnapshotConfiguration",
		"tags":                                 "Tags",
		"target_account_ids":                   "TargetAccountIds",
		"target_repository":                    "TargetRepository",
		"target_resource_count":                "TargetResourceCount",
		"user_groups":                          "UserGroups",
		"user_ids":                             "UserIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
