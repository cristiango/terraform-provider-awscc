// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sample_test

import (
	"regexp"
	"testing"

	"github.com/eVisionSoftware/axiom/terraform-provider-axiom/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAxiomSampleResourceTypeDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "Axiom::Sample::ResourceType", "awscc_sample_resource_type", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAxiomSampleResourceTypeDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "Axiom::Sample::ResourceType", "awscc_sample_resource_type", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
