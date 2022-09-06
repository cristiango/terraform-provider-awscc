//go:generate go run generators/schema/main.go -config all_schemas.hcl -generated-code-root .. -import-path-root github.com/eVisionSoftware/axiom/axiom-terraform-provider/internal -- resources.go singular_data_sources.go plural_data_sources.go

package provider
