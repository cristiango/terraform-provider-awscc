// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package customerprofiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_customerprofiles_integration", integrationDataSourceType)
}

// integrationDataSourceType returns the Terraform awscc_customerprofiles_integration data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CustomerProfiles::Integration resource type.
func integrationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got created",
			//   "type": "string"
			// }
			Description: "The time of this integration got created",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the domain.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9_-]+$",
			//   "type": "string"
			// }
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"flow_definition": {
			// Property: FlowDefinition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Description": {
			//       "maxLength": 2048,
			//       "pattern": "[\\w!@#\\-.?,\\s]*",
			//       "type": "string"
			//     },
			//     "FlowName": {
			//       "maxLength": 256,
			//       "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
			//       "type": "string"
			//     },
			//     "KmsArn": {
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "arn:aws:kms:.*:[0-9]+:.*",
			//       "type": "string"
			//     },
			//     "SourceFlowConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ConnectorProfileName": {
			//           "maxLength": 256,
			//           "pattern": "[\\w/!@#+=.-]+",
			//           "type": "string"
			//         },
			//         "ConnectorType": {
			//           "enum": [
			//             "Salesforce",
			//             "Marketo",
			//             "ServiceNow",
			//             "Zendesk",
			//             "S3"
			//           ],
			//           "type": "string"
			//         },
			//         "IncrementalPullConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "DatetimeTypeFieldName": {
			//               "maxLength": 256,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SourceConnectorProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Marketo": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "\\S+",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "S3": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "BucketName": {
			//                   "maxLength": 63,
			//                   "minLength": 3,
			//                   "pattern": "\\S+",
			//                   "type": "string"
			//                 },
			//                 "BucketPrefix": {
			//                   "maxLength": 512,
			//                   "pattern": ".*",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "BucketName"
			//               ],
			//               "type": "object"
			//             },
			//             "Salesforce": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "EnableDynamicFieldUpdate": {
			//                   "type": "boolean"
			//                 },
			//                 "IncludeDeletedRecords": {
			//                   "type": "boolean"
			//                 },
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "\\S+",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "ServiceNow": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "\\S+",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             },
			//             "Zendesk": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Object": {
			//                   "additionalProperties": false,
			//                   "maxLength": 512,
			//                   "pattern": "\\S+",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Object"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "ConnectorType",
			//         "SourceConnectorProperties"
			//       ],
			//       "type": "object"
			//     },
			//     "Tasks": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ConnectorOperator": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Marketo": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "S3": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Salesforce": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "ServiceNow": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "LESS_THAN",
			//                   "GREATER_THAN",
			//                   "CONTAINS",
			//                   "BETWEEN",
			//                   "LESS_THAN_OR_EQUAL_TO",
			//                   "GREATER_THAN_OR_EQUAL_TO",
			//                   "EQUAL_TO",
			//                   "NOT_EQUAL_TO",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Zendesk": {
			//                 "enum": [
			//                   "PROJECTION",
			//                   "GREATER_THAN",
			//                   "ADDITION",
			//                   "MULTIPLICATION",
			//                   "DIVISION",
			//                   "SUBTRACTION",
			//                   "MASK_ALL",
			//                   "MASK_FIRST_N",
			//                   "MASK_LAST_N",
			//                   "VALIDATE_NON_NULL",
			//                   "VALIDATE_NON_ZERO",
			//                   "VALIDATE_NON_NEGATIVE",
			//                   "VALIDATE_NUMERIC",
			//                   "NO_OP"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "DestinationField": {
			//             "maxLength": 256,
			//             "pattern": ".*",
			//             "type": "string"
			//           },
			//           "SourceFields": {
			//             "items": {
			//               "maxLength": 2048,
			//               "pattern": ".*",
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "TaskProperties": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "OperatorPropertyKey": {
			//                   "enum": [
			//                     "VALUE",
			//                     "VALUES",
			//                     "DATA_TYPE",
			//                     "UPPER_BOUND",
			//                     "LOWER_BOUND",
			//                     "SOURCE_DATA_TYPE",
			//                     "DESTINATION_DATA_TYPE",
			//                     "VALIDATION_ACTION",
			//                     "MASK_VALUE",
			//                     "MASK_LENGTH",
			//                     "TRUNCATE_LENGTH",
			//                     "MATH_OPERATION_FIELDS_ORDER",
			//                     "CONCAT_FORMAT",
			//                     "SUBFIELD_CATEGORY_MAP"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Property": {
			//                   "maxLength": 2048,
			//                   "pattern": ".+",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "OperatorPropertyKey",
			//                 "Property"
			//               ],
			//               "type": "object"
			//             },
			//             "type": "array"
			//           },
			//           "TaskType": {
			//             "enum": [
			//               "Arithmetic",
			//               "Filter",
			//               "Map",
			//               "Mask",
			//               "Merge",
			//               "Truncate",
			//               "Validate"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "SourceFields",
			//           "TaskType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "TriggerConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "TriggerProperties": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Scheduled": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "DataPullMode": {
			//                   "enum": [
			//                     "Incremental",
			//                     "Complete"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "FirstExecutionFrom": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleEndTime": {
			//                   "type": "number"
			//                 },
			//                 "ScheduleExpression": {
			//                   "maxLength": 256,
			//                   "pattern": ".*",
			//                   "type": "string"
			//                 },
			//                 "ScheduleOffset": {
			//                   "maximum": 36000,
			//                   "minimum": 0,
			//                   "type": "integer"
			//                 },
			//                 "ScheduleStartTime": {
			//                   "type": "number"
			//                 },
			//                 "Timezone": {
			//                   "maxLength": 256,
			//                   "pattern": ".*",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "ScheduleExpression"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "TriggerType": {
			//           "enum": [
			//             "Scheduled",
			//             "Event",
			//             "OnDemand"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "TriggerType"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "FlowName",
			//     "KmsArn",
			//     "Tasks",
			//     "TriggerConfig",
			//     "SourceFlowConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Computed: true,
					},
					"flow_name": {
						// Property: FlowName
						Type:     types.StringType,
						Computed: true,
					},
					"kms_arn": {
						// Property: KmsArn
						Type:     types.StringType,
						Computed: true,
					},
					"source_flow_config": {
						// Property: SourceFlowConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"connector_profile_name": {
									// Property: ConnectorProfileName
									Type:     types.StringType,
									Computed: true,
								},
								"connector_type": {
									// Property: ConnectorType
									Type:     types.StringType,
									Computed: true,
								},
								"incremental_pull_config": {
									// Property: IncrementalPullConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"datetime_type_field_name": {
												// Property: DatetimeTypeFieldName
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"source_connector_properties": {
									// Property: SourceConnectorProperties
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"marketo": {
												// Property: Marketo
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"s3": {
												// Property: S3
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"bucket_name": {
															// Property: BucketName
															Type:     types.StringType,
															Computed: true,
														},
														"bucket_prefix": {
															// Property: BucketPrefix
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"salesforce": {
												// Property: Salesforce
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"enable_dynamic_field_update": {
															// Property: EnableDynamicFieldUpdate
															Type:     types.BoolType,
															Computed: true,
														},
														"include_deleted_records": {
															// Property: IncludeDeletedRecords
															Type:     types.BoolType,
															Computed: true,
														},
														"object": {
															// Property: Object
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"service_now": {
												// Property: ServiceNow
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"zendesk": {
												// Property: Zendesk
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"object": {
															// Property: Object
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
					"tasks": {
						// Property: Tasks
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"connector_operator": {
									// Property: ConnectorOperator
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"marketo": {
												// Property: Marketo
												Type:     types.StringType,
												Computed: true,
											},
											"s3": {
												// Property: S3
												Type:     types.StringType,
												Computed: true,
											},
											"salesforce": {
												// Property: Salesforce
												Type:     types.StringType,
												Computed: true,
											},
											"service_now": {
												// Property: ServiceNow
												Type:     types.StringType,
												Computed: true,
											},
											"zendesk": {
												// Property: Zendesk
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"destination_field": {
									// Property: DestinationField
									Type:     types.StringType,
									Computed: true,
								},
								"source_fields": {
									// Property: SourceFields
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"task_properties": {
									// Property: TaskProperties
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"operator_property_key": {
												// Property: OperatorPropertyKey
												Type:     types.StringType,
												Computed: true,
											},
											"property": {
												// Property: Property
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"task_type": {
									// Property: TaskType
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"trigger_config": {
						// Property: TriggerConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"trigger_properties": {
									// Property: TriggerProperties
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"scheduled": {
												// Property: Scheduled
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"data_pull_mode": {
															// Property: DataPullMode
															Type:     types.StringType,
															Computed: true,
														},
														"first_execution_from": {
															// Property: FirstExecutionFrom
															Type:     types.Float64Type,
															Computed: true,
														},
														"schedule_end_time": {
															// Property: ScheduleEndTime
															Type:     types.Float64Type,
															Computed: true,
														},
														"schedule_expression": {
															// Property: ScheduleExpression
															Type:     types.StringType,
															Computed: true,
														},
														"schedule_offset": {
															// Property: ScheduleOffset
															Type:     types.Int64Type,
															Computed: true,
														},
														"schedule_start_time": {
															// Property: ScheduleStartTime
															Type:     types.Float64Type,
															Computed: true,
														},
														"timezone": {
															// Property: Timezone
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
								"trigger_type": {
									// Property: TriggerType
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
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got last updated at",
			//   "type": "string"
			// }
			Description: "The time of this integration got last updated at",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_type_name": {
			// Property: ObjectTypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the ObjectType defined for the 3rd party data in Profile Service",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
			//   "type": "string"
			// }
			Description: "The name of the ObjectType defined for the 3rd party data in Profile Service",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_type_names": {
			// Property: ObjectTypeNames
			// CloudFormation resource type schema:
			// {
			//   "description": "The mapping between 3rd party event types and ObjectType names",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The mapping between 3rd party event types and ObjectType names",
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
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the integration",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
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
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the integration",
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
		"uri": {
			// Property: Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URI of the S3 bucket or any other type of data source.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The URI of the S3 bucket or any other type of data source.",
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
		Description: "Data Source schema for AWS::CustomerProfiles::Integration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Integration").WithTerraformTypeName("awscc_customerprofiles_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":                 "BucketName",
		"bucket_prefix":               "BucketPrefix",
		"connector_operator":          "ConnectorOperator",
		"connector_profile_name":      "ConnectorProfileName",
		"connector_type":              "ConnectorType",
		"created_at":                  "CreatedAt",
		"data_pull_mode":              "DataPullMode",
		"datetime_type_field_name":    "DatetimeTypeFieldName",
		"description":                 "Description",
		"destination_field":           "DestinationField",
		"domain_name":                 "DomainName",
		"enable_dynamic_field_update": "EnableDynamicFieldUpdate",
		"first_execution_from":        "FirstExecutionFrom",
		"flow_definition":             "FlowDefinition",
		"flow_name":                   "FlowName",
		"include_deleted_records":     "IncludeDeletedRecords",
		"incremental_pull_config":     "IncrementalPullConfig",
		"key":                         "Key",
		"kms_arn":                     "KmsArn",
		"last_updated_at":             "LastUpdatedAt",
		"marketo":                     "Marketo",
		"object":                      "Object",
		"object_type_name":            "ObjectTypeName",
		"object_type_names":           "ObjectTypeNames",
		"operator_property_key":       "OperatorPropertyKey",
		"property":                    "Property",
		"s3":                          "S3",
		"salesforce":                  "Salesforce",
		"schedule_end_time":           "ScheduleEndTime",
		"schedule_expression":         "ScheduleExpression",
		"schedule_offset":             "ScheduleOffset",
		"schedule_start_time":         "ScheduleStartTime",
		"scheduled":                   "Scheduled",
		"service_now":                 "ServiceNow",
		"source_connector_properties": "SourceConnectorProperties",
		"source_fields":               "SourceFields",
		"source_flow_config":          "SourceFlowConfig",
		"tags":                        "Tags",
		"task_properties":             "TaskProperties",
		"task_type":                   "TaskType",
		"tasks":                       "Tasks",
		"timezone":                    "Timezone",
		"trigger_config":              "TriggerConfig",
		"trigger_properties":          "TriggerProperties",
		"trigger_type":                "TriggerType",
		"uri":                         "Uri",
		"value":                       "Value",
		"zendesk":                     "Zendesk",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
