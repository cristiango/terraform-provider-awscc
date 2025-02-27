// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package evidently

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_evidently_feature", featureDataSourceType)
}

// featureDataSourceType returns the Terraform awscc_evidently_feature data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Evidently::Feature resource type.
func featureDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/feature/[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"default_variation": {
			// Property: DefaultVariation
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "pattern": "[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 160,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"entity_overrides": {
			// Property: EntityOverrides
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "EntityId": {
			//         "type": "string"
			//       },
			//       "Variation": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "[-a-zA-Z0-9._]*",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"entity_id": {
						// Property: EntityId
						Type:     types.StringType,
						Computed: true,
					},
					"variation": {
						// Property: Variation
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"evaluation_strategy": {
			// Property: EvaluationStrategy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ALL_RULES",
			//     "DEFAULT_VARIATION"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "pattern": "[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"project": {
			// Property: Project
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"variations": {
			// Property: Variations
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "oneOf": [
			//       {
			//         "required": [
			//           "VariationName",
			//           "StringValue"
			//         ]
			//       },
			//       {
			//         "required": [
			//           "VariationName",
			//           "BooleanValue"
			//         ]
			//       },
			//       {
			//         "required": [
			//           "VariationName",
			//           "LongValue"
			//         ]
			//       },
			//       {
			//         "required": [
			//           "VariationName",
			//           "DoubleValue"
			//         ]
			//       }
			//     ],
			//     "properties": {
			//       "BooleanValue": {
			//         "type": "boolean"
			//       },
			//       "DoubleValue": {
			//         "type": "number"
			//       },
			//       "LongValue": {
			//         "type": "number"
			//       },
			//       "StringValue": {
			//         "maxLength": 512,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "VariationName": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "[-a-zA-Z0-9._]*",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"boolean_value": {
						// Property: BooleanValue
						Type:     types.BoolType,
						Computed: true,
					},
					"double_value": {
						// Property: DoubleValue
						Type:     types.Float64Type,
						Computed: true,
					},
					"long_value": {
						// Property: LongValue
						Type:     types.Float64Type,
						Computed: true,
					},
					"string_value": {
						// Property: StringValue
						Type:     types.StringType,
						Computed: true,
					},
					"variation_name": {
						// Property: VariationName
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Evidently::Feature",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Feature").WithTerraformTypeName("awscc_evidently_feature")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"boolean_value":       "BooleanValue",
		"default_variation":   "DefaultVariation",
		"description":         "Description",
		"double_value":        "DoubleValue",
		"entity_id":           "EntityId",
		"entity_overrides":    "EntityOverrides",
		"evaluation_strategy": "EvaluationStrategy",
		"key":                 "Key",
		"long_value":          "LongValue",
		"name":                "Name",
		"project":             "Project",
		"string_value":        "StringValue",
		"tags":                "Tags",
		"value":               "Value",
		"variation":           "Variation",
		"variation_name":      "VariationName",
		"variations":          "Variations",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
