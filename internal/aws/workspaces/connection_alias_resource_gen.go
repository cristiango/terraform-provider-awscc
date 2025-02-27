// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspaces

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
	registry.AddResourceTypeFactory("awscc_workspaces_connection_alias", connectionAliasResourceType)
}

// connectionAliasResourceType returns the Terraform awscc_workspaces_connection_alias resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WorkSpaces::ConnectionAlias resource type.
func connectionAliasResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias_id": {
			// Property: AliasId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 68,
			//   "minLength": 13,
			//   "pattern": "^wsca-[0-9a-z]{8,63}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"associations": {
			// Property: Associations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AssociatedAccountId": {
			//         "type": "string"
			//       },
			//       "AssociationStatus": {
			//         "enum": [
			//           "NOT_ASSOCIATED",
			//           "PENDING_ASSOCIATION",
			//           "ASSOCIATED_WITH_OWNER_ACCOUNT",
			//           "ASSOCIATED_WITH_SHARED_ACCOUNT",
			//           "PENDING_DISASSOCIATION"
			//         ],
			//         "type": "string"
			//       },
			//       "ConnectionIdentifier": {
			//         "maxLength": 20,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9]+$",
			//         "type": "string"
			//       },
			//       "ResourceId": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "pattern": ".+",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxLength": 25,
			//   "minLength": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"associated_account_id": {
						// Property: AssociatedAccountId
						Type:     types.StringType,
						Optional: true,
					},
					"association_status": {
						// Property: AssociationStatus
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NOT_ASSOCIATED",
								"PENDING_ASSOCIATION",
								"ASSOCIATED_WITH_OWNER_ACCOUNT",
								"ASSOCIATED_WITH_SHARED_ACCOUNT",
								"PENDING_DISASSOCIATION",
							}),
						},
					},
					"connection_identifier": {
						// Property: ConnectionIdentifier
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 20),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+$"), ""),
						},
					},
					"resource_id": {
						// Property: ResourceId
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1000),
							validate.StringMatch(regexp.MustCompile(".+"), ""),
						},
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"connection_alias_state": {
			// Property: ConnectionAliasState
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CREATING",
			//     "CREATED",
			//     "DELETING"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"connection_string": {
			// Property: ConnectionString
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[.0-9a-zA-Z\\-]{1,255}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("^[.0-9a-zA-Z\\-]{1,255}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::WorkSpaces::ConnectionAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpaces::ConnectionAlias").WithTerraformTypeName("awscc_workspaces_connection_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_id":               "AliasId",
		"associated_account_id":  "AssociatedAccountId",
		"association_status":     "AssociationStatus",
		"associations":           "Associations",
		"connection_alias_state": "ConnectionAliasState",
		"connection_identifier":  "ConnectionIdentifier",
		"connection_string":      "ConnectionString",
		"key":                    "Key",
		"resource_id":            "ResourceId",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
