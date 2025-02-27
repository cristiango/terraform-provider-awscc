// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudfront_response_headers_policy", responseHeadersPolicyDataSourceType)
}

// responseHeadersPolicyDataSourceType returns the Terraform awscc_cloudfront_response_headers_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFront::ResponseHeadersPolicy resource type.
func responseHeadersPolicyDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"response_headers_policy_config": {
			// Property: ResponseHeadersPolicyConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "CorsConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AccessControlAllowCredentials": {
			//           "type": "boolean"
			//         },
			//         "AccessControlAllowHeaders": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Items": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "required": [
			//             "Items"
			//           ],
			//           "type": "object"
			//         },
			//         "AccessControlAllowMethods": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Items": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "required": [
			//             "Items"
			//           ],
			//           "type": "object"
			//         },
			//         "AccessControlAllowOrigins": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Items": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "required": [
			//             "Items"
			//           ],
			//           "type": "object"
			//         },
			//         "AccessControlExposeHeaders": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Items": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "required": [
			//             "Items"
			//           ],
			//           "type": "object"
			//         },
			//         "AccessControlMaxAgeSec": {
			//           "type": "integer"
			//         },
			//         "OriginOverride": {
			//           "type": "boolean"
			//         }
			//       },
			//       "required": [
			//         "AccessControlAllowOrigins",
			//         "AccessControlAllowHeaders",
			//         "AccessControlAllowMethods",
			//         "AccessControlAllowCredentials",
			//         "OriginOverride"
			//       ],
			//       "type": "object"
			//     },
			//     "CustomHeadersConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Items": {
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Header": {
			//                 "type": "string"
			//               },
			//               "Override": {
			//                 "type": "boolean"
			//               },
			//               "Value": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Header",
			//               "Value",
			//               "Override"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "Items"
			//       ],
			//       "type": "object"
			//     },
			//     "Name": {
			//       "type": "string"
			//     },
			//     "SecurityHeadersConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ContentSecurityPolicy": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ContentSecurityPolicy": {
			//               "type": "string"
			//             },
			//             "Override": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Override",
			//             "ContentSecurityPolicy"
			//           ],
			//           "type": "object"
			//         },
			//         "ContentTypeOptions": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Override": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Override"
			//           ],
			//           "type": "object"
			//         },
			//         "FrameOptions": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "FrameOption": {
			//               "pattern": "^(DENY|SAMEORIGIN)$",
			//               "type": "string"
			//             },
			//             "Override": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Override",
			//             "FrameOption"
			//           ],
			//           "type": "object"
			//         },
			//         "ReferrerPolicy": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Override": {
			//               "type": "boolean"
			//             },
			//             "ReferrerPolicy": {
			//               "pattern": "^(no-referrer|no-referrer-when-downgrade|origin|origin-when-cross-origin|same-origin|strict-origin|strict-origin-when-cross-origin|unsafe-url)$",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Override",
			//             "ReferrerPolicy"
			//           ],
			//           "type": "object"
			//         },
			//         "StrictTransportSecurity": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AccessControlMaxAgeSec": {
			//               "type": "integer"
			//             },
			//             "IncludeSubdomains": {
			//               "type": "boolean"
			//             },
			//             "Override": {
			//               "type": "boolean"
			//             },
			//             "Preload": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Override",
			//             "AccessControlMaxAgeSec"
			//           ],
			//           "type": "object"
			//         },
			//         "XSSProtection": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ModeBlock": {
			//               "type": "boolean"
			//             },
			//             "Override": {
			//               "type": "boolean"
			//             },
			//             "Protection": {
			//               "type": "boolean"
			//             },
			//             "ReportUri": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Override",
			//             "Protection"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ServerTimingHeadersConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Enabled": {
			//           "type": "boolean"
			//         },
			//         "SamplingRate": {
			//           "maximum": 100,
			//           "minimum": 0,
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "Enabled"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Name"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Computed: true,
					},
					"cors_config": {
						// Property: CorsConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"access_control_allow_credentials": {
									// Property: AccessControlAllowCredentials
									Type:     types.BoolType,
									Computed: true,
								},
								"access_control_allow_headers": {
									// Property: AccessControlAllowHeaders
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"items": {
												// Property: Items
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"access_control_allow_methods": {
									// Property: AccessControlAllowMethods
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"items": {
												// Property: Items
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"access_control_allow_origins": {
									// Property: AccessControlAllowOrigins
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"items": {
												// Property: Items
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"access_control_expose_headers": {
									// Property: AccessControlExposeHeaders
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"items": {
												// Property: Items
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"access_control_max_age_sec": {
									// Property: AccessControlMaxAgeSec
									Type:     types.Int64Type,
									Computed: true,
								},
								"origin_override": {
									// Property: OriginOverride
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"custom_headers_config": {
						// Property: CustomHeadersConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"items": {
									// Property: Items
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"header": {
												// Property: Header
												Type:     types.StringType,
												Computed: true,
											},
											"override": {
												// Property: Override
												Type:     types.BoolType,
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
							},
						),
						Computed: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"security_headers_config": {
						// Property: SecurityHeadersConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"content_security_policy": {
									// Property: ContentSecurityPolicy
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"content_security_policy": {
												// Property: ContentSecurityPolicy
												Type:     types.StringType,
												Computed: true,
											},
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"content_type_options": {
									// Property: ContentTypeOptions
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"frame_options": {
									// Property: FrameOptions
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"frame_option": {
												// Property: FrameOption
												Type:     types.StringType,
												Computed: true,
											},
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"referrer_policy": {
									// Property: ReferrerPolicy
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
											"referrer_policy": {
												// Property: ReferrerPolicy
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"strict_transport_security": {
									// Property: StrictTransportSecurity
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"access_control_max_age_sec": {
												// Property: AccessControlMaxAgeSec
												Type:     types.Int64Type,
												Computed: true,
											},
											"include_subdomains": {
												// Property: IncludeSubdomains
												Type:     types.BoolType,
												Computed: true,
											},
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
											"preload": {
												// Property: Preload
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"xss_protection": {
									// Property: XSSProtection
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"mode_block": {
												// Property: ModeBlock
												Type:     types.BoolType,
												Computed: true,
											},
											"override": {
												// Property: Override
												Type:     types.BoolType,
												Computed: true,
											},
											"protection": {
												// Property: Protection
												Type:     types.BoolType,
												Computed: true,
											},
											"report_uri": {
												// Property: ReportUri
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
					"server_timing_headers_config": {
						// Property: ServerTimingHeadersConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Computed: true,
								},
								"sampling_rate": {
									// Property: SamplingRate
									Type:     types.Float64Type,
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFront::ResponseHeadersPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::ResponseHeadersPolicy").WithTerraformTypeName("awscc_cloudfront_response_headers_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_control_allow_credentials": "AccessControlAllowCredentials",
		"access_control_allow_headers":     "AccessControlAllowHeaders",
		"access_control_allow_methods":     "AccessControlAllowMethods",
		"access_control_allow_origins":     "AccessControlAllowOrigins",
		"access_control_expose_headers":    "AccessControlExposeHeaders",
		"access_control_max_age_sec":       "AccessControlMaxAgeSec",
		"comment":                          "Comment",
		"content_security_policy":          "ContentSecurityPolicy",
		"content_type_options":             "ContentTypeOptions",
		"cors_config":                      "CorsConfig",
		"custom_headers_config":            "CustomHeadersConfig",
		"enabled":                          "Enabled",
		"frame_option":                     "FrameOption",
		"frame_options":                    "FrameOptions",
		"header":                           "Header",
		"id":                               "Id",
		"include_subdomains":               "IncludeSubdomains",
		"items":                            "Items",
		"last_modified_time":               "LastModifiedTime",
		"mode_block":                       "ModeBlock",
		"name":                             "Name",
		"origin_override":                  "OriginOverride",
		"override":                         "Override",
		"preload":                          "Preload",
		"protection":                       "Protection",
		"referrer_policy":                  "ReferrerPolicy",
		"report_uri":                       "ReportUri",
		"response_headers_policy_config":   "ResponseHeadersPolicyConfig",
		"sampling_rate":                    "SamplingRate",
		"security_headers_config":          "SecurityHeadersConfig",
		"server_timing_headers_config":     "ServerTimingHeadersConfig",
		"strict_transport_security":        "StrictTransportSecurity",
		"value":                            "Value",
		"xss_protection":                   "XSSProtection",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
