// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53

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
	registry.AddResourceTypeFactory("awscc_route53_hosted_zone", hostedZoneResourceType)
}

// hostedZoneResourceType returns the Terraform awscc_route53_hosted_zone resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53::HostedZone resource type.
func hostedZoneResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"hosted_zone_config": {
			// Property: HostedZoneConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A complex type that contains an optional comment.\n\nIf you don't want to specify a comment, omit the HostedZoneConfig and Comment elements.",
			//   "properties": {
			//     "Comment": {
			//       "description": "Any comments that you want to include about the hosted zone.",
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A complex type that contains an optional comment.\n\nIf you don't want to specify a comment, omit the HostedZoneConfig and Comment elements.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Description: "Any comments that you want to include about the hosted zone.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
		},
		"hosted_zone_tags": {
			// Property: HostedZoneTags
			// CloudFormation resource type schema:
			// {
			//   "description": "Adds, edits, or deletes tags for a health check or a hosted zone.\n\nFor information about using tags for cost allocation, see Using Cost Allocation Tags in the AWS Billing and Cost Management User Guide.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A complex type that contains information about a tag that you want to add or edit for the specified health check or hosted zone.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag.",
			//         "maxLength": 128,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag.",
			//         "maxLength": 256,
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
			//   "uniqueItems": true
			// }
			Description: "Adds, edits, or deletes tags for a health check or a hosted zone.\n\nFor information about using tags for cost allocation, see Using Cost Allocation Tags in the AWS Billing and Cost Management User Guide.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
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
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the domain. Specify a fully qualified domain name, for example, www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats www.example.com (without a trailing dot) and www.example.com. (with a trailing dot) as identical.\n\nIf you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of NameServers that are returned by the Fn::GetAtt intrinsic function.",
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Description: "The name of the domain. Specify a fully qualified domain name, for example, www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats www.example.com (without a trailing dot) and www.example.com. (with a trailing dot) as identical.\n\nIf you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of NameServers that are returned by the Fn::GetAtt intrinsic function.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1024),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"name_servers": {
			// Property: NameServers
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
		"query_logging_config": {
			// Property: QueryLoggingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A complex type that contains information about a configuration for DNS query logging.",
			//   "properties": {
			//     "CloudWatchLogsLogGroupArn": {
			//       "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "CloudWatchLogsLogGroupArn"
			//   ],
			//   "type": "object"
			// }
			Description: "A complex type that contains information about a configuration for DNS query logging.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cloudwatch_logs_log_group_arn": {
						// Property: CloudWatchLogsLogGroupArn
						Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"vp_cs": {
			// Property: VPCs
			// CloudFormation resource type schema:
			// {
			//   "description": "A complex type that contains information about the VPCs that are associated with the specified hosted zone.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A complex type that contains information about an Amazon VPC. Route 53 Resolver uses the records in the private hosted zone to route traffic in that VPC.",
			//     "properties": {
			//       "VPCId": {
			//         "description": "The ID of an Amazon VPC.",
			//         "type": "string"
			//       },
			//       "VPCRegion": {
			//         "description": "The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "VPCId",
			//       "VPCRegion"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A complex type that contains information about the VPCs that are associated with the specified hosted zone.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"vpc_id": {
						// Property: VPCId
						Description: "The ID of an Amazon VPC.",
						Type:        types.StringType,
						Required:    true,
					},
					"vpc_region": {
						// Property: VPCRegion
						Description: "The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53::HostedZone.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::HostedZone").WithTerraformTypeName("awscc_route53_hosted_zone")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudwatch_logs_log_group_arn": "CloudWatchLogsLogGroupArn",
		"comment":                       "Comment",
		"hosted_zone_config":            "HostedZoneConfig",
		"hosted_zone_tags":              "HostedZoneTags",
		"id":                            "Id",
		"key":                           "Key",
		"name":                          "Name",
		"name_servers":                  "NameServers",
		"query_logging_config":          "QueryLoggingConfig",
		"value":                         "Value",
		"vp_cs":                         "VPCs",
		"vpc_id":                        "VPCId",
		"vpc_region":                    "VPCRegion",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
