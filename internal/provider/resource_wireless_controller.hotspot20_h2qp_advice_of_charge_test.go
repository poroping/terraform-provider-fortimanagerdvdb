// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure advice of charge.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWirelessControllerHotspot20H2qpAdviceOfCharge_basic(t *testing.T) {
	rName := "wireless_controller.hotspot20_h2qp_advice_of_charge"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
