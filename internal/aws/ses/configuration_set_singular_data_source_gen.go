// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ses_configuration_set", configurationSetDataSourceType)
}

// configurationSetDataSourceType returns the Terraform awscc_ses_configuration_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SES::ConfigurationSet resource type.
func configurationSetDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"delivery_options": {
			// Property: DeliveryOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.",
			//   "properties": {
			//     "SendingPoolName": {
			//       "description": "The name of the dedicated IP pool to associate with the configuration set.",
			//       "type": "string"
			//     },
			//     "TlsPolicy": {
			//       "description": "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.",
			//       "pattern": "REQUIRE|OPTIONAL",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sending_pool_name": {
						// Property: SendingPoolName
						Description: "The name of the dedicated IP pool to associate with the configuration set.",
						Type:        types.StringType,
						Computed:    true,
					},
					"tls_policy": {
						// Property: TlsPolicy
						Description: "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the configuration set.",
			//   "pattern": "^[a-zA-Z0-9_-]{1,64}$",
			//   "type": "string"
			// }
			Description: "The name of the configuration set.",
			Type:        types.StringType,
			Computed:    true,
		},
		"reputation_options": {
			// Property: ReputationOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.",
			//   "properties": {
			//     "ReputationMetricsEnabled": {
			//       "description": "If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.",
			//       "pattern": "true|false",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"reputation_metrics_enabled": {
						// Property: ReputationMetricsEnabled
						Description: "If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"sending_options": {
			// Property: SendingOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that defines whether or not Amazon SES can send email that you send using the configuration set.",
			//   "properties": {
			//     "SendingEnabled": {
			//       "pattern": "true|false",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that defines whether or not Amazon SES can send email that you send using the configuration set.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sending_enabled": {
						// Property: SendingEnabled
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"suppression_options": {
			// Property: SuppressionOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that contains information about the suppression list preferences for your account.",
			//   "properties": {
			//     "SuppressedReasons": {
			//       "description": "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "The reason that the address was added to the suppression list for your account",
			//         "pattern": "BOUNCE|COMPLAINT",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that contains information about the suppression list preferences for your account.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"suppressed_reasons": {
						// Property: SuppressedReasons
						Description: "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tracking_options": {
			// Property: TrackingOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that defines the open and click tracking options for emails that you send using the configuration set.",
			//   "properties": {
			//     "CustomRedirectDomain": {
			//       "description": "The domain to use for tracking open and click events.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that defines the open and click tracking options for emails that you send using the configuration set.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"custom_redirect_domain": {
						// Property: CustomRedirectDomain
						Description: "The domain to use for tracking open and click events.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SES::ConfigurationSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ConfigurationSet").WithTerraformTypeName("awscc_ses_configuration_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"custom_redirect_domain":     "CustomRedirectDomain",
		"delivery_options":           "DeliveryOptions",
		"name":                       "Name",
		"reputation_metrics_enabled": "ReputationMetricsEnabled",
		"reputation_options":         "ReputationOptions",
		"sending_enabled":            "SendingEnabled",
		"sending_options":            "SendingOptions",
		"sending_pool_name":          "SendingPoolName",
		"suppressed_reasons":         "SuppressedReasons",
		"suppression_options":        "SuppressionOptions",
		"tls_policy":                 "TlsPolicy",
		"tracking_options":           "TrackingOptions",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
