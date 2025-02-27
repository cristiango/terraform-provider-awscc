// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_s3outposts_bucket_policy", bucketPolicyDataSourceType)
}

// bucketPolicyDataSourceType returns the Terraform awscc_s3outposts_bucket_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::S3Outposts::BucketPolicy resource type.
func bucketPolicyDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified bucket.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/bucket\\/[^:]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "A policy document containing permissions to add to the specified bucket.",
			//   "type": "object"
			// }
			Description: "A policy document containing permissions to add to the specified bucket.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3Outposts::BucketPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::BucketPolicy").WithTerraformTypeName("awscc_s3outposts_bucket_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":          "Bucket",
		"policy_document": "PolicyDocument",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
