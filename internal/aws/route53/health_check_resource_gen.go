// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_route53_health_check", healthCheckResourceType)
}

// healthCheckResourceType returns the Terraform awscc_route53_health_check resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53::HealthCheck resource type.
func healthCheckResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"health_check_config": {
			// Property: HealthCheckConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A complex type that contains information about the health check.",
			//   "properties": {
			//     "AlarmIdentifier": {
			//       "additionalProperties": false,
			//       "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
			//       "properties": {
			//         "Name": {
			//           "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
			//           "maxLength": 256,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Region": {
			//           "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Name",
			//         "Region"
			//       ],
			//       "type": "object"
			//     },
			//     "ChildHealthChecks": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 256,
			//       "type": "array"
			//     },
			//     "EnableSNI": {
			//       "type": "boolean"
			//     },
			//     "FailureThreshold": {
			//       "maximum": 10,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "FullyQualifiedDomainName": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "HealthThreshold": {
			//       "maximum": 256,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "IPAddress": {
			//       "maxLength": 45,
			//       "pattern": "^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$",
			//       "type": "string"
			//     },
			//     "InsufficientDataHealthStatus": {
			//       "enum": [
			//         "Healthy",
			//         "LastKnownStatus",
			//         "Unhealthy"
			//       ],
			//       "type": "string"
			//     },
			//     "Inverted": {
			//       "type": "boolean"
			//     },
			//     "MeasureLatency": {
			//       "type": "boolean"
			//     },
			//     "Port": {
			//       "maximum": 65535,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "Regions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 64,
			//       "type": "array"
			//     },
			//     "RequestInterval": {
			//       "maximum": 30,
			//       "minimum": 10,
			//       "type": "integer"
			//     },
			//     "ResourcePath": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "RoutingControlArn": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SearchString": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "Type": {
			//       "enum": [
			//         "CALCULATED",
			//         "CLOUDWATCH_METRIC",
			//         "HTTP",
			//         "HTTP_STR_MATCH",
			//         "HTTPS",
			//         "HTTPS_STR_MATCH",
			//         "TCP",
			//         "RECOVERY_CONTROL"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "A complex type that contains information about the health check.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"alarm_identifier": {
						// Property: AlarmIdentifier
						Description: "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 256),
									},
								},
								"region": {
									// Property: Region
									Description: "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"child_health_checks": {
						// Property: ChildHealthChecks
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"enable_sni": {
						// Property: EnableSNI
						Type:     types.BoolType,
						Optional: true,
					},
					"failure_threshold": {
						// Property: FailureThreshold
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 10),
						},
					},
					"fully_qualified_domain_name": {
						// Property: FullyQualifiedDomainName
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
					"health_threshold": {
						// Property: HealthThreshold
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 256),
						},
					},
					"ip_address": {
						// Property: IPAddress
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(45),
							validate.StringMatch(regexp.MustCompile("^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$"), ""),
						},
					},
					"insufficient_data_health_status": {
						// Property: InsufficientDataHealthStatus
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Healthy",
								"LastKnownStatus",
								"Unhealthy",
							}),
						},
					},
					"inverted": {
						// Property: Inverted
						Type:     types.BoolType,
						Optional: true,
					},
					"measure_latency": {
						// Property: MeasureLatency
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"port": {
						// Property: Port
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 65535),
						},
					},
					"regions": {
						// Property: Regions
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(64),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"request_interval": {
						// Property: RequestInterval
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(10, 30),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"resource_path": {
						// Property: ResourcePath
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
					"routing_control_arn": {
						// Property: RoutingControlArn
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"search_string": {
						// Property: SearchString
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CALCULATED",
								"CLOUDWATCH_METRIC",
								"HTTP",
								"HTTP_STR_MATCH",
								"HTTPS",
								"HTTPS_STR_MATCH",
								"TCP",
								"RECOVERY_CONTROL",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
					},
				},
			),
			Required: true,
		},
		"health_check_id": {
			// Property: HealthCheckId
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
		"health_check_tags": {
			// Property: HealthCheckTags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
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
			Description: "An array of key-value pairs to apply to this resource.",
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
		Description: "Resource schema for AWS::Route53::HealthCheck.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::HealthCheck").WithTerraformTypeName("awscc_route53_health_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_identifier":                "AlarmIdentifier",
		"child_health_checks":             "ChildHealthChecks",
		"enable_sni":                      "EnableSNI",
		"failure_threshold":               "FailureThreshold",
		"fully_qualified_domain_name":     "FullyQualifiedDomainName",
		"health_check_config":             "HealthCheckConfig",
		"health_check_id":                 "HealthCheckId",
		"health_check_tags":               "HealthCheckTags",
		"health_threshold":                "HealthThreshold",
		"insufficient_data_health_status": "InsufficientDataHealthStatus",
		"inverted":                        "Inverted",
		"ip_address":                      "IPAddress",
		"key":                             "Key",
		"measure_latency":                 "MeasureLatency",
		"name":                            "Name",
		"port":                            "Port",
		"region":                          "Region",
		"regions":                         "Regions",
		"request_interval":                "RequestInterval",
		"resource_path":                   "ResourcePath",
		"routing_control_arn":             "RoutingControlArn",
		"search_string":                   "SearchString",
		"type":                            "Type",
		"value":                           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
