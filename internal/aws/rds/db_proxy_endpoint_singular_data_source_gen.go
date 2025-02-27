// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_rds_db_proxy_endpoint", dBProxyEndpointDataSourceType)
}

// dBProxyEndpointDataSourceType returns the Terraform awscc_rds_db_proxy_endpoint data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::RDS::DBProxyEndpoint resource type.
func dBProxyEndpointDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"db_proxy_endpoint_arn": {
			// Property: DBProxyEndpointArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
			//   "pattern": "arn:aws[A-Za-z0-9-]{0,64}:rds:[A-Za-z0-9-]{1,64}:[0-9]{12}:.*",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"db_proxy_endpoint_name": {
			// Property: DBProxyEndpointName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
			//   "maxLength": 64,
			//   "pattern": "[0-z]*",
			//   "type": "string"
			// }
			Description: "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"db_proxy_name": {
			// Property: DBProxyName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			//   "maxLength": 64,
			//   "pattern": "[0-z]*",
			//   "type": "string"
			// }
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"is_default": {
			// Property: IsDefault
			// CloudFormation resource type schema:
			// {
			//   "description": "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
			//   "type": "boolean"
			// }
			Description: "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
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
		"target_role": {
			// Property: TargetRole
			// CloudFormation resource type schema:
			// {
			//   "description": "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
			//   "enum": [
			//     "READ_WRITE",
			//     "READ_ONLY"
			//   ],
			//   "type": "string"
			// }
			Description: "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC ID to associate with the new DB proxy endpoint.",
			//   "type": "string"
			// }
			Description: "VPC ID to associate with the new DB proxy endpoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_security_group_ids": {
			// Property: VpcSecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC security group IDs to associate with the new DB proxy endpoint.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "VPC security group IDs to associate with the new DB proxy endpoint.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"vpc_subnet_ids": {
			// Property: VpcSubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC subnet IDs to associate with the new DB proxy endpoint.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 2,
			//   "type": "array"
			// }
			Description: "VPC subnet IDs to associate with the new DB proxy endpoint.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::RDS::DBProxyEndpoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyEndpoint").WithTerraformTypeName("awscc_rds_db_proxy_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"db_proxy_endpoint_arn":  "DBProxyEndpointArn",
		"db_proxy_endpoint_name": "DBProxyEndpointName",
		"db_proxy_name":          "DBProxyName",
		"endpoint":               "Endpoint",
		"is_default":             "IsDefault",
		"key":                    "Key",
		"tags":                   "Tags",
		"target_role":            "TargetRole",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_subnet_ids":         "VpcSubnetIds",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
