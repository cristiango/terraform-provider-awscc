// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudwatch

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
	registry.AddResourceTypeFactory("awscc_cloudwatch_metric_stream", metricStreamResourceType)
}

// metricStreamResourceType returns the Terraform awscc_cloudwatch_metric_stream resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudWatch::MetricStream resource type.
func metricStreamResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name of the metric stream.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date of creation of the metric stream.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date of creation of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"exclude_filters": {
			// Property: ExcludeFilters
			// CloudFormation resource type schema:
			// {
			//   "description": "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure defines the metrics that will be streamed.",
			//     "properties": {
			//       "Namespace": {
			//         "description": "Only metrics with Namespace matching this value will be streamed.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Namespace"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"namespace": {
						// Property: Namespace
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(1000),
				validate.UniqueItems(),
			},
		},
		"firehose_arn": {
			// Property: FirehoseArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the Kinesis Firehose where to stream the data.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the Kinesis Firehose where to stream the data.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
		},
		"include_filters": {
			// Property: IncludeFilters
			// CloudFormation resource type schema:
			// {
			//   "description": "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure defines the metrics that will be streamed.",
			//     "properties": {
			//       "Namespace": {
			//         "description": "Only metrics with Namespace matching this value will be streamed.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Namespace"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"namespace": {
						// Property: Namespace
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(1000),
				validate.UniqueItems(),
			},
		},
		"last_update_date": {
			// Property: LastUpdateDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date of the last update of the metric stream.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date of the last update of the metric stream.",
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
			//   "description": "Name of the metric stream.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the metric stream.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"output_format": {
			// Property: OutputFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "The output format of the data streamed to the Kinesis Firehose.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The output format of the data streamed to the Kinesis Firehose.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the role that provides access to the Kinesis Firehose.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the role that provides access to the Kinesis Firehose.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "Displays the state of the Metric Stream.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Displays the state of the Metric Stream.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"statistics_configurations": {
			// Property: StatisticsConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure specifies a list of additional statistics to stream, and the metrics to stream those additional statistics for. All metrics that match the combination of metric name and namespace will be streamed with the extended statistics, no matter their dimensions.",
			//     "properties": {
			//       "AdditionalStatistics": {
			//         "description": "The additional statistics to stream for the metrics listed in IncludeMetrics.",
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 20,
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "IncludeMetrics": {
			//         "description": "An array that defines the metrics that are to have additional statistics streamed.",
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "A structure that specifies the metric name and namespace for one metric that is going to have additional statistics included in the stream.",
			//           "properties": {
			//             "MetricName": {
			//               "description": "The name of the metric.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Namespace": {
			//               "description": "The namespace of the metric.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "MetricName",
			//             "Namespace"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 100,
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "AdditionalStatistics",
			//       "IncludeMetrics"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 100,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"additional_statistics": {
						// Property: AdditionalStatistics
						Description: "The additional statistics to stream for the metrics listed in IncludeMetrics.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(20),
							validate.UniqueItems(),
						},
					},
					"include_metrics": {
						// Property: IncludeMetrics
						Description: "An array that defines the metrics that are to have additional statistics streamed.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"metric_name": {
									// Property: MetricName
									Description: "The name of the metric.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
								"namespace": {
									// Property: Namespace
									Description: "The namespace of the metric.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
							},
						),
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(100),
							validate.UniqueItems(),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(100),
				validate.UniqueItems(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A set of tags to assign to the delivery stream.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Metadata that you can assign to a Metric Stream, consisting of a key-value pair.",
			//     "properties": {
			//       "Key": {
			//         "description": "A unique identifier for the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "An optional string, which you can use to describe or define the tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A set of tags to assign to the delivery stream.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A unique identifier for the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "An optional string, which you can use to describe or define the tag.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
				validate.UniqueItems(),
			},
			// Tags is a write-only property.
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
		Description: "Resource Type definition for Metric Stream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::MetricStream").WithTerraformTypeName("awscc_cloudwatch_metric_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_statistics":     "AdditionalStatistics",
		"arn":                       "Arn",
		"creation_date":             "CreationDate",
		"exclude_filters":           "ExcludeFilters",
		"firehose_arn":              "FirehoseArn",
		"include_filters":           "IncludeFilters",
		"include_metrics":           "IncludeMetrics",
		"key":                       "Key",
		"last_update_date":          "LastUpdateDate",
		"metric_name":               "MetricName",
		"name":                      "Name",
		"namespace":                 "Namespace",
		"output_format":             "OutputFormat",
		"role_arn":                  "RoleArn",
		"state":                     "State",
		"statistics_configurations": "StatisticsConfigurations",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.AnyOfRequired(
		validate.Required(
			"firehose_arn",
			"role_arn",
			"output_format",
		),
		validate.AllOfRequired(
			validate.Required(
				"firehose_arn",
				"role_arn",
				"output_format",
			),
		),
		validate.OneOfRequired(
			validate.Required(
				"include_filters",
			),
			validate.Required(
				"exclude_filters",
			),
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
