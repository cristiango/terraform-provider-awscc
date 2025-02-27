// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ivs_stream_key", streamKeyDataSourceType)
}

// streamKeyDataSourceType returns the Terraform awscc_ivs_stream_key data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IVS::StreamKey resource type.
func streamKeyDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Stream Key ARN is automatically generated on creation and assigned as the unique identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:stream-key/[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Description: "Stream Key ARN is automatically generated on creation and assigned as the unique identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"channel_arn": {
			// Property: ChannelArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Channel ARN for the stream.",
			//   "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:channel/[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Description: "Channel ARN for the stream.",
			Type:        types.StringType,
			Computed:    true,
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
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"value": {
			// Property: Value
			// CloudFormation resource type schema:
			// {
			//   "description": "Stream-key value.",
			//   "type": "string"
			// }
			Description: "Stream-key value.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IVS::StreamKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::StreamKey").WithTerraformTypeName("awscc_ivs_stream_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"channel_arn": "ChannelArn",
		"key":         "Key",
		"tags":        "Tags",
		"value":       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
