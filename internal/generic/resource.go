package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
	"github.com/mattbaird/jsonpatch"
)

// ResourceTypeOptionsFunc is a type alias for a resource type functional option.
type ResourceTypeOptionsFunc func(*genericResourceType) error

// resourceWithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name maps.
// If multiple resourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func resourceWithAttributeNameMap(v map[string]string) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		if _, ok := v["id"]; !ok {
			// Synthesize a mapping for the reserved top-level "id" attribute.
			v["id"] = "ID"
		}

		cfToTfNameMap := make(map[string]string, len(v))

		for tfName, cfName := range v {
			_, ok := cfToTfNameMap[cfName]
			if ok {
				return fmt.Errorf("duplicate attribute name mapping for CloudFormation property %s", cfName)
			}
			cfToTfNameMap[cfName] = tfName
		}

		o.tfToCfNameMap = v
		o.cfToTfNameMap = cfToTfNameMap

		return nil
	}
}

// resourceWithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name.
// If multiple resourceWithCloudFormationTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithCloudFormationTypeName(v string) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.cfTypeName = v

		return nil
	}
}

// resourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple resourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformSchema(v tfsdk.Schema) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.tfSchema = v

		return nil
	}
}

// resourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple resourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformTypeName(v string) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.tfTypeName = v

		return nil
	}
}

// resourceIsImmutableType is a helper function to construct functional options
// that set a resource type's immutability flag.
// If multiple resourceIsImmutableType calls are made, the last call overrides
// the previous calls' values.
func resourceIsImmutableType(v bool) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.isImmutableType = v

		return nil
	}
}

// resourceWithSyntheticIDAttribute is a helper function to construct functional options
// that set a resource type's synthetic ID attribute flag.
// If multiple resourceWithSyntheticIDAttribute calls are made, the last call overrides
// the previous calls' values.
func resourceWithSyntheticIDAttribute(v bool) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.syntheticIDAttribute = v

		return nil
	}
}

// resourceWithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths (JSON Pointer).
// If multiple resourceWithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func resourceWithWriteOnlyPropertyPaths(v []string) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		writeOnlyAttributePaths := make([]*path.Path, 0)

		for _, writeOnlyPropertyPath := range v {
			writeOnlyAttributePath, err := o.propertyPathToAttributePath(writeOnlyPropertyPath)

			if err != nil {
				return fmt.Errorf("creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
			}

			writeOnlyAttributePaths = append(writeOnlyAttributePaths, writeOnlyAttributePath)
		}

		o.writeOnlyAttributePaths = writeOnlyAttributePaths

		return nil
	}
}

const (
	resourceMaxWaitTimeCreate = 120 * time.Minute
	resourceMaxWaitTimeUpdate = 120 * time.Minute
	resourceMaxWaitTimeDelete = 120 * time.Minute
)

// resourceWithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout (in minutes).
// If multiple resourceWithCreateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithCreateTimeoutInMinutes(v int) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		if v > 0 {
			o.createTimeout = time.Duration(v) * time.Minute
		} else {
			o.createTimeout = resourceMaxWaitTimeCreate
		}

		return nil
	}
}

// resourceWithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout (in minutes).
// If multiple resourceWithUpdateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithUpdateTimeoutInMinutes(v int) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		if v > 0 {
			o.updateTimeout = time.Duration(v) * time.Minute
		} else {
			o.updateTimeout = resourceMaxWaitTimeUpdate
		}

		return nil
	}
}

// resourceWithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout (in minutes).
// If multiple resourceWithDeleteTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithDeleteTimeoutInMinutes(v int) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		if v > 0 {
			o.deleteTimeout = time.Duration(v) * time.Minute
		} else {
			o.deleteTimeout = resourceMaxWaitTimeDelete
		}

		return nil
	}
}

// resourceWithRequiredAttributesValidators is a helper function to construct functional options
// that set a resource type's required attributes validators.
// If multiple resourceWithRequiredAttributesValidators calls are made, the last call overrides
// the previous calls' values.
func resourceWithRequiredAttributesValidators(fs ...validate.RequiredAttributesFunc) ResourceTypeOptionsFunc {
	return func(o *genericResourceType) error {
		o.requiredAttributesValidators = fs

		return nil
	}
}

