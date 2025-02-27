// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotevents

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
	registry.AddResourceTypeFactory("awscc_iotevents_input", inputResourceType)
}

// inputResourceType returns the Terraform awscc_iotevents_input resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTEvents::Input resource type.
func inputResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"input_definition": {
			// Property: InputDefinition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The definition of the input.",
			//   "properties": {
			//     "Attributes": {
			//       "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
			//         "properties": {
			//           "JsonPath": {
			//             "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `\u003cfield-name\u003e.\u003cfield-name\u003e...`",
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "^((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+))(\\.((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+)))*$",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "JsonPath"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 200,
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "required": [
			//     "Attributes"
			//   ],
			//   "type": "object"
			// }
			Description: "The definition of the input.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Description: "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"json_path": {
									// Property: JsonPath
									Description: "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (`BatchPutMessage`). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.\n\n_Syntax_: `<field-name>.<field-name>...`",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
										validate.StringMatch(regexp.MustCompile("^((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+))(\\.((`[a-zA-Z0-9_\\- ]+`)|([a-zA-Z0-9_\\-]+)))*$"), ""),
									},
								},
							},
						),
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 200),
						},
					},
				},
			),
			Required: true,
		},
		"input_description": {
			// Property: InputDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A brief description of the input.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A brief description of the input.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
		},
		"input_name": {
			// Property: InputName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the input.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z][a-zA-Z0-9_]*$",
			//   "type": "string"
			// }
			Description: "The name of the input.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Tags to be applied to Input.",
			//     "properties": {
			//       "Key": {
			//         "description": "Key of the Tag.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Value of the Tag.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "Key of the Tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "Value of the Tag.",
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

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTEvents::Input").WithTerraformTypeName("awscc_iotevents_input")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":        "Attributes",
		"input_definition":  "InputDefinition",
		"input_description": "InputDescription",
		"input_name":        "InputName",
		"json_path":         "JsonPath",
		"key":               "Key",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
