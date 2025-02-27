// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_gateway_route_table_association", gatewayRouteTableAssociationDataSourceType)
}

// gatewayRouteTableAssociationDataSourceType returns the Terraform awscc_ec2_gateway_route_table_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::GatewayRouteTableAssociation resource type.
func gatewayRouteTableAssociationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "description": "The route table association ID.",
			//   "type": "string"
			// }
			Description: "The route table association ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"gateway_id": {
			// Property: GatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"route_table_id": {
			// Property: RouteTableId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the route table.",
			//   "type": "string"
			// }
			Description: "The ID of the route table.",
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
		Description: "Data Source schema for AWS::EC2::GatewayRouteTableAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::GatewayRouteTableAssociation").WithTerraformTypeName("awscc_ec2_gateway_route_table_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"gateway_id":     "GatewayId",
		"route_table_id": "RouteTableId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
