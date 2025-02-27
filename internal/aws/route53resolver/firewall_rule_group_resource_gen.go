// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver

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
	registry.AddResourceTypeFactory("awscc_route53resolver_firewall_rule_group", firewallRuleGroupResourceType)
}

// firewallRuleGroupResourceType returns the Terraform awscc_route53resolver_firewall_rule_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53Resolver::FirewallRuleGroup resource type.
func firewallRuleGroupResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn",
			//   "maxLength": 600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Arn",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creator_request_id": {
			// Property: CreatorRequestId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the creator request.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The id of the creator request.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"firewall_rules": {
			// Property: FirewallRules
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRules",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Firewall Rule associating the Rule Group to a Domain List",
			//     "properties": {
			//       "Action": {
			//         "description": "Rule Action",
			//         "enum": [
			//           "ALLOW",
			//           "BLOCK",
			//           "ALERT"
			//         ],
			//         "type": "string"
			//       },
			//       "BlockOverrideDnsType": {
			//         "description": "BlockOverrideDnsType",
			//         "enum": [
			//           "CNAME"
			//         ],
			//         "type": "string"
			//       },
			//       "BlockOverrideDomain": {
			//         "description": "BlockOverrideDomain",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "BlockOverrideTtl": {
			//         "description": "BlockOverrideTtl",
			//         "maximum": 604800,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "BlockResponse": {
			//         "description": "BlockResponse",
			//         "enum": [
			//           "NODATA",
			//           "NXDOMAIN",
			//           "OVERRIDE"
			//         ],
			//         "type": "string"
			//       },
			//       "FirewallDomainListId": {
			//         "description": "ResourceId",
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Priority": {
			//         "description": "Rule Priority",
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "FirewallDomainListId",
			//       "Priority",
			//       "Action"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "FirewallRules",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"action": {
						// Property: Action
						Description: "Rule Action",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ALLOW",
								"BLOCK",
								"ALERT",
							}),
						},
					},
					"block_override_dns_type": {
						// Property: BlockOverrideDnsType
						Description: "BlockOverrideDnsType",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CNAME",
							}),
						},
					},
					"block_override_domain": {
						// Property: BlockOverrideDomain
						Description: "BlockOverrideDomain",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"block_override_ttl": {
						// Property: BlockOverrideTtl
						Description: "BlockOverrideTtl",
						Type:        types.Int64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 604800),
						},
					},
					"block_response": {
						// Property: BlockResponse
						Description: "BlockResponse",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NODATA",
								"NXDOMAIN",
								"OVERRIDE",
							}),
						},
					},
					"firewall_domain_list_id": {
						// Property: FirewallDomainListId
						Description: "ResourceId",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 64),
						},
					},
					"priority": {
						// Property: Priority
						Description: "Rule Priority",
						Type:        types.Int64Type,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "ResourceId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResourceId",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"modification_time": {
			// Property: ModificationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupName",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupName",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "AccountId",
			//   "maxLength": 32,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "AccountId",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"rule_count": {
			// Property: RuleCount
			// CloudFormation resource type schema:
			// {
			//   "description": "Count",
			//   "type": "integer"
			// }
			Description: "Count",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"share_status": {
			// Property: ShareStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			//   "enum": [
			//     "NOT_SHARED",
			//     "SHARED_WITH_ME",
			//     "SHARED_BY_ME"
			//   ],
			//   "type": "string"
			// }
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			//   "enum": [
			//     "COMPLETE",
			//     "DELETING",
			//     "UPDATING",
			//     "INACTIVE_OWNER_ACCOUNT_CLOSED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"status_message": {
			// Property: StatusMessage
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupStatus",
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupStatus",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Tags",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
					},
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53Resolver::FirewallRuleGroup.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallRuleGroup").WithTerraformTypeName("awscc_route53resolver_firewall_rule_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                  "Action",
		"arn":                     "Arn",
		"block_override_dns_type": "BlockOverrideDnsType",
		"block_override_domain":   "BlockOverrideDomain",
		"block_override_ttl":      "BlockOverrideTtl",
		"block_response":          "BlockResponse",
		"creation_time":           "CreationTime",
		"creator_request_id":      "CreatorRequestId",
		"firewall_domain_list_id": "FirewallDomainListId",
		"firewall_rules":          "FirewallRules",
		"id":                      "Id",
		"key":                     "Key",
		"modification_time":       "ModificationTime",
		"name":                    "Name",
		"owner_id":                "OwnerId",
		"priority":                "Priority",
		"rule_count":              "RuleCount",
		"share_status":            "ShareStatus",
		"status":                  "Status",
		"status_message":          "StatusMessage",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
