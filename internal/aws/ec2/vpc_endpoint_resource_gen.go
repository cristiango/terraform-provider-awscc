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
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_vpc_endpoint", vPCEndpointResourceType)
}

// vPCEndpointResourceType returns the Terraform awscc_ec2_vpc_endpoint resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::VPCEndpoint resource type.
func vPCEndpointResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_timestamp": {
			// Property: CreationTimestamp
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"private_dns_enabled": {
			// Property: PrivateDnsEnabled
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
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
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
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
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
		"service_name": {
			// Property: ServiceName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
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
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
		"vpc_endpoint_type": {
			// Property: VpcEndpointType
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::VPCEndpoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCEndpoint").WithTerraformTypeName("awscc_ec2_vpc_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
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

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
