package main

import (
	"context"
	"flag"

	"github.com/eVisionSoftware/axiom/terraform-provider-axiom/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/enablon/axiom",
		Debug:   debugMode,
	})
}
