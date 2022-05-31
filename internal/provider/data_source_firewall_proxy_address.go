// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web proxy address.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortimanagerdvdb/utils"
)

func dataSourceFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web proxy address.",

		ReadContext: dataSourceFirewallProxyAddressRead,

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
			"case_sensitivity": {
				Type:        schema.TypeString,
				Description: "Enable to make the pattern case sensitive.",
				Computed:    true,
			},
			"category": {
				Type:        schema.TypeList,
				Description: "FortiGuard category ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "FortiGuard category ID.",
							Computed:    true,
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"header": {
				Type:        schema.TypeString,
				Description: "HTTP header name as a regular expression.",
				Computed:    true,
			},
			"header_group": {
				Type:        schema.TypeList,
				Description: "HTTP header group.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitivity": {
							Type:        schema.TypeString,
							Description: "Case sensitivity in pattern.",
							Computed:    true,
						},
						"header": {
							Type:        schema.TypeString,
							Description: "HTTP header regular expression.",
							Computed:    true,
						},
						"header_name": {
							Type:        schema.TypeString,
							Description: "HTTP header.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
					},
				},
			},
			"header_name": {
				Type:        schema.TypeString,
				Description: "Name of HTTP header.",
				Computed:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "Address object for the host.",
				Computed:    true,
			},
			"host_regex": {
				Type:        schema.TypeString,
				Description: "Host name as a regular expression.",
				Computed:    true,
			},
			"method": {
				Type:        schema.TypeString,
				Description: "HTTP request methods to be used.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Address name.",
				Required:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "URL path as a regular expression.",
				Computed:    true,
			},
			"query": {
				Type:        schema.TypeString,
				Description: "Match the query part of the URL as a regular expression.",
				Computed:    true,
			},
			"referrer": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of referrer field in the HTTP header to match the address.",
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Tag category.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Tagging entry name.",
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Tag name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Proxy address type.",
				Computed:    true,
			},
			"ua": {
				Type:        schema.TypeString,
				Description: "Names of browsers to be used as user agent.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable visibility of the object in the GUI.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallProxyAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProxyAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProxyAddress dataSource: %v", err)
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

	diags := refreshObjectFirewallProxyAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
