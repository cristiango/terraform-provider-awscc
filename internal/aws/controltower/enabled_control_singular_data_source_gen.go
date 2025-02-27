// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package controltower

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_controltower_enabled_control", enabledControlDataSourceType)
}

// enabledControlDataSourceType returns the Terraform awscc_controltower_enabled_control data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ControlTower::EnabledControl resource type.
func enabledControlDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"control_identifier": {
			// Property: ControlIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn of the control.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
			//   "type": "string"
			// }
			Description: "Arn of the control.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_identifier": {
			// Property: TargetIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn for Organizational unit to which the control needs to be applied",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:aws[0-9a-zA-Z_\\-:\\/]+$",
			//   "type": "string"
			// }
			Description: "Arn for Organizational unit to which the control needs to be applied",
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
		Description: "Data Source schema for AWS::ControlTower::EnabledControl",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ControlTower::EnabledControl").WithTerraformTypeName("awscc_controltower_enabled_control")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"control_identifier": "ControlIdentifier",
		"target_identifier":  "TargetIdentifier",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
