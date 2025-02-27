// Code generated by generators/resource/main.go; DO NOT EDIT.

package apprunner

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_apprunner_observability_configuration", observabilityConfigurationResourceType)
}

// observabilityConfigurationResourceType returns the Terraform awscc_apprunner_observability_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppRunner::ObservabilityConfiguration resource type.
func observabilityConfigurationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"latest": {
			// Property: Latest
			// CloudFormation resource type schema:
			// {
			//   "description": "It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.",
			//   "type": "boolean"
			// }
			Description: "It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.",
			Type:        types.BoolType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"observability_configuration_arn": {
			// Property: ObservabilityConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of this ObservabilityConfiguration",
			//   "maxLength": 1011,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of this ObservabilityConfiguration",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"observability_configuration_name": {
			// Property: ObservabilityConfigurationName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.",
			//   "maxLength": 32,
			//   "minLength": 4,
			//   "pattern": "[A-Za-z0-9][A-Za-z0-9\\-_]{3,31}",
			//   "type": "string"
			// }
			Description: "A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(4, 32),
				validate.StringMatch(regexp.MustCompile("[A-Za-z0-9][A-Za-z0-9\\-_]{3,31}"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"observability_configuration_revision": {
			// Property: ObservabilityConfigurationRevision
			// CloudFormation resource type schema:
			// {
			//   "description": "The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.",
			//   "type": "integer"
			// }
			Description: "The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.",
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
			//   "description": "A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
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
			Description: "A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.",
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
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Tags is a write-only property.
		},
		"trace_configuration": {
			// Property: TraceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.",
			//   "properties": {
			//     "Vendor": {
			//       "description": "The implementation provider chosen for tracing App Runner services.",
			//       "enum": [
			//         "AWSXRAY"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Vendor"
			//   ],
			//   "type": "object"
			// }
			Description: "The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"vendor": {
						// Property: Vendor
						Description: "The implementation provider chosen for tracing App Runner services.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AWSXRAY",
							}),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
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
		Description: "The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::ObservabilityConfiguration").WithTerraformTypeName("awscc_apprunner_observability_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                                  "Key",
		"latest":                               "Latest",
		"observability_configuration_arn":      "ObservabilityConfigurationArn",
		"observability_configuration_name":     "ObservabilityConfigurationName",
		"observability_configuration_revision": "ObservabilityConfigurationRevision",
		"tags":                                 "Tags",
		"trace_configuration":                  "TraceConfiguration",
		"value":                                "Value",
		"vendor":                               "Vendor",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
