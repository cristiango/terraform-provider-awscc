// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_robomaker_robot", robotDataSourceType)
}

// robotDataSourceType returns the Terraform awscc_robomaker_robot data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::RoboMaker::Robot resource type.
func robotDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"architecture": {
			// Property: Architecture
			// CloudFormation resource type schema:
			// {
			//   "description": "The target architecture of the robot.",
			//   "enum": [
			//     "X86_64",
			//     "ARM64",
			//     "ARMHF"
			//   ],
			//   "type": "string"
			// }
			Description: "The target architecture of the robot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"fleet": {
			// Property: Fleet
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the fleet.",
			//   "maxLength": 1224,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the fleet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"greengrass_group_id": {
			// Property: GreengrassGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Greengrass group id.",
			//   "maxLength": 1224,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Greengrass group id.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the robot.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name for the robot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::RoboMaker::Robot",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::Robot").WithTerraformTypeName("awscc_robomaker_robot")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":        "Architecture",
		"arn":                 "Arn",
		"fleet":               "Fleet",
		"greengrass_group_id": "GreengrassGroupId",
		"name":                "Name",
		"tags":                "Tags",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
