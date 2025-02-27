// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3objectlambda

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
	registry.AddResourceTypeFactory("awscc_s3objectlambda_access_point_policy", accessPointPolicyResourceType)
}

// accessPointPolicyResourceType returns the Terraform awscc_s3objectlambda_access_point_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPointPolicy resource type.
func accessPointPolicyResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"object_lambda_access_point": {
			// Property: ObjectLambdaAccessPoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
			//   "maxLength": 45,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
			//   "type": "string"
			// }
			Description: "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 45),
				validate.StringMatch(regexp.MustCompile("^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
			//   "type": "object"
			// }
			Description: "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
			Type:        types.MapType{ElemType: types.StringType},
			Required:    true,
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
		Description: "AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPointPolicy").WithTerraformTypeName("awscc_s3objectlambda_access_point_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"object_lambda_access_point": "ObjectLambdaAccessPoint",
		"policy_document":            "PolicyDocument",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
