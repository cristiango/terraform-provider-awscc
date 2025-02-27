// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

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
	registry.AddResourceTypeFactory("awscc_cloudfront_realtime_log_config", realtimeLogConfigResourceType)
}

// realtimeLogConfigResourceType returns the Terraform awscc_cloudfront_realtime_log_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::RealtimeLogConfig resource type.
func realtimeLogConfigResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"end_points": {
			// Property: EndPoints
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "KinesisStreamConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "RoleArn": {
			//             "type": "string"
			//           },
			//           "StreamArn": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "RoleArn",
			//           "StreamArn"
			//         ],
			//         "type": "object"
			//       },
			//       "StreamType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "KinesisStreamConfig",
			//       "StreamType"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"kinesis_stream_config": {
						// Property: KinesisStreamConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Required: true,
								},
								"stream_arn": {
									// Property: StreamArn
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"stream_type": {
						// Property: StreamType
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
		},
		"fields": {
			// Property: Fields
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"sampling_rate": {
			// Property: SamplingRate
			// CloudFormation resource type schema:
			// {
			//   "maximum": 100,
			//   "minimum": 1,
			//   "type": "number"
			// }
			Type:     types.Float64Type,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.FloatBetween(1.000000, 100.000000),
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
		Description: "Resource Type definition for AWS::CloudFront::RealtimeLogConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::RealtimeLogConfig").WithTerraformTypeName("awscc_cloudfront_realtime_log_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"end_points":            "EndPoints",
		"fields":                "Fields",
		"kinesis_stream_config": "KinesisStreamConfig",
		"name":                  "Name",
		"role_arn":              "RoleArn",
		"sampling_rate":         "SamplingRate",
		"stream_arn":            "StreamArn",
		"stream_type":           "StreamType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
