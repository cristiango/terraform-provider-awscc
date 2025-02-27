// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalogappregistry

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
	registry.AddResourceTypeFactory("awscc_servicecatalogappregistry_attribute_group_association", attributeGroupAssociationResourceType)
}

// attributeGroupAssociationResourceType returns the Terraform awscc_servicecatalogappregistry_attribute_group_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation resource type.
func attributeGroupAssociationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application": {
			// Property: Application
			// CloudFormation resource type schema:
			// {
			//   "description": "The name or the Id of the Application.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "\\w+|[a-z0-9]{12}",
			//   "type": "string"
			// }
			Description: "The name or the Id of the Application.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("\\w+|[a-z0-9]{12}"), ""),
			},
		},
		"application_arn": {
			// Property: ApplicationArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/applications/[a-z0-9]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"attribute_group": {
			// Property: AttributeGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "The name or the Id of the AttributeGroup.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "\\w+|[a-z0-9]{12}",
			//   "type": "string"
			// }
			Description: "The name or the Id of the AttributeGroup.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("\\w+|[a-z0-9]{12}"), ""),
			},
		},
		"attribute_group_arn": {
			// Property: AttributeGroupArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/attribute-groups/[a-z0-9]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
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
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation").WithTerraformTypeName("awscc_servicecatalogappregistry_attribute_group_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application":         "Application",
		"application_arn":     "ApplicationArn",
		"attribute_group":     "AttributeGroup",
		"attribute_group_arn": "AttributeGroupArn",
		"id":                  "Id",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
