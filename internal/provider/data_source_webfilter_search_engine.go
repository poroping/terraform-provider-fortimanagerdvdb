// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web filter search engines.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func dataSourceWebfilterSearchEngine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web filter search engines.",

		ReadContext: dataSourceWebfilterSearchEngineRead,

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
			"charset": {
				Type:        schema.TypeString,
				Description: "Search engine charset.",
				Computed:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "Hostname (regular expression).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Search engine name.",
				Required:    true,
			},
			"query": {
				Type:        schema.TypeString,
				Description: "Code used to prefix a query (must end with an equals character).",
				Computed:    true,
			},
			"safesearch": {
				Type:        schema.TypeString,
				Description: "Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.",
				Computed:    true,
			},
			"safesearch_str": {
				Type:        schema.TypeString,
				Description: "Safe search parameter used in the URL.",
				Computed:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL (regular expression).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterSearchEngineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWebfilterSearchEngine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterSearchEngine dataSource: %v", err)
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

	diags := refreshObjectWebfilterSearchEngine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
