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
	registry.AddDataSourceTypeFactory("awscc_ec2_dhcp_options", dHCPOptionsDataSourceType)
}

// dHCPOptionsDataSourceType returns the Terraform awscc_ec2_dhcp_options data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::DHCPOptions resource type.
func dHCPOptionsDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"dhcp_options_id": {
			// Property: DhcpOptionsId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "This value is used to complete unqualified DNS hostnames.",
			//   "type": "string"
			// }
			Description: "This value is used to complete unqualified DNS hostnames.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name_servers": {
			// Property: DomainNameServers
			// CloudFormation resource type schema:
			// {
			//   "description": "The IPv4 addresses of up to four domain name servers, or AmazonProvidedDNS.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The IPv4 addresses of up to four domain name servers, or AmazonProvidedDNS.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"netbios_name_servers": {
			// Property: NetbiosNameServers
			// CloudFormation resource type schema:
			// {
			//   "description": "The IPv4 addresses of up to four NetBIOS name servers.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The IPv4 addresses of up to four NetBIOS name servers.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"netbios_node_type": {
			// Property: NetbiosNodeType
			// CloudFormation resource type schema:
			// {
			//   "description": "The NetBIOS node type (1, 2, 4, or 8).",
			//   "type": "integer"
			// }
			Description: "The NetBIOS node type (1, 2, 4, or 8).",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"ntp_servers": {
			// Property: NtpServers
			// CloudFormation resource type schema:
			// {
			//   "description": "The IPv4 addresses of up to four Network Time Protocol (NTP) servers.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The IPv4 addresses of up to four Network Time Protocol (NTP) servers.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Any tags assigned to the DHCP options set.",
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
			//         "minLength": 0,
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
			Description: "Any tags assigned to the DHCP options set.",
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::DHCPOptions",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::DHCPOptions").WithTerraformTypeName("awscc_ec2_dhcp_options")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dhcp_options_id":      "DhcpOptionsId",
		"domain_name":          "DomainName",
		"domain_name_servers":  "DomainNameServers",
		"key":                  "Key",
		"netbios_name_servers": "NetbiosNameServers",
		"netbios_node_type":    "NetbiosNodeType",
		"ntp_servers":          "NtpServers",
		"tags":                 "Tags",
		"value":                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
