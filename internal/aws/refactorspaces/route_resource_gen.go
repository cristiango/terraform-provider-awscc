// Code generated by generators/resource/main.go; DO NOT EDIT.

package refactorspaces

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
	registry.AddResourceTypeFactory("awscc_refactorspaces_route", routeResourceType)
}

// routeResourceType returns the Terraform awscc_refactorspaces_route resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RefactorSpaces::Route resource type.
func routeResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_identifier": {
			// Property: ApplicationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^app-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(14, 14),
				validate.StringMatch(regexp.MustCompile("^app-([0-9A-Za-z]{10}$)"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:(aws[a-zA-Z-]*)?:refactor-spaces:[a-zA-Z0-9\\-]+:\\w{12}:[a-zA-Z_0-9+=,.@\\-_/]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"default_route": {
			// Property: DefaultRoute
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ActivationState": {
			//       "enum": [
			//         "INACTIVE",
			//         "ACTIVE"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ActivationState"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"activation_state": {
						// Property: ActivationState
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INACTIVE",
								"ACTIVE",
							}),
						},
					},
				},
			),
			Optional: true,
			// DefaultRoute is a write-only property.
		},
		"environment_identifier": {
			// Property: EnvironmentIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^env-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(14, 14),
				validate.StringMatch(regexp.MustCompile("^env-([0-9A-Za-z]{10}$)"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"path_resource_to_id": {
			// Property: PathResourceToId
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
		"route_identifier": {
			// Property: RouteIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^rte-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"route_type": {
			// Property: RouteType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "URI_PATH"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DEFAULT",
					"URI_PATH",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// RouteType is a write-only property.
		},
		"service_identifier": {
			// Property: ServiceIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^svc-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(14, 14),
				validate.StringMatch(regexp.MustCompile("^svc-([0-9A-Za-z]{10}$)"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
			// ServiceIdentifier is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Environment resource",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"uri_path_route": {
			// Property: UriPathRoute
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ActivationState": {
			//       "enum": [
			//         "INACTIVE",
			//         "ACTIVE"
			//       ],
			//       "type": "string"
			//     },
			//     "IncludeChildPaths": {
			//       "type": "boolean"
			//     },
			//     "Methods": {
			//       "insertionOrder": false,
			//       "items": {
			//         "enum": [
			//           "DELETE",
			//           "GET",
			//           "HEAD",
			//           "OPTIONS",
			//           "PATCH",
			//           "POST",
			//           "PUT"
			//         ],
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "SourcePath": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "pattern": "^(/[a-zA-Z0-9._-]+)+$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ActivationState"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"activation_state": {
						// Property: ActivationState
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INACTIVE",
								"ACTIVE",
							}),
						},
					},
					"include_child_paths": {
						// Property: IncludeChildPaths
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"methods": {
						// Property: Methods
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayForEach(validate.StringInSlice([]string{
								"DELETE",
								"GET",
								"HEAD",
								"OPTIONS",
								"PATCH",
								"POST",
								"PUT",
							})),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"source_path": {
						// Property: SourcePath
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
							validate.StringMatch(regexp.MustCompile("^(/[a-zA-Z0-9._-]+)+$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
				},
			),
			Optional: true,
			// UriPathRoute is a write-only property.
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
		Description: "Definition of AWS::RefactorSpaces::Route Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Route").WithTerraformTypeName("awscc_refactorspaces_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activation_state":       "ActivationState",
		"application_identifier": "ApplicationIdentifier",
		"arn":                    "Arn",
		"default_route":          "DefaultRoute",
		"environment_identifier": "EnvironmentIdentifier",
		"include_child_paths":    "IncludeChildPaths",
		"key":                    "Key",
		"methods":                "Methods",
		"path_resource_to_id":    "PathResourceToId",
		"route_identifier":       "RouteIdentifier",
		"route_type":             "RouteType",
		"service_identifier":     "ServiceIdentifier",
		"source_path":            "SourcePath",
		"tags":                   "Tags",
		"uri_path_route":         "UriPathRoute",
		"value":                  "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RouteType",
		"/properties/ServiceIdentifier",
		"/properties/DefaultRoute",
		"/properties/UriPathRoute",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
