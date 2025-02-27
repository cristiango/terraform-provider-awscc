package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

// pluralDataSourceType is a type alias for a data source type.
type pluralDataSourceType genericDataSourceType

// NewPluralDataSourceType returns a new pluralDataSourceType from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewPluralDataSourceType(_ context.Context, optFns ...DataSourceTypeOptionsFunc) (provider.DataSourceType, error) {
	dataSourceType := &genericDataSourceType{}

	for _, optFn := range optFns {
		err := optFn(dataSourceType)

		if err != nil {
			return nil, err
		}
	}

	if dataSourceType.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}

	if dataSourceType.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	pluralDataSourceType := pluralDataSourceType(*dataSourceType)

	return &pluralDataSourceType, nil
}

func (pdt *pluralDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return pdt.tfSchema, nil
}

func (pdt *pluralDataSourceType) NewDataSource(ctx context.Context, provider provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	return newGenericPluralDataSource(provider, pdt), nil
}

// Implements datasource.DataSource
type pluralDataSource struct {
	provider       tfcloudcontrol.Provider
	dataSourceType *pluralDataSourceType
}

func newGenericPluralDataSource(provider provider.Provider, pluralDataSourceType *pluralDataSourceType) datasource.DataSource {
	return &pluralDataSource{
		provider:       provider.(tfcloudcontrol.Provider),
		dataSourceType: pluralDataSourceType,
	}
}

func (pd *pluralDataSource) Read(ctx context.Context, _ datasource.ReadRequest, response *datasource.ReadResponse) {
	ctx = pd.cfnTypeContext(ctx)

	traceEntry(ctx, "PluralDataSource.Read")

	conn := pd.provider.CloudControlApiClient(ctx)

	descriptions, err := pd.list(ctx, conn)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("CloudControl", "ListResources", err))

		return
	}

	val := GetCloudControlResourceDescriptionsValue(pd.provider.Region(ctx), descriptions)

	response.State = tfsdk.State{
		Schema: pd.dataSourceType.tfSchema,
		Raw:    val,
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "PluralDataSource.Read")
}

// list returns the ResourceDescriptions of the specified CloudFormation type.
func (pd *pluralDataSource) list(ctx context.Context, conn *cloudcontrol.Client) ([]cctypes.ResourceDescription, error) {
	return tfcloudcontrol.ListResourcesByTypeName(ctx, conn, pd.provider.RoleARN(ctx), pd.dataSourceType.cfTypeName)
}

// cfnTypeContext injects the CloudFormation type name into logger contexts.
func (pd *pluralDataSource) cfnTypeContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCFNType, pd.dataSourceType.cfTypeName)

	return ctx
}

// GetCloudControlResourceDescriptionsValue returns the Terraform Value for the specified Cloud Control API ResourceDescriptions.
func GetCloudControlResourceDescriptionsValue(id string, descriptions []cctypes.ResourceDescription) tftypes.Value {
	m := map[string]tftypes.Value{
		"id": tftypes.NewValue(tftypes.String, id),
	}

	ids := make([]tftypes.Value, 0, len(descriptions))

	for _, description := range descriptions {
		ids = append(ids, tftypes.NewValue(tftypes.String, aws.ToString(description.Identifier)))
	}

	m["ids"] = tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, ids)

	return tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"id":  tftypes.String,
			"ids": tftypes.Set{ElementType: tftypes.String},
		}}, m)
}
