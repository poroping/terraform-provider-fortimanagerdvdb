// Unofficial FortiManager Device Database Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure authentication setting.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func dataSourceAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure authentication setting.",

		ReadContext: dataSourceAuthenticationSettingRead,

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
			"active_auth_scheme": {
				Type:        schema.TypeString,
				Description: "Active authentication method (scheme name).",
				Computed:    true,
			},
			"auth_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirecting HTTP user authentication to HTTPS.",
				Computed:    true,
			},
			"captive_portal": {
				Type:        schema.TypeString,
				Description: "Captive portal host name.",
				Computed:    true,
			},
			"captive_portal_ip": {
				Type:        schema.TypeString,
				Description: "Captive portal IP address.",
				Computed:    true,
			},
			"captive_portal_ip6": {
				Type:        schema.TypeString,
				Description: "Captive portal IPv6 address.",
				Computed:    true,
			},
			"captive_portal_port": {
				Type:        schema.TypeInt,
				Description: "Captive portal port number (1 - 65535, default = 7830).",
				Computed:    true,
			},
			"captive_portal_ssl_port": {
				Type:        schema.TypeInt,
				Description: "Captive portal SSL port number (1 - 65535, default = 7831).",
				Computed:    true,
			},
			"captive_portal_type": {
				Type:        schema.TypeString,
				Description: "Captive portal type.",
				Computed:    true,
			},
			"captive_portal6": {
				Type:        schema.TypeString,
				Description: "IPv6 captive portal host name.",
				Computed:    true,
			},
			"cert_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirecting certificate authentication to HTTPS portal.",
				Computed:    true,
			},
			"cert_captive_portal": {
				Type:        schema.TypeString,
				Description: "Certificate captive portal host name.",
				Computed:    true,
			},
			"cert_captive_portal_ip": {
				Type:        schema.TypeString,
				Description: "Certificate captive portal IP address.",
				Computed:    true,
			},
			"cert_captive_portal_port": {
				Type:        schema.TypeInt,
				Description: "Certificate captive portal port number (1 - 65535, default = 7832).",
				Computed:    true,
			},
			"dev_range": {
				Type:        schema.TypeList,
				Description: "Address range for the IP based device query.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"sso_auth_scheme": {
				Type:        schema.TypeString,
				Description: "Single-Sign-On authentication method (scheme name).",
				Computed:    true,
			},
			"user_cert_ca": {
				Type:        schema.TypeList,
				Description: "CA certificate used for client certificate verification.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "CA certificate list.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAuthenticationSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "AuthenticationSetting"

	o, err := c.Cmdb.ReadAuthenticationSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationSetting dataSource: %v", err)
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

	diags := refreshObjectAuthenticationSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
