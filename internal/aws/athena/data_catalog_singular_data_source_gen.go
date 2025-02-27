// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_athena_data_catalog", dataCatalogDataSourceType)
}

// dataCatalogDataSourceType returns the Terraform awscc_athena_data_catalog data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Athena::DataCatalog resource type.
func dataCatalogDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the data catalog to be created. ",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A description of the data catalog to be created. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. ",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type. ",
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 51200,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type. ",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of comma separated tags to add to the data catalog that is created. ",
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
			//   "type": "array"
			// }
			Description: "A list of comma separated tags to add to the data catalog that is created. ",
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
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. ",
			//   "enum": [
			//     "LAMBDA",
			//     "GLUE",
			//     "HIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. ",
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
		Description: "Data Source schema for AWS::Athena::DataCatalog",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::DataCatalog").WithTerraformTypeName("awscc_athena_data_catalog")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description": "Description",
		"key":         "Key",
		"name":        "Name",
		"parameters":  "Parameters",
		"tags":        "Tags",
		"type":        "Type",
		"value":       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
