// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kafkaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kafkaconnect_connector", connectorDataSourceType)
}

// connectorDataSourceType returns the Terraform awscc_kafkaconnect_connector data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::KafkaConnect::Connector resource type.
func connectorDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity": {
			// Property: Capacity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the capacity allocated to the connector.",
			//   "oneOf": [
			//     {
			//       "required": [
			//         "AutoScaling"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "ProvisionedCapacity"
			//       ]
			//     }
			//   ],
			//   "properties": {
			//     "AutoScaling": {
			//       "additionalProperties": false,
			//       "description": "Details about auto scaling of a connector. ",
			//       "properties": {
			//         "MaxWorkerCount": {
			//           "description": "The maximum number of workers for a connector.",
			//           "type": "integer"
			//         },
			//         "McuCount": {
			//           "description": "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
			//           "enum": [
			//             1,
			//             2,
			//             4,
			//             8
			//           ],
			//           "type": "integer"
			//         },
			//         "MinWorkerCount": {
			//           "description": "The minimum number of workers for a connector.",
			//           "type": "integer"
			//         },
			//         "ScaleInPolicy": {
			//           "additionalProperties": false,
			//           "description": "Information about the scale in policy of the connector.",
			//           "properties": {
			//             "CpuUtilizationPercentage": {
			//               "description": "Specifies the CPU utilization percentage threshold at which connector scale in should trigger.",
			//               "maximum": 100,
			//               "minimum": 1,
			//               "type": "integer"
			//             }
			//           },
			//           "required": [
			//             "CpuUtilizationPercentage"
			//           ],
			//           "type": "object"
			//         },
			//         "ScaleOutPolicy": {
			//           "additionalProperties": false,
			//           "description": "Information about the scale out policy of the connector.",
			//           "properties": {
			//             "CpuUtilizationPercentage": {
			//               "description": "Specifies the CPU utilization percentage threshold at which connector scale out should trigger.",
			//               "maximum": 100,
			//               "minimum": 1,
			//               "type": "integer"
			//             }
			//           },
			//           "required": [
			//             "CpuUtilizationPercentage"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "MaxWorkerCount",
			//         "MinWorkerCount",
			//         "ScaleInPolicy",
			//         "ScaleOutPolicy",
			//         "McuCount"
			//       ],
			//       "type": "object"
			//     },
			//     "ProvisionedCapacity": {
			//       "additionalProperties": false,
			//       "description": "Details about a fixed capacity allocated to a connector.",
			//       "properties": {
			//         "McuCount": {
			//           "description": "Specifies how many MSK Connect Units (MCU) are allocated to the connector.",
			//           "enum": [
			//             1,
			//             2,
			//             4,
			//             8
			//           ],
			//           "type": "integer"
			//         },
			//         "WorkerCount": {
			//           "description": "Number of workers for a connector.",
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "WorkerCount"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about the capacity allocated to the connector.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_scaling": {
						// Property: AutoScaling
						Description: "Details about auto scaling of a connector. ",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"max_worker_count": {
									// Property: MaxWorkerCount
									Description: "The maximum number of workers for a connector.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"mcu_count": {
									// Property: McuCount
									Description: "Specifies how many MSK Connect Units (MCU) as the minimum scaling unit.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"min_worker_count": {
									// Property: MinWorkerCount
									Description: "The minimum number of workers for a connector.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"scale_in_policy": {
									// Property: ScaleInPolicy
									Description: "Information about the scale in policy of the connector.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cpu_utilization_percentage": {
												// Property: CpuUtilizationPercentage
												Description: "Specifies the CPU utilization percentage threshold at which connector scale in should trigger.",
												Type:        types.Int64Type,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"scale_out_policy": {
									// Property: ScaleOutPolicy
									Description: "Information about the scale out policy of the connector.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cpu_utilization_percentage": {
												// Property: CpuUtilizationPercentage
												Description: "Specifies the CPU utilization percentage threshold at which connector scale out should trigger.",
												Type:        types.Int64Type,
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
					"provisioned_capacity": {
						// Property: ProvisionedCapacity
						Description: "Details about a fixed capacity allocated to a connector.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"mcu_count": {
									// Property: McuCount
									Description: "Specifies how many MSK Connect Units (MCU) are allocated to the connector.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"worker_count": {
									// Property: WorkerCount
									Description: "Number of workers for a connector.",
									Type:        types.Int64Type,
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
		"connector_arn": {
			// Property: ConnectorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name for the created Connector.",
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name for the created Connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connector_configuration": {
			// Property: ConnectorConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration for the connector.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration for the connector.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"connector_description": {
			// Property: ConnectorDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A summary description of the connector.",
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Description: "A summary description of the connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connector_name": {
			// Property: ConnectorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the connector.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kafka_cluster": {
			// Property: KafkaCluster
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Details of how to connect to the Kafka cluster.",
			//   "properties": {
			//     "ApacheKafkaCluster": {
			//       "additionalProperties": false,
			//       "description": "Details of how to connect to an Apache Kafka cluster.",
			//       "properties": {
			//         "BootstrapServers": {
			//           "description": "The bootstrap servers string of the Apache Kafka cluster.",
			//           "type": "string"
			//         },
			//         "Vpc": {
			//           "additionalProperties": false,
			//           "description": "Information about a VPC used with the connector.",
			//           "properties": {
			//             "SecurityGroups": {
			//               "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "Subnets": {
			//               "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "SecurityGroups",
			//             "Subnets"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "BootstrapServers",
			//         "Vpc"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "ApacheKafkaCluster"
			//   ],
			//   "type": "object"
			// }
			Description: "Details of how to connect to the Kafka cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"apache_kafka_cluster": {
						// Property: ApacheKafkaCluster
						Description: "Details of how to connect to an Apache Kafka cluster.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bootstrap_servers": {
									// Property: BootstrapServers
									Description: "The bootstrap servers string of the Apache Kafka cluster.",
									Type:        types.StringType,
									Computed:    true,
								},
								"vpc": {
									// Property: Vpc
									Description: "Information about a VPC used with the connector.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"security_groups": {
												// Property: SecurityGroups
												Description: "The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.",
												Type:        types.SetType{ElemType: types.StringType},
												Computed:    true,
											},
											"subnets": {
												// Property: Subnets
												Description: "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
												Type:        types.SetType{ElemType: types.StringType},
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
				},
			),
			Computed: true,
		},
		"kafka_cluster_client_authentication": {
			// Property: KafkaClusterClientAuthentication
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Details of the client authentication used by the Kafka cluster.",
			//   "properties": {
			//     "AuthenticationType": {
			//       "description": "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
			//       "enum": [
			//         "NONE",
			//         "IAM"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AuthenticationType"
			//   ],
			//   "type": "object"
			// }
			Description: "Details of the client authentication used by the Kafka cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"authentication_type": {
						// Property: AuthenticationType
						Description: "The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"kafka_cluster_encryption_in_transit": {
			// Property: KafkaClusterEncryptionInTransit
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Details of encryption in transit to the Kafka cluster.",
			//   "properties": {
			//     "EncryptionType": {
			//       "description": "The type of encryption in transit to the Kafka cluster.",
			//       "enum": [
			//         "PLAINTEXT",
			//         "TLS"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "EncryptionType"
			//   ],
			//   "type": "object"
			// }
			Description: "Details of encryption in transit to the Kafka cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "The type of encryption in transit to the Kafka cluster.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"kafka_connect_version": {
			// Property: KafkaConnectVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
			//   "type": "string"
			// }
			Description: "The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_delivery": {
			// Property: LogDelivery
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Details of what logs are delivered and where they are delivered.",
			//   "properties": {
			//     "WorkerLogDelivery": {
			//       "additionalProperties": false,
			//       "description": "Specifies where worker logs are delivered.",
			//       "properties": {
			//         "CloudWatchLogs": {
			//           "additionalProperties": false,
			//           "description": "Details about delivering logs to Amazon CloudWatch Logs.",
			//           "properties": {
			//             "Enabled": {
			//               "description": "Specifies whether the logs get sent to the specified CloudWatch Logs destination.",
			//               "type": "boolean"
			//             },
			//             "LogGroup": {
			//               "description": "The CloudWatch log group that is the destination for log delivery.",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         },
			//         "Firehose": {
			//           "additionalProperties": false,
			//           "description": "Details about delivering logs to Amazon Kinesis Data Firehose.",
			//           "properties": {
			//             "DeliveryStream": {
			//               "description": "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
			//               "type": "string"
			//             },
			//             "Enabled": {
			//               "description": "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         },
			//         "S3": {
			//           "additionalProperties": false,
			//           "description": "Details about delivering logs to Amazon S3.",
			//           "properties": {
			//             "Bucket": {
			//               "description": "The name of the S3 bucket that is the destination for log delivery.",
			//               "type": "string"
			//             },
			//             "Enabled": {
			//               "description": "Specifies whether the logs get sent to the specified Amazon S3 destination.",
			//               "type": "boolean"
			//             },
			//             "Prefix": {
			//               "description": "The S3 prefix that is the destination for log delivery.",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "WorkerLogDelivery"
			//   ],
			//   "type": "object"
			// }
			Description: "Details of what logs are delivered and where they are delivered.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"worker_log_delivery": {
						// Property: WorkerLogDelivery
						Description: "Specifies where worker logs are delivered.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cloudwatch_logs": {
									// Property: CloudWatchLogs
									Description: "Details about delivering logs to Amazon CloudWatch Logs.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Description: "Specifies whether the logs get sent to the specified CloudWatch Logs destination.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"log_group": {
												// Property: LogGroup
												Description: "The CloudWatch log group that is the destination for log delivery.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"firehose": {
									// Property: Firehose
									Description: "Details about delivering logs to Amazon Kinesis Data Firehose.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delivery_stream": {
												// Property: DeliveryStream
												Description: "The Kinesis Data Firehose delivery stream that is the destination for log delivery.",
												Type:        types.StringType,
												Computed:    true,
											},
											"enabled": {
												// Property: Enabled
												Description: "Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.",
												Type:        types.BoolType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"s3": {
									// Property: S3
									Description: "Details about delivering logs to Amazon S3.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Description: "The name of the S3 bucket that is the destination for log delivery.",
												Type:        types.StringType,
												Computed:    true,
											},
											"enabled": {
												// Property: Enabled
												Description: "Specifies whether the logs get sent to the specified Amazon S3 destination.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"prefix": {
												// Property: Prefix
												Description: "The S3 prefix that is the destination for log delivery.",
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
				},
			),
			Computed: true,
		},
		"plugins": {
			// Property: Plugins
			// CloudFormation resource type schema:
			// {
			//   "description": "List of plugins to use with the connector.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Details about a Kafka Connect plugin which will be used with the connector.",
			//     "properties": {
			//       "CustomPlugin": {
			//         "additionalProperties": false,
			//         "description": "Details about a custom plugin.",
			//         "properties": {
			//           "CustomPluginArn": {
			//             "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
			//             "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
			//             "type": "string"
			//           },
			//           "Revision": {
			//             "description": "The revision of the custom plugin to use.",
			//             "format": "int64",
			//             "minimum": 1,
			//             "type": "integer"
			//           }
			//         },
			//         "required": [
			//           "CustomPluginArn",
			//           "Revision"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "CustomPlugin"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "List of plugins to use with the connector.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"custom_plugin": {
						// Property: CustomPlugin
						Description: "Details about a custom plugin.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"custom_plugin_arn": {
									// Property: CustomPluginArn
									Description: "The Amazon Resource Name (ARN) of the custom plugin to use.",
									Type:        types.StringType,
									Computed:    true,
								},
								"revision": {
									// Property: Revision
									Description: "The revision of the custom plugin to use.",
									Type:        types.Int64Type,
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
		"service_execution_role_arn": {
			// Property: ServiceExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn):iam:.*",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.",
			Type:        types.StringType,
			Computed:    true,
		},
		"worker_configuration": {
			// Property: WorkerConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies the worker configuration to use with the connector.",
			//   "properties": {
			//     "Revision": {
			//       "description": "The revision of the worker configuration to use.",
			//       "format": "int64",
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "WorkerConfigurationArn": {
			//       "description": "The Amazon Resource Name (ARN) of the worker configuration to use.",
			//       "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Revision",
			//     "WorkerConfigurationArn"
			//   ],
			//   "type": "object"
			// }
			Description: "Specifies the worker configuration to use with the connector.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"revision": {
						// Property: Revision
						Description: "The revision of the worker configuration to use.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"worker_configuration_arn": {
						// Property: WorkerConfigurationArn
						Description: "The Amazon Resource Name (ARN) of the worker configuration to use.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::KafkaConnect::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KafkaConnect::Connector").WithTerraformTypeName("awscc_kafkaconnect_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"apache_kafka_cluster":                "ApacheKafkaCluster",
		"authentication_type":                 "AuthenticationType",
		"auto_scaling":                        "AutoScaling",
		"bootstrap_servers":                   "BootstrapServers",
		"bucket":                              "Bucket",
		"capacity":                            "Capacity",
		"cloudwatch_logs":                     "CloudWatchLogs",
		"connector_arn":                       "ConnectorArn",
		"connector_configuration":             "ConnectorConfiguration",
		"connector_description":               "ConnectorDescription",
		"connector_name":                      "ConnectorName",
		"cpu_utilization_percentage":          "CpuUtilizationPercentage",
		"custom_plugin":                       "CustomPlugin",
		"custom_plugin_arn":                   "CustomPluginArn",
		"delivery_stream":                     "DeliveryStream",
		"enabled":                             "Enabled",
		"encryption_type":                     "EncryptionType",
		"firehose":                            "Firehose",
		"kafka_cluster":                       "KafkaCluster",
		"kafka_cluster_client_authentication": "KafkaClusterClientAuthentication",
		"kafka_cluster_encryption_in_transit": "KafkaClusterEncryptionInTransit",
		"kafka_connect_version":               "KafkaConnectVersion",
		"log_delivery":                        "LogDelivery",
		"log_group":                           "LogGroup",
		"max_worker_count":                    "MaxWorkerCount",
		"mcu_count":                           "McuCount",
		"min_worker_count":                    "MinWorkerCount",
		"plugins":                             "Plugins",
		"prefix":                              "Prefix",
		"provisioned_capacity":                "ProvisionedCapacity",
		"revision":                            "Revision",
		"s3":                                  "S3",
		"scale_in_policy":                     "ScaleInPolicy",
		"scale_out_policy":                    "ScaleOutPolicy",
		"security_groups":                     "SecurityGroups",
		"service_execution_role_arn":          "ServiceExecutionRoleArn",
		"subnets":                             "Subnets",
		"vpc":                                 "Vpc",
		"worker_configuration":                "WorkerConfiguration",
		"worker_configuration_arn":            "WorkerConfigurationArn",
		"worker_count":                        "WorkerCount",
		"worker_log_delivery":                 "WorkerLogDelivery",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
