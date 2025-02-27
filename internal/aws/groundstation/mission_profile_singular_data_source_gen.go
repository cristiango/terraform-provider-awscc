// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_groundstation_mission_profile", missionProfileDataSourceType)
}

// missionProfileDataSourceType returns the Terraform awscc_groundstation_mission_profile data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GroundStation::MissionProfile resource type.
func missionProfileDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"contact_post_pass_duration_seconds": {
			// Property: ContactPostPassDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Post-pass time needed after the contact.",
			//   "type": "integer"
			// }
			Description: "Post-pass time needed after the contact.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"contact_pre_pass_duration_seconds": {
			// Property: ContactPrePassDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Pre-pass time needed before the contact.",
			//   "type": "integer"
			// }
			Description: "Pre-pass time needed before the contact.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"dataflow_edges": {
			// Property: DataflowEdges
			// CloudFormation resource type schema:
			// {
			//   "description": "",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "type": "string"
			//       },
			//       "Source": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Type:     types.StringType,
						Computed: true,
					},
					"source": {
						// Property: Source
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"minimum_viable_contact_duration_seconds": {
			// Property: MinimumViableContactDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
			//   "type": "integer"
			// }
			Description: "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name used to identify a mission profile.",
			//   "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
			//   "type": "string"
			// }
			Description: "A name used to identify a mission profile.",
			Type:        types.StringType,
			Computed:    true,
		},
		"region": {
			// Property: Region
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
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
		"tracking_config_arn": {
			// Property: TrackingConfigArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::GroundStation::MissionProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::MissionProfile").WithTerraformTypeName("awscc_groundstation_mission_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"contact_post_pass_duration_seconds": "ContactPostPassDurationSeconds",
		"contact_pre_pass_duration_seconds":  "ContactPrePassDurationSeconds",
		"dataflow_edges":                     "DataflowEdges",
		"destination":                        "Destination",
		"id":                                 "Id",
		"key":                                "Key",
		"minimum_viable_contact_duration_seconds": "MinimumViableContactDurationSeconds",
		"name":                "Name",
		"region":              "Region",
		"source":              "Source",
		"tags":                "Tags",
		"tracking_config_arn": "TrackingConfigArn",
		"value":               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
