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
	registry.AddDataSourceTypeFactory("awscc_athena_named_query", namedQueryDataSourceType)
}

// namedQueryDataSourceType returns the Terraform awscc_athena_named_query data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Athena::NamedQuery resource type.
func namedQueryDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"database": {
			// Property: Database
			// CloudFormation resource type schema:
			// {
			//   "description": "The database to which the query belongs.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The database to which the query belongs.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The query description.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The query name.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"named_query_id": {
			// Property: NamedQueryId
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique ID of the query.",
			//   "type": "string"
			// }
			Description: "The unique ID of the query.",
			Type:        types.StringType,
			Computed:    true,
		},
		"query_string": {
			// Property: QueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The contents of the query with all query statements.",
			//   "maxLength": 262144,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The contents of the query with all query statements.",
			Type:        types.StringType,
			Computed:    true,
		},
		"work_group": {
			// Property: WorkGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the workgroup that contains the named query.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the workgroup that contains the named query.",
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
		Description: "Data Source schema for AWS::Athena::NamedQuery",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::NamedQuery").WithTerraformTypeName("awscc_athena_named_query")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"database":       "Database",
		"description":    "Description",
		"name":           "Name",
		"named_query_id": "NamedQueryId",
		"query_string":   "QueryString",
		"work_group":     "WorkGroup",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
