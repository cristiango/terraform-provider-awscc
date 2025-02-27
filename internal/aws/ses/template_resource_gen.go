// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

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
	registry.AddResourceTypeFactory("awscc_ses_template", templateResourceType)
}

// templateResourceType returns the Terraform awscc_ses_template resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SES::Template resource type.
func templateResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
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
		"template": {
			// Property: Template
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The content of the email, composed of a subject line, an HTML part, and a text-only part",
			//   "properties": {
			//     "HtmlPart": {
			//       "description": "The HTML body of the email.",
			//       "type": "string"
			//     },
			//     "SubjectPart": {
			//       "description": "The subject line of the email.",
			//       "type": "string"
			//     },
			//     "TemplateName": {
			//       "description": "The name of the template.",
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "pattern": "^[a-zA-Z0-9_-]{1,64}$",
			//       "type": "string"
			//     },
			//     "TextPart": {
			//       "description": "The email body that is visible to recipients whose email clients do not display HTML content.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SubjectPart"
			//   ],
			//   "type": "object"
			// }
			Description: "The content of the email, composed of a subject line, an HTML part, and a text-only part",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"html_part": {
						// Property: HtmlPart
						Description: "The HTML body of the email.",
						Type:        types.StringType,
						Optional:    true,
					},
					"subject_part": {
						// Property: SubjectPart
						Description: "The subject line of the email.",
						Type:        types.StringType,
						Required:    true,
					},
					"template_name": {
						// Property: TemplateName
						Description: "The name of the template.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 64),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"text_part": {
						// Property: TextPart
						Description: "The email body that is visible to recipients whose email clients do not display HTML content.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SES::Template",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::Template").WithTerraformTypeName("awscc_ses_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"html_part":     "HtmlPart",
		"id":            "Id",
		"subject_part":  "SubjectPart",
		"template":      "Template",
		"template_name": "TemplateName",
		"text_part":     "TextPart",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