// ResourceTypeOptions is a type alias for a slice of resource type functional options.
type ResourceTypeOptions []ResourceTypeOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithAttributeNameMap(v map[string]string) ResourceTypeOptions {
	return append(opts, resourceWithAttributeNameMap(v))
}

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithCloudFormationTypeName(v string) ResourceTypeOptions {
	return append(opts, resourceWithCloudFormationTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithTerraformSchema(v tfsdk.Schema) ResourceTypeOptions {
	return append(opts, resourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithTerraformTypeName(v string) ResourceTypeOptions {
	return append(opts, resourceWithTerraformTypeName(v))
}

// IsImmutableType is a helper function to construct functional options
// that set a resource type's Terraform immutability flag, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) IsImmutableType(v bool) ResourceTypeOptions {
	return append(opts, resourceIsImmutableType(v))
}

// WithSyntheticIDAttribute is a helper function to construct functional options
// that set a resource type's synthetic ID attribute flag, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithSyntheticIDAttribute(v bool) ResourceTypeOptions {
	return append(opts, resourceWithSyntheticIDAttribute(v))
}

// WithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithWriteOnlyPropertyPaths(v []string) ResourceTypeOptions {
	return append(opts, resourceWithWriteOnlyPropertyPaths(v))
}

// WithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithCreateTimeoutInMinutes(v int) ResourceTypeOptions {
	return append(opts, resourceWithCreateTimeoutInMinutes(v))
}

// WithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithUpdateTimeoutInMinutes(v int) ResourceTypeOptions {
	return append(opts, resourceWithUpdateTimeoutInMinutes(v))
}

// WithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithDeleteTimeoutInMinutes(v int) ResourceTypeOptions {
	return append(opts, resourceWithDeleteTimeoutInMinutes(v))
}

// WithRequiredAttributesValidator is a helper function to construct functional options
// that set a resource type's required attribyte validator, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithRequiredAttributesValidators(v ...validate.RequiredAttributesFunc) ResourceTypeOptions {
	return append(opts, resourceWithRequiredAttributesValidators(v...))
}

// genericResourceType implements provider.ResourceType.
type genericResourceType struct {
	cfTypeName                   string                            // CloudFormation type name for the resource type
	tfSchema                     tfsdk.Schema                      // Terraform schema for the resource type
	tfTypeName                   string                            // Terraform type name for resource type
	tfToCfNameMap                map[string]string                 // Map of Terraform attribute name to CloudFormation property name
	cfToTfNameMap                map[string]string                 // Map of CloudFormation property name to Terraform attribute name
	isImmutableType              bool                              // Resources cannot be updated and must be recreated
	syntheticIDAttribute         bool                              // Resource type has a synthetic ID attribute
	writeOnlyAttributePaths      []*path.Path                      // Paths to any write-only attributes
	createTimeout                time.Duration                     // Maximum wait time for resource creation
	updateTimeout                time.Duration                     // Maximum wait time for resource update
	deleteTimeout                time.Duration                     // Maximum wait time for resource deletion
	requiredAttributesValidators []validate.RequiredAttributesFunc // Required attributes validators
}

// NewResourceType returns a new ResourceType from the specified varidaic list of functional options.
// It's public as it's called from generated code.
func NewResourceType(_ context.Context, optFns ...ResourceTypeOptionsFunc) (provider.ResourceType, error) {
	resourceType := &genericResourceType{}

	for _, optFn := range optFns {
		err := optFn(resourceType)

		if err != nil {
			return nil, err
		}
	}

	if resourceType.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}
	if resourceType.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	return resourceType, nil
}

func (rt *genericResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return rt.tfSchema, nil
}

func (rt *genericResourceType) NewResource(ctx context.Context, provider provider.Provider) (resource.Resource, diag.Diagnostics) {
	return newGenericResource(provider, rt), nil
}

// propertyPathToAttributePath returns the AttributePath for the specified JSON Pointer property path.
func (rt *genericResourceType) propertyPathToAttributePath(propertyPath string) (*path.Path, error) {
	segments := strings.Split(propertyPath, "/")

	if got, expected := len(segments), 3; got < expected {
		return nil, fmt.Errorf("expected at least %d property path segments, got: %d", expected, got)
	}

	if got, expected := segments[0], ""; got != expected {
		return nil, fmt.Errorf("expected %q for the initial property path segment, got: %q", expected, got)
	}

	if got, expected := segments[1], "properties"; got != expected {
		return nil, fmt.Errorf("expected %q for the second property path segment, got: %q", expected, got)
	}

	attributePath := path.Empty()

	for _, segment := range segments[2:] {
		switch segment {
		case "", "*":
			return nil, fmt.Errorf("invalid property path segment: %q", segment)

		default:
			attributeName, ok := rt.cfToTfNameMap[segment]
			if !ok {
				return nil, fmt.Errorf("attribute name mapping not found: %s", segment)
			}
			attributePath = attributePath.AtName(attributeName)
		}
	}

	return &attributePath, nil
}

