// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_imagebuilder_component", componentResourceType)
}

// componentResourceType returns the Terraform awscc_imagebuilder_component resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::Component resource type.
func componentResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the component.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the component.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"change_description": {
			// Property: ChangeDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The change description of the component.",
			//   "type": "string"
			// }
			Description: "The change description of the component.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"data": {
			// Property: Data
			// CloudFormation resource type schema:
			// {
			//   "description": "The data of the component.",
			//   "maxLength": 16000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The data of the component.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 16000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Data is a write-only property.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the component.",
			//   "type": "string"
			// }
			Description: "The description of the component.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"encrypted": {
			// Property: Encrypted
			// CloudFormation resource type schema:
			// {
			//   "description": "The encryption status of the component.",
			//   "type": "boolean"
			// }
			Description: "The encryption status of the component.",
			Type:        types.BoolType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key identifier used to encrypt the component.",
			//   "type": "string"
			// }
			Description: "The KMS key identifier used to encrypt the component.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the component.",
			//   "type": "string"
			// }
			Description: "The name of the component.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"platform": {
			// Property: Platform
			// CloudFormation resource type schema:
			// {
			//   "description": "The platform of the component.",
			//   "enum": [
			//     "Windows",
			//     "Linux"
			//   ],
			//   "type": "string"
			// }
			Description: "The platform of the component.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"Windows",
					"Linux",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
			// Platform is a write-only property.
		},
		"supported_os_versions": {
			// Property: SupportedOsVersions
			// CloudFormation resource type schema:
			// {
			//   "description": "The operating system (OS) version supported by the component.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The operating system (OS) version supported by the component.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags associated with the component.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags associated with the component.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the component denotes whether the component is used to build the image or only to test it. ",
			//   "enum": [
			//     "BUILD",
			//     "TEST"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of the component denotes whether the component is used to build the image or only to test it. ",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"uri": {
			// Property: Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "The uri of the component.",
			//   "type": "string"
			// }
			Description: "The uri of the component.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Uri is a write-only property.
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the component.",
			//   "type": "string"
			// }
			Description: "The version of the component.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
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
		Description: "Resource schema for AWS::ImageBuilder::Component",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::Component").WithTerraformTypeName("awscc_imagebuilder_component")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"change_description":    "ChangeDescription",
		"data":                  "Data",
		"description":           "Description",
		"encrypted":             "Encrypted",
		"kms_key_id":            "KmsKeyId",
		"name":                  "Name",
		"platform":              "Platform",
		"supported_os_versions": "SupportedOsVersions",
		"tags":                  "Tags",
		"type":                  "Type",
		"uri":                   "Uri",
		"version":               "Version",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Data",
		"/properties/Uri",
		"/properties/Platform",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
