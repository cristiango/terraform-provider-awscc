// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lakeformation_data_cells_filter", dataCellsFilterDataSourceType)
}

// dataCellsFilterDataSourceType returns the Terraform awscc_lakeformation_data_cells_filter data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LakeFormation::DataCellsFilter resource type.
func dataCellsFilterDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"column_names": {
			// Property: ColumnNames
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of columns to be included in this Data Cells Filter.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "A string representing a resource's name.",
			//     "maxLength": 255,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of columns to be included in this Data Cells Filter.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"column_wildcard": {
			// Property: ColumnWildcard
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required",
			//   "properties": {
			//     "ExcludedColumnNames": {
			//       "description": "A list of column names to be excluded from the Data Cells Filter.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "A string representing a resource's name.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"excluded_column_names": {
						// Property: ExcludedColumnNames
						Description: "A list of column names to be excluded from the Data Cells Filter.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"database_name": {
			// Property: DatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Database that the Table resides in.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Database that the Table resides in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The desired name of the Data Cells Filter.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The desired name of the Data Cells Filter.",
			Type:        types.StringType,
			Computed:    true,
		},
		"row_filter": {
			// Property: RowFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required",
			//   "properties": {
			//     "AllRowsWildcard": {
			//       "additionalProperties": false,
			//       "description": "An empty object representing a row wildcard.",
			//       "type": "object"
			//     },
			//     "FilterExpression": {
			//       "description": "A PartiQL predicate.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"all_rows_wildcard": {
						// Property: AllRowsWildcard
						Description: "An empty object representing a row wildcard.",
						Type:        types.MapType{ElemType: types.StringType},
						Computed:    true,
					},
					"filter_expression": {
						// Property: FilterExpression
						Description: "A PartiQL predicate.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"table_catalog_id": {
			// Property: TableCatalogId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Catalog Id of the Table on which to create a Data Cells Filter.",
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "The Catalog Id of the Table on which to create a Data Cells Filter.",
			Type:        types.StringType,
			Computed:    true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Table to create a Data Cells Filter for.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Table to create a Data Cells Filter for.",
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
		Description: "Data Source schema for AWS::LakeFormation::DataCellsFilter",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::DataCellsFilter").WithTerraformTypeName("awscc_lakeformation_data_cells_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"all_rows_wildcard":     "AllRowsWildcard",
		"column_names":          "ColumnNames",
		"column_wildcard":       "ColumnWildcard",
		"database_name":         "DatabaseName",
		"excluded_column_names": "ExcludedColumnNames",
		"filter_expression":     "FilterExpression",
		"name":                  "Name",
		"row_filter":            "RowFilter",
		"table_catalog_id":      "TableCatalogId",
		"table_name":            "TableName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
