package validate

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// stringLenBetweenValidator validates that a string Attribute's length is in a range.
type stringLenBetweenValidator struct {
	minLength, maxLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", validator.minLength, validator.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if l := len(s); l < validator.minLength || l > validator.maxLength {
		response.Diagnostics.Append(ccdiag.NewInvalidLengthBetweenAttributeError(
			request.AttributePath, validator.minLength, validator.maxLength, l,
		))

		return
	}
}

// StringLenBetween returns a new string length between validator.
func StringLenBetween(minLength, maxLength int) tfsdk.AttributeValidator {
	if minLength < 0 || maxLength < 0 || minLength > maxLength {
		return nil
	}

	return stringLenBetweenValidator{
		minLength: minLength,
		maxLength: maxLength,
	}
}

// stringLenAtLeastValidator validates that a string Attribute's length is at least a certain value.
type stringLenAtLeastValidator struct {
	minLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be at least %d", validator.minLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if l := len(s); l < validator.minLength {
		response.Diagnostics.Append(ccdiag.NewInvalidLengthAtLeastAttributeError(
			request.AttributePath, validator.minLength, l,
		))

		return
	}
}

// StringLenAtLeast returns a new string length at least validator.
func StringLenAtLeast(minLength int) tfsdk.AttributeValidator {
	if minLength < 0 {
		return nil
	}

	return stringLenAtLeastValidator{
		minLength: minLength,
	}
}

// stringLenAtMostValidator validates that a string Attribute's length is at most a certain value.
type stringLenAtMostValidator struct {
	maxLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be at most %d", validator.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenAtMostValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if l := len(s); l > validator.maxLength {
		response.Diagnostics.Append(ccdiag.NewInvalidLengthAtMostAttributeError(
			request.AttributePath, validator.maxLength, l,
		))

		return
	}
}

// StringLenAtMost returns a new string length at least validator.
func StringLenAtMost(maxLength int) tfsdk.AttributeValidator {
	if maxLength < 0 {
		return nil
	}

	return stringLenAtMostValidator{
		maxLength: maxLength,
	}
}

// stringInSliceValidator validates that a string Attribute's value matches the value of an element in the valid slice.
type stringInSliceValidator struct {
	valid []string
}

// Description describes the validation in plain text formatting.
func (validator stringInSliceValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be one of %v", validator.valid)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringInSliceValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringInSliceValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	for _, val := range validator.valid {
		if s == val {
			return
		}
	}

	response.Diagnostics.Append(newStringNotInSliceError(
		request.AttributePath,
		validator.valid,
		s,
	))
}

func newStringNotInSliceError(p path.Path, valid []string, value string) diag.Diagnostic {
	return ccdiag.NewInvalidValueAttributeError(
		p,
		fmt.Sprintf("expected value to be one of %v, got %s", valid, value),
	)
}

// StringInSlice returns a new string in slice validator.
func StringInSlice(valid []string) tfsdk.AttributeValidator {
	return stringInSliceValidator{
		valid: valid,
	}
}

// stringMatchValidator validates that a string Attribute's value matches the specified regular expression.
type stringMatchValidator struct {
	re      *regexp.Regexp
	message string
}

// Description describes the validation in plain text formatting.
func (validator stringMatchValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must match regular expression '%s'", validator.re)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringMatchValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringMatchValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if ok := validator.re.MatchString(s); ok {
		return
	}

	if v := validator.message; v != "" {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(request.AttributePath, v))
	}

	response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
		request.AttributePath,
		fmt.Sprintf("expected value of %s to match regular expression '%s'", s, validator.re)))
}

// StringMatch returns a new string match validator.
func StringMatch(re *regexp.Regexp, message string) tfsdk.AttributeValidator {
	return stringMatchValidator{
		re:      re,
		message: message,
	}
}

// stringIsJsonObjectValidator validates that a string Attribute's value is a valid JSON object.
type stringIsJsonObjectValidator struct{}

// Description describes the validation in plain text formatting.
func (validator stringIsJsonObjectValidator) Description(_ context.Context) string {
	return "value must be a valid JSON object"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringIsJsonObjectValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringIsJsonObjectValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	// A JSON object starts with a '{'
	if s[:1] != "{" {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.AttributePath,
			"expected value to be a valid JSON object",
		))
		return
	}

	// Use json.Unmarshal() instead of just json.Valid() to get the parsing error
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be valid JSON: %s", err),
		))
	}
}

// StringIsJsonObject returns a new string is JSON validator.
func StringIsJsonObject() tfsdk.AttributeValidator {
	return stringIsJsonObjectValidator{}
}

func validateString(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (string, bool) {
	var s types.String
	diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &s)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return "", false
	}

	if s.Unknown || s.Null {
		return "", false
	}

	return s.Value, true
}
