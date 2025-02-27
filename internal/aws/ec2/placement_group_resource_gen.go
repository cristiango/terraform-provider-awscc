// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_placement_group", placementGroupResourceType)
}

// placementGroupResourceType returns the Terraform awscc_ec2_placement_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::PlacementGroup resource type.
func placementGroupResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"group_name": {
			// Property: GroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Group Name of Placement Group.",
			//   "type": "string"
			// }
			Description: "The Group Name of Placement Group.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"spread_level": {
			// Property: SpreadLevel
			// CloudFormation resource type schema:
			// {
			//   "description": "The Spread Level of Placement Group is an enum where it accepts either host or rack when strategy is spread",
			//   "type": "string"
			// }
			Description: "The Spread Level of Placement Group is an enum where it accepts either host or rack when strategy is spread",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"strategy": {
			// Property: Strategy
			// CloudFormation resource type schema:
			// {
			//   "description": "The placement strategy.",
			//   "type": "string"
			// }
			Description: "The placement strategy.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
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
		Description: "Resource Type definition for AWS::EC2::PlacementGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::PlacementGroup").WithTerraformTypeName("awscc_ec2_placement_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_name":   "GroupName",
		"spread_level": "SpreadLevel",
		"strategy":     "Strategy",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
