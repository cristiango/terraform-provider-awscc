// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_networkfirewall_firewall_policy", firewallPolicyDataSourceType)
}

// firewallPolicyDataSourceType returns the Terraform awscc_networkfirewall_firewall_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NetworkFirewall::FirewallPolicy resource type.
func firewallPolicyDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "^.*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"firewall_policy": {
			// Property: FirewallPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "StatefulDefaultActions": {
			//       "insertionOrder": true,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StatefulEngineOptions": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "RuleOrder": {
			//           "enum": [
			//             "DEFAULT_ACTION_ORDER",
			//             "STRICT_ORDER"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "StatefulRuleGroupReferences": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Priority": {
			//             "maximum": 65535,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "ResourceArn": {
			//             "description": "A resource ARN.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "pattern": "^(arn:aws.*)$",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ResourceArn"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StatelessCustomActions": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ActionDefinition": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "PublishMetricAction": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Dimensions": {
			//                     "insertionOrder": true,
			//                     "items": {
			//                       "additionalProperties": false,
			//                       "properties": {
			//                         "Value": {
			//                           "maxLength": 128,
			//                           "minLength": 1,
			//                           "pattern": "^[a-zA-Z0-9-_ ]+$",
			//                           "type": "string"
			//                         }
			//                       },
			//                       "required": [
			//                         "Value"
			//                       ],
			//                       "type": "object"
			//                     },
			//                     "type": "array",
			//                     "uniqueItems": false
			//                   }
			//                 },
			//                 "required": [
			//                   "Dimensions"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ActionName": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "^[a-zA-Z0-9]+$",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ActionName",
			//           "ActionDefinition"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StatelessDefaultActions": {
			//       "insertionOrder": true,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StatelessFragmentDefaultActions": {
			//       "insertionOrder": true,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "StatelessRuleGroupReferences": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Priority": {
			//             "maximum": 65535,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "ResourceArn": {
			//             "description": "A resource ARN.",
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "pattern": "^(arn:aws.*)$",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ResourceArn",
			//           "Priority"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "required": [
			//     "StatelessDefaultActions",
			//     "StatelessFragmentDefaultActions"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stateful_default_actions": {
						// Property: StatefulDefaultActions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"stateful_engine_options": {
						// Property: StatefulEngineOptions
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"rule_order": {
									// Property: RuleOrder
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"stateful_rule_group_references": {
						// Property: StatefulRuleGroupReferences
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"priority": {
									// Property: Priority
									Type:     types.Int64Type,
									Computed: true,
								},
								"resource_arn": {
									// Property: ResourceArn
									Description: "A resource ARN.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"stateless_custom_actions": {
						// Property: StatelessCustomActions
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"action_definition": {
									// Property: ActionDefinition
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"publish_metric_action": {
												// Property: PublishMetricAction
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"dimensions": {
															// Property: Dimensions
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"value": {
																		// Property: Value
																		Type:     types.StringType,
																		Computed: true,
																	},
																},
															),
															Computed: true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"action_name": {
									// Property: ActionName
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"stateless_default_actions": {
						// Property: StatelessDefaultActions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"stateless_fragment_default_actions": {
						// Property: StatelessFragmentDefaultActions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"stateless_rule_group_references": {
						// Property: StatelessRuleGroupReferences
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"priority": {
									// Property: Priority
									Type:     types.Int64Type,
									Computed: true,
								},
								"resource_arn": {
									// Property: ResourceArn
									Description: "A resource ARN.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"firewall_policy_arn": {
			// Property: FirewallPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^(arn:aws.*)$",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"firewall_policy_id": {
			// Property: FirewallPolicyId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "^([0-9a-f]{8})-([0-9a-f]{4}-){3}([0-9a-f]{12})$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"firewall_policy_name": {
			// Property: FirewallPolicyName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^.*$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "pattern": "^.*$",
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
			Attributes: tfsdk.SetNestedAttributes(
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
		Description: "Data Source schema for AWS::NetworkFirewall::FirewallPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::FirewallPolicy").WithTerraformTypeName("awscc_networkfirewall_firewall_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action_definition":                  "ActionDefinition",
		"action_name":                        "ActionName",
		"description":                        "Description",
		"dimensions":                         "Dimensions",
		"firewall_policy":                    "FirewallPolicy",
		"firewall_policy_arn":                "FirewallPolicyArn",
		"firewall_policy_id":                 "FirewallPolicyId",
		"firewall_policy_name":               "FirewallPolicyName",
		"key":                                "Key",
		"priority":                           "Priority",
		"publish_metric_action":              "PublishMetricAction",
		"resource_arn":                       "ResourceArn",
		"rule_order":                         "RuleOrder",
		"stateful_default_actions":           "StatefulDefaultActions",
		"stateful_engine_options":            "StatefulEngineOptions",
		"stateful_rule_group_references":     "StatefulRuleGroupReferences",
		"stateless_custom_actions":           "StatelessCustomActions",
		"stateless_default_actions":          "StatelessDefaultActions",
		"stateless_fragment_default_actions": "StatelessFragmentDefaultActions",
		"stateless_rule_group_references":    "StatelessRuleGroupReferences",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
