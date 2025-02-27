// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

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
	registry.AddResourceTypeFactory("awscc_ssmcontacts_contact_channel", contactChannelResourceType)
}

// contactChannelResourceType returns the Terraform awscc_ssmcontacts_contact_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMContacts::ContactChannel resource type.
func contactChannelResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the engagement to a contact channel.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the engagement to a contact channel.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"channel_address": {
			// Property: ChannelAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "The details that SSM Incident Manager uses when trying to engage the contact channel.",
			//   "type": "string"
			// }
			Description: "The details that SSM Incident Manager uses when trying to engage the contact channel.",
			Type:        types.StringType,
			Optional:    true,
		},
		"channel_name": {
			// Property: ChannelName
			// CloudFormation resource type schema:
			// {
			//   "description": "The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
		},
		"channel_type": {
			// Property: ChannelType
			// CloudFormation resource type schema:
			// {
			//   "description": "Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.",
			//   "enum": [
			//     "SMS",
			//     "VOICE",
			//     "EMAIL"
			//   ],
			//   "type": "string"
			// }
			Description: "Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SMS",
					"VOICE",
					"EMAIL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"contact_id": {
			// Property: ContactId
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the contact resource",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "arn:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]*:[0-9]+:([\\w+=\\/,.@:-]+)*",
			//   "type": "string"
			// }
			Description: "ARN of the contact resource",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
				validate.StringMatch(regexp.MustCompile("arn:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]*:[0-9]+:([\\w+=\\/,.@:-]+)*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"defer_activation": {
			// Property: DeferActivation
			// CloudFormation resource type schema:
			// {
			//   "description": "If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.",
			//   "type": "boolean"
			// }
			Description: "If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.",
			Type:        types.BoolType,
			Optional:    true,
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
		Description: "Resource Type definition for AWS::SSMContacts::ContactChannel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::ContactChannel").WithTerraformTypeName("awscc_ssmcontacts_contact_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":              "Arn",
		"channel_address":  "ChannelAddress",
		"channel_name":     "ChannelName",
		"channel_type":     "ChannelType",
		"contact_id":       "ContactId",
		"defer_activation": "DeferActivation",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"contact_id",
			"channel_name",
			"channel_type",
			"channel_address",
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
