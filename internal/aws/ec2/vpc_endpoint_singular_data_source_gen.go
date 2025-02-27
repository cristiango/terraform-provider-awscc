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
	registry.AddDataSourceTypeFactory("awscc_ec2_vpc_endpoint", vPCEndpointDataSourceType)
}

// vPCEndpointDataSourceType returns the Terraform awscc_ec2_vpc_endpoint data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::VPCEndpoint resource type.
func vPCEndpointDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_timestamp": {
			// Property: CreationTimestamp
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"dns_entries": {
			// Property: DnsEntries
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
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
		"network_interface_ids": {
			// Property: NetworkInterfaceIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"private_dns_enabled": {
			// Property: PrivateDnsEnabled
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"route_table_ids": {
			// Property: RouteTableIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"service_name": {
			// Property: ServiceName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"vpc_endpoint_type": {
			// Property: VpcEndpointType
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"vpc_id": {
			// Property: VpcId
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
		Description: "Data Source schema for AWS::EC2::VPCEndpoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCEndpoint").WithTerraformTypeName("awscc_ec2_vpc_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_timestamp":    "CreationTimestamp",
		"dns_entries":           "DnsEntries",
		"id":                    "Id",
		"network_interface_ids": "NetworkInterfaceIds",
		"policy_document":       "PolicyDocument",
		"private_dns_enabled":   "PrivateDnsEnabled",
		"route_table_ids":       "RouteTableIds",
		"security_group_ids":    "SecurityGroupIds",
		"service_name":          "ServiceName",
		"subnet_ids":            "SubnetIds",
		"vpc_endpoint_type":     "VpcEndpointType",
		"vpc_id":                "VpcId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
