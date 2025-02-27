// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("awscc_sagemaker_image", imageResourceType)
}

// imageResourceType returns the Terraform awscc_sagemaker_image resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Image resource type.
func imageResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"image_arn": {
			// Property: ImageArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image\\/[a-z0-9]([-.]?[a-z0-9])*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"image_description": {
			// Property: ImageDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the image.",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": ".+",
			//   "type": "string"
			// }
			Description: "A description of the image.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
				validate.StringMatch(regexp.MustCompile(".+"), ""),
			},
		},
		"image_display_name": {
			// Property: ImageDisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the image.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9 -_]+$",
			//   "type": "string"
			// }
			Description: "The display name of the image.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9 -_]+$"), ""),
			},
		},
		"image_name": {
			// Property: ImageName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$",
			//   "type": "string"
			// }
			Description: "The name of the image.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"image_role_arn": {
			// Property: ImageRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:iam::[0-9]{12}:role/.*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^arn:aws(-[\\w]+)*:iam::[0-9]{12}:role/.*$"), ""),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
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
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 50),
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
		Description: "Resource Type definition for AWS::SageMaker::Image",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Image").WithTerraformTypeName("awscc_sagemaker_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"image_arn":          "ImageArn",
		"image_description":  "ImageDescription",
		"image_display_name": "ImageDisplayName",
		"image_name":         "ImageName",
		"image_role_arn":     "ImageRoleArn",
		"key":                "Key",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
