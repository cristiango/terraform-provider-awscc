// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceTypeFactory("awscc_ecs_capacity_provider", capacityProviderResourceType)
}

// capacityProviderResourceType returns the Terraform awscc_ecs_capacity_provider resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECS::CapacityProvider resource type.
func capacityProviderResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_provider": {
			// Property: AutoScalingGroupProvider
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AutoScalingGroupArn": {
			//       "type": "string"
			//     },
			//     "ManagedScaling": {
			//       "additionalProperties": false,
			//       "description": "The managed scaling settings for the Auto Scaling group capacity provider.",
			//       "properties": {
			//         "InstanceWarmupPeriod": {
			//           "type": "integer"
			//         },
			//         "MaximumScalingStepSize": {
			//           "type": "integer"
			//         },
			//         "MinimumScalingStepSize": {
			//           "type": "integer"
			//         },
			//         "Status": {
			//           "enum": [
			//             "DISABLED",
			//             "ENABLED"
			//           ],
			//           "type": "string"
			//         },
			//         "TargetCapacity": {
			//           "type": "integer"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ManagedTerminationProtection": {
			//       "enum": [
			//         "DISABLED",
			//         "ENABLED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AutoScalingGroupArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_scaling_group_arn": {
						// Property: AutoScalingGroupArn
						Type:     types.StringType,
						Required: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
					},
					"managed_scaling": {
						// Property: ManagedScaling
						Description: "The managed scaling settings for the Auto Scaling group capacity provider.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"instance_warmup_period": {
									// Property: InstanceWarmupPeriod
									Type:     types.Int64Type,
									Optional: true,
								},
								"maximum_scaling_step_size": {
									// Property: MaximumScalingStepSize
									Type:     types.Int64Type,
									Optional: true,
								},
								"minimum_scaling_step_size": {
									// Property: MinimumScalingStepSize
									Type:     types.Int64Type,
									Optional: true,
								},
								"status": {
									// Property: Status
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"DISABLED",
											"ENABLED",
										}),
									},
								},
								"target_capacity": {
									// Property: TargetCapacity
									Type:     types.Int64Type,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"managed_termination_protection": {
						// Property: ManagedTerminationProtection
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"DISABLED",
								"ENABLED",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "minLength": 1,
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
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
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
		Description: "Resource Type definition for AWS::ECS::CapacityProvider.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::CapacityProvider").WithTerraformTypeName("awscc_ecs_capacity_provider")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_arn":         "AutoScalingGroupArn",
		"auto_scaling_group_provider":    "AutoScalingGroupProvider",
		"instance_warmup_period":         "InstanceWarmupPeriod",
		"key":                            "Key",
		"managed_scaling":                "ManagedScaling",
		"managed_termination_protection": "ManagedTerminationProtection",
		"maximum_scaling_step_size":      "MaximumScalingStepSize",
		"minimum_scaling_step_size":      "MinimumScalingStepSize",
		"name":                           "Name",
		"status":                         "Status",
		"tags":                           "Tags",
		"target_capacity":                "TargetCapacity",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
