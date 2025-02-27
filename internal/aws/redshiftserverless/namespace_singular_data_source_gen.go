// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_redshiftserverless_namespace", namespaceDataSourceType)
}

// namespaceDataSourceType returns the Terraform awscc_redshiftserverless_namespace data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::RedshiftServerless::Namespace resource type.
func namespaceDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"admin_user_password": {
			// Property: AdminUserPassword
			// CloudFormation resource type schema:
			// {
			//   "description": "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit.",
			//   "maxLength": 64,
			//   "minLength": 8,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit.",
			Type:        types.StringType,
			Computed:    true,
		},
		"admin_username": {
			// Property: AdminUsername
			// CloudFormation resource type schema:
			// {
			//   "description": "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			//   "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
			//   "type": "string"
			// }
			Description: "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"db_name": {
			// Property: DbName
			// CloudFormation resource type schema:
			// {
			//   "description": "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			//   "maxLength": 127,
			//   "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
			//   "type": "string"
			// }
			Description: "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_iam_role_arn": {
			// Property: DefaultIamRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The default IAM role ARN for the namespace that is being created.",
			//   "type": "string"
			// }
			Description: "The default IAM role ARN for the namespace that is being created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"final_snapshot_name": {
			// Property: FinalSnapshotName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
			//   "maxLength": 255,
			//   "pattern": "[a-z][a-z0-9]*(-[a-z0-9]+)*",
			//   "type": "string"
			// }
			Description: "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
			Type:        types.StringType,
			Computed:    true,
		},
		"final_snapshot_retention_period": {
			// Property: FinalSnapshotRetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
			//   "type": "integer"
			// }
			Description: "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"iam_roles": {
			// Property: IamRoles
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 512,
			//     "minLength": 0,
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
			//   "type": "string"
			// }
			Description: "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_exports": {
			// Property: LogExports
			// CloudFormation resource type schema:
			// {
			//   "description": "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
			//   "insertionOrder": false,
			//   "items": {
			//     "enum": [
			//       "useractivitylog",
			//       "userlog",
			//       "connectionlog"
			//     ],
			//     "type": "string"
			//   },
			//   "maxItems": 16,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"namespace": {
			// Property: Namespace
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AdminUsername": {
			//       "type": "string"
			//     },
			//     "CreationDate": {
			//       "type": "string"
			//     },
			//     "DbName": {
			//       "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
			//       "type": "string"
			//     },
			//     "DefaultIamRoleArn": {
			//       "type": "string"
			//     },
			//     "IamRoles": {
			//       "insertionOrder": false,
			//       "items": {
			//         "maxLength": 512,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "KmsKeyId": {
			//       "type": "string"
			//     },
			//     "LogExports": {
			//       "insertionOrder": false,
			//       "items": {
			//         "enum": [
			//           "useractivitylog",
			//           "userlog",
			//           "connectionlog"
			//         ],
			//         "type": "string"
			//       },
			//       "maxItems": 16,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "NamespaceArn": {
			//       "type": "string"
			//     },
			//     "NamespaceId": {
			//       "type": "string"
			//     },
			//     "NamespaceName": {
			//       "maxLength": 64,
			//       "minLength": 3,
			//       "pattern": "^[a-z0-9-]+$",
			//       "type": "string"
			//     },
			//     "Status": {
			//       "enum": [
			//         "AVAILABLE",
			//         "MODIFYING",
			//         "DELETING"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"admin_username": {
						// Property: AdminUsername
						Type:     types.StringType,
						Computed: true,
					},
					"creation_date": {
						// Property: CreationDate
						Type:     types.StringType,
						Computed: true,
					},
					"db_name": {
						// Property: DbName
						Type:     types.StringType,
						Computed: true,
					},
					"default_iam_role_arn": {
						// Property: DefaultIamRoleArn
						Type:     types.StringType,
						Computed: true,
					},
					"iam_roles": {
						// Property: IamRoles
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"kms_key_id": {
						// Property: KmsKeyId
						Type:     types.StringType,
						Computed: true,
					},
					"log_exports": {
						// Property: LogExports
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"namespace_arn": {
						// Property: NamespaceArn
						Type:     types.StringType,
						Computed: true,
					},
					"namespace_id": {
						// Property: NamespaceId
						Type:     types.StringType,
						Computed: true,
					},
					"namespace_name": {
						// Property: NamespaceName
						Type:     types.StringType,
						Computed: true,
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"namespace_name": {
			// Property: NamespaceName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
			//   "maxLength": 64,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9-]+$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of tags for the namespace.",
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
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The list of tags for the namespace.",
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
		Description: "Data Source schema for AWS::RedshiftServerless::Namespace",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Namespace").WithTerraformTypeName("awscc_redshiftserverless_namespace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_user_password":             "AdminUserPassword",
		"admin_username":                  "AdminUsername",
		"creation_date":                   "CreationDate",
		"db_name":                         "DbName",
		"default_iam_role_arn":            "DefaultIamRoleArn",
		"final_snapshot_name":             "FinalSnapshotName",
		"final_snapshot_retention_period": "FinalSnapshotRetentionPeriod",
		"iam_roles":                       "IamRoles",
		"key":                             "Key",
		"kms_key_id":                      "KmsKeyId",
		"log_exports":                     "LogExports",
		"namespace":                       "Namespace",
		"namespace_arn":                   "NamespaceArn",
		"namespace_id":                    "NamespaceId",
		"namespace_name":                  "NamespaceName",
		"status":                          "Status",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
