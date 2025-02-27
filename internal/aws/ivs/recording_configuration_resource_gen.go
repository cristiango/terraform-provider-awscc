// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

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
	registry.AddResourceTypeFactory("awscc_ivs_recording_configuration", recordingConfigurationResourceType)
}

// recordingConfigurationResourceType returns the Terraform awscc_ivs_recording_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IVS::RecordingConfiguration resource type.
func recordingConfigurationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[-a-z]*:ivs:[a-z0-9-]+:[0-9]+:recording-configuration/[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Description: "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"destination_configuration": {
			// Property: DestinationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Recording Destination Configuration.",
			//   "properties": {
			//     "S3": {
			//       "additionalProperties": false,
			//       "description": "Recording S3 Destination Configuration.",
			//       "properties": {
			//         "BucketName": {
			//           "maxLength": 63,
			//           "minLength": 3,
			//           "pattern": "^[a-z0-9-.]+$",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "BucketName"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "S3"
			//   ],
			//   "type": "object"
			// }
			Description: "Recording Destination Configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3": {
						// Property: S3
						Description: "Recording S3 Destination Configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_name": {
									// Property: BucketName
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 63),
										validate.StringMatch(regexp.MustCompile("^[a-z0-9-.]+$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.RequiresReplace(),
									},
								},
							},
						),
						Required: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration Name.",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "^[a-zA-Z0-9-_]*$",
			//   "type": "string"
			// }
			Description: "Recording Configuration Name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 128),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "Recording Configuration State.",
			//   "enum": [
			//     "CREATING",
			//     "CREATE_FAILED",
			//     "ACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "Recording Configuration State.",
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
			//   "description": "A list of key-value pairs that contain metadata for the asset model.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"thumbnail_configuration": {
			// Property: ThumbnailConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Recording Thumbnail Configuration.",
			//   "properties": {
			//     "RecordingMode": {
			//       "description": "Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.",
			//       "enum": [
			//         "INTERVAL",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     },
			//     "TargetIntervalSeconds": {
			//       "description": "Thumbnail recording Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.",
			//       "maximum": 60,
			//       "minimum": 5,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "RecordingMode"
			//   ],
			//   "type": "object"
			// }
			Description: "Recording Thumbnail Configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"recording_mode": {
						// Property: RecordingMode
						Description: "Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INTERVAL",
								"DISABLED",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
					},
					"target_interval_seconds": {
						// Property: TargetIntervalSeconds
						Description: "Thumbnail recording Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(5, 60),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::IVS::RecordingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::RecordingConfiguration").WithTerraformTypeName("awscc_ivs_recording_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"bucket_name":               "BucketName",
		"destination_configuration": "DestinationConfiguration",
		"key":                       "Key",
		"name":                      "Name",
		"recording_mode":            "RecordingMode",
		"s3":                        "S3",
		"state":                     "State",
		"tags":                      "Tags",
		"target_interval_seconds":   "TargetIntervalSeconds",
		"thumbnail_configuration":   "ThumbnailConfiguration",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
