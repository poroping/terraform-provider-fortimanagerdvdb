// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.4.0,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Neighbor.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccRouterBfdNeighbor_basic(t *testing.T) {
	rName := "router_bfd_neighbor"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
