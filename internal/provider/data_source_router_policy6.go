// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 routing policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func dataSourceRouterPolicy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 routing policies.",

		ReadContext: dataSourceRouterPolicy6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"deviceparam": {
				Type:        schema.TypeString,
				Description: "Specifies the device to which the resource will be applied to on the FortiManager. If you want to inherit the configuration from the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeString,
				Description: "Destination IPv6 prefix.",
				Computed:    true,
			},
			"end_port": {
				Type:        schema.TypeInt,
				Description: "End destination port number (1 - 65535).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the gateway.",
				Computed:    true,
			},
			"input_device": {
				Type:        schema.TypeList,
				Description: "Incoming interface name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"output_device": {
				Type:        schema.TypeString,
				Description: "Outgoing interface name.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Protocol number (0 - 255).",
				Computed:    true,
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number(1-65535).",
				Computed:    true,
			},
			"src": {
				Type:        schema.TypeString,
				Description: "Source IPv6 prefix.",
				Computed:    true,
			},
			"start_port": {
				Type:        schema.TypeInt,
				Description: "Start destination port number (1 - 65535).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this policy route.",
				Computed:    true,
			},
			"tos": {
				Type:        schema.TypeString,
				Description: "Type of service bit pattern.",
				Computed:    true,
			},
			"tos_mask": {
				Type:        schema.TypeString,
				Description: "Type of service evaluated bits.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterPolicy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	deviceparam := ""
	if v, ok := d.GetOk("deviceparam"); ok {
		if s, ok := v.(string); ok {
			deviceparam = s
		}
	}
	urlparams.Scope = deviceparam

	i := d.Get("seq_num")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadRouterPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPolicy6 dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectRouterPolicy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