// Implements genericResource.Resource.
type genericResource struct {
	provider     tfcloudcontrol.Provider
	resourceType *genericResourceType
}

func newGenericResource(provider provider.Provider, resourceType *genericResourceType) resource.Resource {
	return &genericResource{
		provider:     provider.(tfcloudcontrol.Provider),
		resourceType: resourceType,
	}
}

var (
	// Path to the "id" attribute which uniquely (for a specific resource type) identifies the resource.
	// This attribute is required for acceptance testing.
	idAttributePath = path.Root("id")
)

func (r *genericResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	ctx = r.cfnTypeContext(ctx)

	traceEntry(ctx, "Resource.Create")

	conn := r.provider.CloudControlApiClient(ctx)

	tflog.Debug(ctx, "Request.Plan.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.Plan.Raw),
	})

	translator := toCloudControl{tfToCfNameMap: r.resourceType.tfToCfNameMap}
	desiredState, err := translator.AsString(ctx, &request.Plan.Schema, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Plan", err))

		return
	}

	tflog.Debug(ctx, "CloudControl DesiredState", map[string]interface{}{
		"value": desiredState,
	})

	input := &cloudcontrol.CreateResourceInput{
		ClientToken:  aws.String(tfresource.UniqueId()),
		DesiredState: aws.String(desiredState),
		TypeName:     aws.String(r.resourceType.cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.CreateResource(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("Cloud Control API", "CreateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("Cloud Control API", "CreateResource"))

		return
	}

	var progressEvent *cctypes.ProgressEvent
	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(conn, func(o *cloudcontrol.ResourceRequestSuccessWaiterOptions) {
		o.Retryable = tfcloudcontrol.RetryGetResourceRequestStatus(&progressEvent)
	})

	err = waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, r.resourceType.createTimeout)

	id := aws.ToString(progressEvent.Identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("Cloud Control API", "CreateResource", err))

		// Save any ID to state so that the resource will be marked as tainted.
		if id != "" {
			err := r.setId(ctx, id, &response.State)

			if err != nil {
				response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotSetDiag(err))
			}
		}

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the synthetic "id" attribute.
	if r.resourceType.syntheticIDAttribute {
		err = r.setId(ctx, id, &response.State)

		if err != nil {
			response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotSetDiag(err))

			return
		}
	}

	diags := r.populateUnknownValues(ctx, id, &response.State)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Create")
}

func (r *genericResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	ctx = r.cfnTypeContext(ctx)

	traceEntry(ctx, "Resource.Read")

	tflog.Debug(ctx, "Request.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.State.Raw),
	})

	conn := r.provider.CloudControlApiClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := r.describe(ctx, conn, id)

	if tfresource.NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, ResourceNotFoundWarningDiag(err))
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return
	}

	translator := toTerraform{cfToTfNameMap: r.resourceType.cfToTfNameMap}
	schema := &currentState.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.Properties))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	response.State = tfsdk.State{
		Schema: *schema,
		Raw:    val,
	}

	// Copy over any write-only values.
	// They can only be in the current state.
	for _, path := range r.resourceType.writeOnlyAttributePaths {
		err = CopyValueAtPath(ctx, &response.State, &request.State, *path)

		if err != nil {
			response.Diagnostics.AddError(
				"Terraform State Value Not Set",
				fmt.Sprintf("Unable to set Terraform State value %s. This is typically an error with the Terraform provider implementation. Original Error: %s", path, err.Error()),
			)

			return
		}
	}

	// Set the "id" attribute.
	if r.resourceType.syntheticIDAttribute {
		err = r.setId(ctx, id, &response.State)

		if err != nil {
			response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotSetDiag(err))

			return
		}
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Read")
}

