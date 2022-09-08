package generic

import (
	"context"

	"github.com/eVisionSoftware/axiom/terraform-provider-axiom/internal/tfresource"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// CopyValueAtPath copies the value at a specified path from source State to destination State.
func CopyValueAtPath(ctx context.Context, dst, src *tfsdk.State, p path.Path) error {
	var val attr.Value
	diags := src.GetAttribute(ctx, p, &val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	diags = dst.SetAttribute(ctx, p, val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	return nil
}