func (r *genericResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	ctx = r.cfnTypeContext(ctx)

	traceEntry(ctx, "Resource.Update")

	conn := r.provider.CloudControlApiClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	translator := toCloudControl{tfToCfNameMap: r.resourceType.tfToCfNameMap}
	currentDesiredState, err := translator.AsString(ctx, &currentState.Schema, currentState.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Prior State", err))

		return
	}

	plannedDesiredState, err := translator.AsString(ctx, &request.Plan.Schema, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Plan", err))

		return
	}

	patchDocument, err := patchDocument(currentDesiredState, plannedDesiredState)

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of JSON Patch Unsuccessful",
			fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	tflog.Debug(ctx, "Cloud Control API PatchDocument", map[string]interface{}{
		"value": patchDocument,
	})

	input := &cloudcontrol.UpdateResourceInput{
		ClientToken:   aws.String(tfresource.UniqueId()),
		Identifier:    aws.String(id),
		PatchDocument: aws.String(patchDocument),
		TypeName:      aws.String(r.resourceType.cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.UpdateResource(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("Cloud Control API", "UpdateResource"))

		return
	}

	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(conn, func(o *cloudcontrol.ResourceRequestSuccessWaiterOptions) {
		o.Retryable = tfcloudcontrol.RetryGetResourceRequestStatus(nil)
	})

	err = waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, r.resourceType.updateTimeout)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("Cloud Control API", "UpdateResource", err))

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	diags := r.populateUnknownValues(ctx, id, &response.State)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)

		return
	}

	traceExit(ctx, "Resource.Update")
}

func (r *genericResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	ctx = r.cfnTypeContext(ctx)

	traceEntry(ctx, "Resource.Delete")

	conn := r.provider.CloudControlApiClient(ctx)

	id, err := r.getId(ctx, &request.State)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	err = tfcloudcontrol.DeleteResource(ctx, conn, r.provider.RoleARN(ctx), r.resourceType.cfTypeName, id, r.resourceType.deleteTimeout)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("Cloud Control API", "DeleteResource", err))

		return
	}

	response.State.RemoveResource(ctx)

	traceExit(ctx, "Resource.Delete")
}

func (r *genericResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	ctx = r.cfnTypeContext(ctx)

	traceEntry(ctx, "Resource.ImportState")

	tflog.Debug(ctx, "Request.ID", map[string]interface{}{
		"value": hclog.Fmt("%v", request.ID),
	})

	resource.ImportStatePassthroughID(ctx, idAttributePath, request, response)

	traceExit(ctx, "Resource.ImportState")
}

// ConfigValidators returns a list of functions which will all be performed during validation.
func (r *genericResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	validators := make([]resource.ConfigValidator, 0)

	if len(r.resourceType.requiredAttributesValidators) > 0 {
		validators = append(validators, validate.ResourceConfigRequiredAttributes(r.resourceType.requiredAttributesValidators...))
	}

	return validators
}

// describe returns the live state of the specified resource.
func (r *genericResource) describe(ctx context.Context, conn *cloudcontrol.Client, id string) (*cctypes.ResourceDescription, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, r.provider.RoleARN(ctx), r.resourceType.cfTypeName, id)
}

// getId returns the resource's primary identifier value from State.
func (r *genericResource) getId(ctx context.Context, state *tfsdk.State) (string, error) {
	var val string
	diags := state.GetAttribute(ctx, idAttributePath, &val)

	if diags.HasError() {
		return "", tfresource.DiagsError(diags)
	}

	return val, nil
}

// setId sets the resource's primary identifier value in State.
func (r *genericResource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	return nil
}

// populateUnknownValues populates and unknown values in State with values from the current resource description.
func (r *genericResource) populateUnknownValues(ctx context.Context, id string, state *tfsdk.State) diag.Diagnostics {
	var diags diag.Diagnostics

	unknowns, err := UnknownValuePaths(ctx, state.Raw)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}

	if len(unknowns) == 0 {
		return nil
	}

	description, err := r.describe(ctx, r.provider.CloudControlApiClient(ctx), id)

	if tfresource.NotFound(err) {
		diags.Append(ResourceNotFoundAfterWriteDiag(err))

		return diags
	}

	if err != nil {
		diags.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return diags
	}

	if description == nil {
		diags.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "GetResource"))

		return diags
	}

	err = SetUnknownValuesFromResourceModel(ctx, state, unknowns, aws.ToString(description.Properties), r.resourceType.cfToTfNameMap)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}

	return nil
}

// cfnTypeContext injects the CloudFormation type name into logger contexts.
func (r *genericResource) cfnTypeContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCFNType, r.resourceType.cfTypeName)

	return ctx
}

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(patch)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
